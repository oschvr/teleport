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

package clusters

import (
	"context"
	"github.com/gravitational/teleport/lib/sshutils/sftp"
	api "github.com/gravitational/teleport/lib/teleterm/api/protogen/golang/v1"

	"github.com/gravitational/teleport/api/client/proto"
	"github.com/gravitational/teleport/api/defaults"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/teleterm/api/uri"

	"github.com/gravitational/trace"
)

// Database describes database
type Server struct {
	// URI is the database URI
	URI uri.ResourceURI

	types.Server
}

// GetServers returns cluster servers
func (c *Cluster) GetServers(ctx context.Context) ([]Server, error) {
	var clusterServers []types.Server
	err := addMetadataToRetryableError(ctx, func() error {
		proxyClient, err := c.clusterClient.ConnectToProxy(ctx)
		if err != nil {
			return trace.Wrap(err)
		}
		defer proxyClient.Close()

		clusterServers, err = proxyClient.FindNodesByFilters(ctx, proto.ListResourcesRequest{
			Namespace: defaults.Namespace,
		})
		if err != nil {
			return trace.Wrap(err)
		}

		return nil
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}

	results := []Server{}
	for _, server := range clusterServers {
		results = append(results, Server{
			URI:    c.URI.AppendServer(server.GetName()),
			Server: server,
		})
	}

	return results, nil
}

func (c *Cluster) TransferFile(request *api.FileTransferRequest, server api.TerminalService_TransferFileServer) error {
	proxyClient, err := c.clusterClient.ConnectToProxy(server.Context())
	if err != nil {
		return err
	}
	defer proxyClient.Close()

	var config *sftp.Config
	var configErr error

	if request.GetDirection() == api.FileTransferDirection_FILE_TRANSFER_DIRECTION_DOWNLOAD {
		config, configErr = sftp.CreateDownloadConfig(request.GetSource(), request.GetDestination(), sftp.Options{})
	} else {
		config, configErr = sftp.CreateUploadConfig([]string{request.GetSource()}, request.GetDestination(), sftp.Options{})
	}

	if configErr != nil {
		return trace.Wrap(configErr)
	}

	config.ProgressWriter = &grpcProgressWriter{TransferFileServer: server}
	config.WriteSimpleProgress = true

	clusterServers, err := proxyClient.FindNodesByFilters(server.Context(), proto.ListResourcesRequest{
		Namespace: defaults.Namespace,
	})

	var foundServer types.Server
	for _, clusterServer := range clusterServers {
		if clusterServer.GetName() == request.GetServerId() {
			foundServer = clusterServer
			break
		}
	}

	if foundServer == nil {
		return trace.BadParameter("Requested server does not exist")
	}

	err = c.clusterClient.TransferFiles(server.Context(), request.GetLogin(), foundServer.GetHostname()+":0", config)
	if err != nil {
		return trace.Wrap(err)
	}

	return nil
}

type grpcProgressWriter struct {
	TransferFileServer api.TerminalService_TransferFileServer
}

func (writer *grpcProgressWriter) Write(bytes []byte) (n int, err error) {
	err = writer.TransferFileServer.Send(&api.FileTransferProgress{Data: string(bytes)})
	if err != nil {
		return 0, err
	}
	return len(bytes), nil
}
