// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"reflect"
	"testing"

	"github.com/kubernetes/dashboard/src/app/backend/api"
	"github.com/kubernetes/dashboard/src/app/backend/resource/common"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/pkg/api/v1"
)

func TestToServiceDetail(t *testing.T) {
	cases := []struct {
		service  *v1.Service
		expected ServiceDetail
	}{
		{
			service: &v1.Service{}, expected: ServiceDetail{
				TypeMeta: api.TypeMeta{Kind: api.ResourceKindService},
			},
		}, {
			service: &v1.Service{
				ObjectMeta: metaV1.ObjectMeta{
					Name: "test-service", Namespace: "test-namespace",
				}},
			expected: ServiceDetail{
				ObjectMeta: api.ObjectMeta{
					Name:      "test-service",
					Namespace: "test-namespace",
				},
				TypeMeta:         api.TypeMeta{Kind: api.ResourceKindService},
				InternalEndpoint: common.Endpoint{Host: "test-service.test-namespace"},
			},
		},
	}

	for _, c := range cases {
		actual := ToServiceDetail(c.service)

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("ToServiceDetail(%#v) == \ngot %#v, \nexpected %#v", c.service, actual,
				c.expected)
		}
	}
}

func TestToService(t *testing.T) {
	cases := []struct {
		service  *v1.Service
		expected Service
	}{
		{
			service: &v1.Service{}, expected: Service{
				TypeMeta: api.TypeMeta{Kind: api.ResourceKindService},
			},
		}, {
			service: &v1.Service{
				ObjectMeta: metaV1.ObjectMeta{
					Name: "test-service", Namespace: "test-namespace",
				}},
			expected: Service{
				ObjectMeta: api.ObjectMeta{
					Name:      "test-service",
					Namespace: "test-namespace",
				},
				TypeMeta:         api.TypeMeta{Kind: api.ResourceKindService},
				InternalEndpoint: common.Endpoint{Host: "test-service.test-namespace"},
			},
		},
	}

	for _, c := range cases {
		actual := ToService(c.service)

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("ToService(%#v) == \ngot %#v, \nexpected %#v", c.service, actual,
				c.expected)
		}
	}
}
