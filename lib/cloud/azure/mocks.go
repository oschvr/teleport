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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/gravitational/trace"
)

type ARMSubscriptionsMock struct {
	Subscriptions []*armsubscription.Subscription
	NoAuth        bool
}

var _ ARMSubscriptions = (*ARMSubscriptionsMock)(nil)

func (m *ARMSubscriptionsMock) NewListPager(_ *armsubscription.SubscriptionsClientListOptions) *runtime.Pager[armsubscription.SubscriptionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[armsubscription.SubscriptionsClientListResponse]{
		More: func(page armsubscription.SubscriptionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *armsubscription.SubscriptionsClientListResponse) (armsubscription.SubscriptionsClientListResponse, error) {
			if m.NoAuth {
				return armsubscription.SubscriptionsClientListResponse{}, trace.AccessDenied("unauthorized")
			}
			return armsubscription.SubscriptionsClientListResponse{
				ListResult: armsubscription.ListResult{
					Value: m.Subscriptions,
				},
			}, nil
		},
	})
}

// ARMMySQLMock mocks Azure armmysql API.
type ARMMySQLMock struct {
	DBServers []*armmysql.Server
	NoAuth    bool
}

var _ ARMMySQL = (*ARMMySQLMock)(nil)

func (m *ARMMySQLMock) Get(_ context.Context, group, name string, _ *armmysql.ServersClientGetOptions) (armmysql.ServersClientGetResponse, error) {
	if m.NoAuth {
		return armmysql.ServersClientGetResponse{}, trace.AccessDenied("unauthorized")
	}
	for _, s := range m.DBServers {
		if name == *s.Name {
			id, err := arm.ParseResourceID(*s.ID)
			if err != nil {
				return armmysql.ServersClientGetResponse{}, trace.Wrap(err)
			}
			if group == id.ResourceGroupName {
				return armmysql.ServersClientGetResponse{Server: *s}, nil
			}
		}
	}
	return armmysql.ServersClientGetResponse{}, trace.NotFound("resource %v in group %v not found", name, group)
}

func (m *ARMMySQLMock) NewListPager(_ *armmysql.ServersClientListOptions) *runtime.Pager[armmysql.ServersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[armmysql.ServersClientListResponse]{
		More: func(_ armmysql.ServersClientListResponse) bool {
			return false
		},
		Fetcher: func(_ context.Context, _ *armmysql.ServersClientListResponse) (armmysql.ServersClientListResponse, error) {
			if m.NoAuth {
				return armmysql.ServersClientListResponse{}, trace.AccessDenied("unauthorized")
			}
			return armmysql.ServersClientListResponse{
				ServerListResult: armmysql.ServerListResult{
					Value: m.DBServers,
				},
			}, nil
		},
	})
}

func (m *ARMMySQLMock) NewListByResourceGroupPager(group string, _ *armmysql.ServersClientListByResourceGroupOptions) *runtime.Pager[armmysql.ServersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[armmysql.ServersClientListByResourceGroupResponse]{
		More: func(_ armmysql.ServersClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(_ context.Context, _ *armmysql.ServersClientListByResourceGroupResponse) (armmysql.ServersClientListByResourceGroupResponse, error) {
			if m.NoAuth {
				return armmysql.ServersClientListByResourceGroupResponse{}, trace.AccessDenied("unauthorized")
			}
			var servers []*armmysql.Server
			for _, s := range m.DBServers {
				id, err := arm.ParseResourceID(*s.ID)
				if err != nil {
					return armmysql.ServersClientListByResourceGroupResponse{}, trace.Wrap(err)
				}
				if group == id.ResourceGroupName {
					servers = append(servers, s)
				}
			}
			if len(servers) == 0 {
				return armmysql.ServersClientListByResourceGroupResponse{}, trace.NotFound("Resource group '%v' could not be found.", group)
			}
			return armmysql.ServersClientListByResourceGroupResponse{
				ServerListResult: armmysql.ServerListResult{
					Value: servers,
				},
			}, nil
		},
	})
}

// ARMPostgresMock mocks Azure armpostgresql API.
type ARMPostgresMock struct {
	DBServers []*armpostgresql.Server
	NoAuth    bool
}

var _ ARMPostgres = (*ARMPostgresMock)(nil)

