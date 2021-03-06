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

package thirdpartyresource

import (
	"reflect"
	"testing"

	"github.com/kubernetes/dashboard/src/app/backend/api"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	extensions "k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

func TestGetThirdPartyResourceDetail(t *testing.T) {
	cases := []struct {
		thirdPartyResource *extensions.ThirdPartyResource
		expected           *ThirdPartyResourceDetail
	}{
		{
			&extensions.ThirdPartyResource{
				ObjectMeta:  metaV1.ObjectMeta{Name: "foo"},
				Description: "test",
				Versions: []extensions.APIVersion{
					{
						Name: "v1",
					},
				},
			},
			&ThirdPartyResourceDetail{
				TypeMeta:    api.TypeMeta{Kind: api.ResourceKindThirdPartyResource},
				ObjectMeta:  api.ObjectMeta{Name: "foo"},
				Description: "test",
				Versions: []extensions.APIVersion{
					{
						Name: "v1",
					},
				},
			},
		},
	}
	for _, c := range cases {
		actual := getThirdPartyResourceDetail(c.thirdPartyResource, ThirdPartyResourceObjectList{})
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("getThirdPartyResourceDetail(%#v) == \n%#v\nexpected \n%#v\n",
				c.thirdPartyResource, actual, c.expected)
		}
	}
}
