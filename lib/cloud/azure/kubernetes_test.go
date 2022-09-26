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
	"reflect"
	"testing"
	"time"

	"k8s.io/client-go/rest"
)

func Test_aKSClient_ClusterCredentials(t *testing.T) {
	type fields struct {
		api        ARMAKS
		azIdentity AzureIdentityFunction
	}
	type args struct {
		ctx context.Context
		cfg ClusterCredentialsConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *rest.Config
		want1   time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewAKSClustersClient(tt.fields.api, tt.fields.azIdentity)
			got, got1, err := c.ClusterCredentials(tt.args.ctx, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("aKSClient.ClusterCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("aKSClient.ClusterCredentials() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("aKSClient.ClusterCredentials() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