func (m *ARMPostgresMock) Get(_ context.Context, group, name string, _ *armpostgresql.ServersClientGetOptions) (armpostgresql.ServersClientGetResponse, error) {
	if m.NoAuth {
		return armpostgresql.ServersClientGetResponse{}, trace.AccessDenied("unauthorized")
	}
	for _, s := range m.DBServers {
		if name == *s.Name {
			id, err := arm.ParseResourceID(*s.ID)
			if err != nil {
				return armpostgresql.ServersClientGetResponse{}, trace.Wrap(err)
			}
			if group == id.ResourceGroupName {
				return armpostgresql.ServersClientGetResponse{Server: *s}, nil
			}
		}
	}
	return armpostgresql.ServersClientGetResponse{}, trace.NotFound("resource %v in group %v not found", name, group)
}

func (m *ARMPostgresMock) NewListPager(_ *armpostgresql.ServersClientListOptions) *runtime.Pager[armpostgresql.ServersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[armpostgresql.ServersClientListResponse]{
		More: func(_ armpostgresql.ServersClientListResponse) bool {
			return false
		},
		Fetcher: func(_ context.Context, _ *armpostgresql.ServersClientListResponse) (armpostgresql.ServersClientListResponse, error) {
			if m.NoAuth {
				return armpostgresql.ServersClientListResponse{}, trace.AccessDenied("unauthorized")
			}
			return armpostgresql.ServersClientListResponse{
				ServerListResult: armpostgresql.ServerListResult{
					Value: m.DBServers,
				},
			}, nil
		},
	})
}

func (m *ARMPostgresMock) NewListByResourceGroupPager(group string, _ *armpostgresql.ServersClientListByResourceGroupOptions) *runtime.Pager[armpostgresql.ServersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[armpostgresql.ServersClientListByResourceGroupResponse]{
		More: func(_ armpostgresql.ServersClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(_ context.Context, _ *armpostgresql.ServersClientListByResourceGroupResponse) (armpostgresql.ServersClientListByResourceGroupResponse, error) {
			if m.NoAuth {
				return armpostgresql.ServersClientListByResourceGroupResponse{}, trace.AccessDenied("unauthorized")
			}
			var servers []*armpostgresql.Server
			for _, s := range m.DBServers {
				id, err := arm.ParseResourceID(*s.ID)
				if err != nil {
					return armpostgresql.ServersClientListByResourceGroupResponse{}, trace.Wrap(err)
				}
				if group == id.ResourceGroupName {
					servers = append(servers, s)
				}
			}
			if len(servers) == 0 {
				return armpostgresql.ServersClientListByResourceGroupResponse{}, trace.NotFound("Resource group '%v' could not be found.", group)
			}
			return armpostgresql.ServersClientListByResourceGroupResponse{
				ServerListResult: armpostgresql.ServerListResult{
					Value: servers,
				},
			}, nil
		},
	})
}

// ARMKubernetesMock mocks Azure armmanagedclusters API.
type ARMKubernetesMock struct {
	KubeServers       []*armcontainerservice.ManagedCluster
	ClusterAdminCreds *armcontainerservice.CredentialResult
	ClusterUserCreds  *armcontainerservice.CredentialResult
	NoAuth            bool
}

var _ ARMAKS = (*ARMKubernetesMock)(nil)

func (m *ARMKubernetesMock) Get(_ context.Context, group, name string, _ *armcontainerservice.ManagedClustersClientGetOptions) (armcontainerservice.ManagedClustersClientGetResponse, error) {
	if m.NoAuth {
		return armcontainerservice.ManagedClustersClientGetResponse{}, trace.AccessDenied("unauthorized")
	}
	for _, s := range m.KubeServers {
		if name == *s.Name {
			id, err := arm.ParseResourceID(*s.ID)
			if err != nil {
				return armcontainerservice.ManagedClustersClientGetResponse{}, trace.Wrap(err)
			}
			if group == id.ResourceGroupName {
				return armcontainerservice.ManagedClustersClientGetResponse{ManagedCluster: *s}, nil
			}
		}
	}
	return armcontainerservice.ManagedClustersClientGetResponse{}, trace.NotFound("resource %v in group %v not found", name, group)
}

func (m *ARMKubernetesMock) NewListPager(_ *armcontainerservice.ManagedClustersClientListOptions) *runtime.Pager[armcontainerservice.ManagedClustersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[armcontainerservice.ManagedClustersClientListResponse]{
		More: func(_ armcontainerservice.ManagedClustersClientListResponse) bool {
			return false
		},
		Fetcher: func(_ context.Context, _ *armcontainerservice.ManagedClustersClientListResponse) (armcontainerservice.ManagedClustersClientListResponse, error) {
			if m.NoAuth {
				return armcontainerservice.ManagedClustersClientListResponse{}, trace.AccessDenied("unauthorized")
			}
			return armcontainerservice.ManagedClustersClientListResponse{
				ManagedClusterListResult: armcontainerservice.ManagedClusterListResult{
					Value: m.KubeServers,
				},
			}, nil
		},
	})
}

