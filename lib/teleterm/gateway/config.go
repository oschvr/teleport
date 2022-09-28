/*
Copyright 2021 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gateway

import (
	"runtime"

	"github.com/gravitational/teleport/api/constants"
	"github.com/gravitational/teleport/lib/defaults"
	"github.com/gravitational/teleport/lib/teleterm/api/uri"
	"github.com/gravitational/trace"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Config describes gateway configuration
type Config struct {
	// URI is the gateway URI
	URI uri.ResourceURI
	// TargetName is the remote resource name
	TargetName string
	// TargetURI is the remote resource URI
	TargetURI string
	// TargetUser is the target user name
	TargetUser string
	// TargetSubresourceName points at a subresource of the remote resource, for example a database
	// name on a database server.
	TargetSubresourceName string

	// Port is the gateway port
	LocalPort string
	// LocalAddress is the local address
	LocalAddress string
	// Protocol is the gateway protocol
	Protocol string
	// CertPath
	CertPath string
	// KeyPath
	KeyPath string
	// Insecure
	Insecure bool
	// WebProxyAddr
	WebProxyAddr string
	// Log is a component logger
	Log *logrus.Entry
	// CLICommandProvider returns a CLI command for the gateway
	CLICommandProvider CLICommandProvider
	// TCPPortAllocator creates listeners on the given ports. This interface lets us avoid occupying
	// hardcoded ports in tests.
	TCPPortAllocator TCPPortAllocator
	// OnNewConnection is a callback called when a new downstream connection is accepted by the
	// gateway. The full gateway struct is not passed as an argument to the callback on purpose to
	// avoid callsites mutating the gateway without acquiring a lock on daemon.Service.mu.
	//
	// Note that the callback blocks handling of the connection.
	//
	// OnNewConnection is copied between gateways when calling gateway.NewWithLocalPort.
	OnNewConnection OnNewConnectionFunc
}

type OnNewConnectionFunc func(gatewayURI uri.ResourceURI, targetURI string)

// CheckAndSetDefaults checks and sets the defaults
func (c *Config) CheckAndSetDefaults() error {
	if c.URI.String() == "" {
		c.URI = uri.NewGatewayURI(uuid.NewString())
	}

	if c.LocalAddress == "" {
		c.LocalAddress = "localhost"
		// SQL Server Management Studio won't connect to localhost:12345, so use 127.0.0.1:12345 instead.
		if runtime.GOOS == constants.WindowsOS && c.Protocol == defaults.ProtocolSQLServer {
			c.LocalAddress = "127.0.0.1"
		}
	}

	if c.LocalPort == "" {
		c.LocalPort = "0"
	}

	if c.Log == nil {
		c.Log = logrus.NewEntry(logrus.StandardLogger())
	}

	c.Log = c.Log.WithField("resource", c.TargetURI).WithField("gateway", c.URI.String())

	if c.TargetName == "" {
		return trace.BadParameter("missing target name")
	}

	if c.TargetURI == "" {
		return trace.BadParameter("missing target URI")
	}

	if c.CLICommandProvider == nil {
		return trace.BadParameter("missing CLICommandProvider")
	}

	if c.TCPPortAllocator == nil {
		c.TCPPortAllocator = NetTCPPortAllocator{}
	}

	return nil
}
