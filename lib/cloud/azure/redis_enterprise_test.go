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
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise"
	"github.com/stretchr/testify/require"
)

func TestRedisEnterpriseClient(t *testing.T) {
	mockClusterAPI := &ARMRedisEnterpriseClusterMock{
		Clusters: []*armredisenterprise.Cluster{
			makeRedisEnterpriceCluster("redis-prod-1", "group-prod"),
			makeRedisEnterpriceCluster("redis-prod-2", "group-prod"),
			makeRedisEnterpriceCluster("redis-dev", "group-dev"),
		},
	}
	mockDatabaseAPI := &ARMRedisEnterpriseDatabaseMock{
		TokensByDatabaseName: map[string]string{
			"default":       "default-token",
			"some-database": "some-database-token",
		},
		Databases: []*armredisenterprise.Database{
			makeRedisEnterpriceDatabase("default", "redis-prod-1", "group-prod"),
			makeRedisEnterpriceDatabase("db-x", "redis-prod-2", "group-prod"),
			makeRedisEnterpriceDatabase("db-y", "redis-prod-2", "group-prod"),
			makeRedisEnterpriceDatabase("default", "redis-dev", "group-dev"),
		},
	}

	mockClusterAPINoAuth := &ARMRedisEnterpriseClusterMock{
		NoAuth: true,
	}
	mockDatabaseAPINoAuth := &ARMRedisEnterpriseDatabaseMock{
		NoAuth: true,
	}

	t.Run("GetToken", func(t *testing.T) {
		tests := []struct {
			name            string
			mockDatabaseAPI armRedisEnterpriseDatabaseClient
			resourceID      string
			expectError     bool
			expectToken     string
		}{
			{
				name:            "access denied",
				resourceID:      "cluster-name",
				mockDatabaseAPI: mockDatabaseAPINoAuth,
				expectError:     true,
			},
			{
				name:            "succeed (default database name)",
				resourceID:      "/subscriptions/sub-id/resourceGroups/group-name/providers/Microsoft.Cache/redisEnterprise/example-teleport",
				mockDatabaseAPI: mockDatabaseAPI,
				expectToken:     "default-token",
			},
			{
				name:            "succeed (specific database name)",
				resourceID:      "/subscriptions/sub-id/resourceGroups/group-name/providers/Microsoft.Cache/redisEnterprise/example-teleport/databases/some-database",
				mockDatabaseAPI: mockDatabaseAPI,
				expectToken:     "some-database-token",
			},
		}

		for _, test := range tests {
			test := test
			t.Run(test.name, func(t *testing.T) {
				t.Parallel()

				c := NewRedisEnterpriseClientByAPI(mockClusterAPI, test.mockDatabaseAPI)
				token, err := c.GetToken(context.TODO(), test.resourceID)

				if test.expectError {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					require.Equal(t, test.expectToken, token)
				}
			})
		}
	})

	t.Run("ListALL", func(t *testing.T) {
		tests := []struct {
			name                   string
			mockDatabaseAPI        armRedisEnterpriseDatabaseClient
			mockClusterAPI         armRedisEnterpriseClusterClient
			expectError            bool
			expectClusterDatabases map[string][]string
		}{
			{
				name:            "access denied",
				mockDatabaseAPI: mockDatabaseAPINoAuth,
				mockClusterAPI:  mockClusterAPINoAuth,
				expectError:     true,
			},
			{
				name:            "succeed",
				mockDatabaseAPI: mockDatabaseAPI,
				mockClusterAPI:  mockClusterAPI,
				expectClusterDatabases: map[string][]string{
					"redis-prod-1": []string{"default"},
					"redis-prod-2": []string{"db-x", "db-y"},
					"redis-dev":    []string{"default"},
				},
			},
		}

		for _, test := range tests {
			test := test
			t.Run(test.name, func(t *testing.T) {
				t.Parallel()

				c := NewRedisEnterpriseClientByAPI(test.mockClusterAPI, test.mockDatabaseAPI)
				resources, err := c.ListAll(context.TODO())
				if test.expectError {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					require.Len(t, resources, 3)
				}
			})
		}
	})
}

func makeRedisEnterpriceCluster(name, group string) *armredisenterprise.Cluster {
	return &armredisenterprise.Cluster{
		Name:     to.Ptr(name),
		ID:       to.Ptr(fmt.Sprintf("/subscriptions/sub-id/resourceGroups/%v/providers/Microsoft.Cache/redisEnterprise/%v", group, name)),
		Type:     to.Ptr("Microsoft.Cache/redisEnterprise"),
		Location: to.Ptr("local"),
	}
}

func makeRedisEnterpriceDatabase(name, clusterName, group string) *armredisenterprise.Database {
	return &armredisenterprise.Database{
		Name: to.Ptr(name),
		ID:   to.Ptr(fmt.Sprintf("/subscriptions/sub-id/resourceGroups/%v/providers/Microsoft.Cache/redisEnterprise/%v/databases/%v", group, clusterName, name)),
	}
}