func (m *ARMKubernetesMock) NewListByResourceGroupPager(group string, _ *armcontainerservice.ManagedClustersClientListByResourceGroupOptions) *runtime.Pager[armcontainerservice.ManagedClustersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[armcontainerservice.ManagedClustersClientListByResourceGroupResponse]{
		More: func(_ armcontainerservice.ManagedClustersClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(_ context.Context, _ *armcontainerservice.ManagedClustersClientListByResourceGroupResponse) (armcontainerservice.ManagedClustersClientListByResourceGroupResponse, error) {
			if m.NoAuth {
				return armcontainerservice.ManagedClustersClientListByResourceGroupResponse{}, trace.AccessDenied("unauthorized")
			}
			var servers []*armcontainerservice.ManagedCluster
			for _, s := range m.KubeServers {
				id, err := arm.ParseResourceID(*s.ID)
				if err != nil {
					return armcontainerservice.ManagedClustersClientListByResourceGroupResponse{}, trace.Wrap(err)
				}
				if group == id.ResourceGroupName {
					servers = append(servers, s)
				}
			}
			if len(servers) == 0 {
				return armcontainerservice.ManagedClustersClientListByResourceGroupResponse{}, trace.NotFound("Resource group '%v' could not be found.", group)
			}
			return armcontainerservice.ManagedClustersClientListByResourceGroupResponse{
				ManagedClusterListResult: armcontainerservice.ManagedClusterListResult{
					Value: servers,
				},
			}, nil
		},
	})
}

func (m *ARMKubernetesMock) GetCommandResult(ctx context.Context, resourceGroupName string, resourceName string, commandID string, options *armcontainerservice.ManagedClustersClientGetCommandResultOptions) (armcontainerservice.ManagedClustersClientGetCommandResultResponse, error) {
	return armcontainerservice.ManagedClustersClientGetCommandResultResponse{
		RunCommandResult: armcontainerservice.RunCommandResult{
			ID: valToPtr(commandID),
		},
	}, nil
}
func (m *ARMKubernetesMock) ListClusterAdminCredentials(ctx context.Context, resourceGroupName string, resourceName string, options *armcontainerservice.ManagedClustersClientListClusterAdminCredentialsOptions) (armcontainerservice.ManagedClustersClientListClusterAdminCredentialsResponse, error) {
	if m.NoAuth {
		return armcontainerservice.ManagedClustersClientListClusterAdminCredentialsResponse{}, trace.AccessDenied("unauthorized")
	}

	return armcontainerservice.ManagedClustersClientListClusterAdminCredentialsResponse{
		CredentialResults: armcontainerservice.CredentialResults{
			Kubeconfigs: []*armcontainerservice.CredentialResult{
				m.ClusterAdminCreds,
			},
		},
	}, nil
}
func (m *ARMKubernetesMock) ListClusterUserCredentials(ctx context.Context, resourceGroupName string, resourceName string, options *armcontainerservice.ManagedClustersClientListClusterUserCredentialsOptions) (armcontainerservice.ManagedClustersClientListClusterUserCredentialsResponse, error) {
	if m.NoAuth {
		return armcontainerservice.ManagedClustersClientListClusterUserCredentialsResponse{}, trace.AccessDenied("unauthorized")
	}
	return armcontainerservice.ManagedClustersClientListClusterUserCredentialsResponse{
		CredentialResults: armcontainerservice.CredentialResults{
			Kubeconfigs: []*armcontainerservice.CredentialResult{
				m.ClusterUserCreds,
			},
		},
	}, nil
}

func (m *ARMKubernetesMock) BeginRunCommand(ctx context.Context, resourceGroupName string, resourceName string, requestPayload armcontainerservice.RunCommandRequest, options *armcontainerservice.ManagedClustersClientBeginRunCommandOptions) (*runtime.Poller[armcontainerservice.ManagedClustersClientRunCommandResponse], error) {
	if m.NoAuth {
		return nil, trace.AccessDenied("unauthorized")
	}
	return &runtime.Poller[armcontainerservice.ManagedClustersClientRunCommandResponse]{}, nil
}
