/*
Copyright 2022 Gravitational, Inc.

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

package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

// DBServersClient provides an interface for fetching Azure DB Servers.
type DBServersClient interface {
	// ListAll returns all Azure DB servers within an Azure subscription.
	ListAll(ctx context.Context) ([]*DBServer, error)
	// ListWithinGroup returns all Azure DB servers within an Azure resource group.
	ListWithinGroup(ctx context.Context, group string) ([]*DBServer, error)
	// Get returns a DBServer within an Azure subscription, queried by group and name
	Get(ctx context.Context, group, name string) (*DBServer, error)
}

// ARMMySQL is an interface for armmysql.ServersClient.
// It exists so that the client can be mocked.
type ARMMySQL interface {
	// Get - gets information about an Azure DB server.
	Get(ctx context.Context, group, name string, opts *armmysql.ServersClientGetOptions) (armmysql.ServersClientGetResponse, error)
	// NewListPager - List all the servers in a given subscription.
	NewListPager(opts *armmysql.ServersClientListOptions) *runtime.Pager[armmysql.ServersClientListResponse]
	// NewListByResourceGroupPager - List all the servers in a given resource group.
	NewListByResourceGroupPager(group string, opts *armmysql.ServersClientListByResourceGroupOptions) *runtime.Pager[armmysql.ServersClientListByResourceGroupResponse]
}

var _ ARMMySQL = (*armmysql.ServersClient)(nil)

// ARMPostgres is an interface for armpostgresql.ServersClient.
// It exists so that the client can be mocked.
type ARMPostgres interface {
	// Get - gets information about an Azure DB server.
	Get(ctx context.Context, group, name string, opts *armpostgresql.ServersClientGetOptions) (armpostgresql.ServersClientGetResponse, error)
	// NewListPager - List all the servers in a given subscription.
	NewListPager(opts *armpostgresql.ServersClientListOptions) *runtime.Pager[armpostgresql.ServersClientListResponse]
	// NewListByResourceGroupPager - List all the servers in a given resource group.
	NewListByResourceGroupPager(group string, opts *armpostgresql.ServersClientListByResourceGroupOptions) *runtime.Pager[armpostgresql.ServersClientListByResourceGroupResponse]
}

var _ ARMPostgres = (*armpostgresql.ServersClient)(nil)
