// Copyright 2021 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apiserver

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"path/filepath"
	"strings"

	api "github.com/gravitational/teleport/lib/teleterm/api/protogen/golang/v1"
	"github.com/gravitational/teleport/lib/teleterm/apiserver/handler"
	"github.com/gravitational/teleport/lib/utils"

	"github.com/gravitational/trace"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// New creates an instance of API Server
func New(ctx context.Context, cfg Config) (*APIServer, error) {
	if err := cfg.CheckAndSetDefaults(); err != nil {
		return nil, trace.Wrap(err)
	}

	closeContext, cancel := context.WithCancel(ctx)
	ok := false
	defer func() {
		if ok {
			return
		}
		cancel()
	}()

	// Create the listener, set up the credentials and the server.

	ls, err := newListener(cfg.HostAddr)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	var tshdKeyPair tls.Certificate
	tshdCreds := grpc.Creds(nil)
	rendererCertPath := filepath.Join(cfg.CertsDir, rendererCertFileName)
	tshdCertPath := filepath.Join(cfg.CertsDir, tshdCertFileName)
	shouldUseMTLS := strings.HasPrefix(cfg.HostAddr, "tcp://")

	if shouldUseMTLS {
		tshdKeyPair, err = generateAndSaveCert(tshdCertPath)
		if err != nil {
			return nil, trace.Wrap(err)
		}

		// rendererCertPath will be read on an incoming client connection so we can assume that at this
		// point the renderer process has saved its public key under that path.
		tshdCreds, err = createServerCredentials(tshdKeyPair, rendererCertPath)
		if err != nil {
			return nil, trace.Wrap(err)
		}
	}

	grpcServer := grpc.NewServer(tshdCreds, grpc.ChainUnaryInterceptor(
		withErrorHandling(cfg.Log),
	))

	// Create Terminal service.

	serviceHandler, err := handler.New(
		handler.Config{
			DaemonService: cfg.Daemon,
		},
	)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	api.RegisterTerminalServiceServer(grpcServer, serviceHandler)

	ok = true
	return &APIServer{cfg, closeContext, cancel, ls, grpcServer}, nil
}

// Serve starts accepting incoming connections
func (s *APIServer) Serve() error {
	return s.grpcServer.Serve(s.ls)
}

// Stop stops the server and closes all listeners
func (s *APIServer) Stop() {
	s.grpcServer.GracefulStop()
	s.cancel()
}

func newListener(hostAddr string) (net.Listener, error) {
	uri, err := utils.ParseAddr(hostAddr)

	if err != nil {
		return nil, trace.BadParameter("invalid host address: %s", hostAddr)
	}

	lis, err := net.Listen(uri.Network(), uri.Addr)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	addr := utils.FromAddr(lis.Addr())
	sendBoundNetworkPortToStdout(addr)

	log.Infof("tsh daemon is listening on %v.", addr.FullAddress())

	return lis, nil
}

func sendBoundNetworkPortToStdout(addr utils.NetAddr) {
	// Connect needs this message to know which port has been assigned to the server.
	fmt.Printf("{CONNECT_GRPC_PORT: %v}\n", addr.Port(1))
}

// Server is a combination of the underlying grpc.Server and its RuntimeOpts.
type APIServer struct {
	Config
	// closeContext is canceled when the APIServer gets stopped.
	closeContext context.Context
	cancel       context.CancelFunc
	// ls is the server listener
	ls net.Listener
	// grpc is an instance of grpc server
	grpcServer *grpc.Server
}
