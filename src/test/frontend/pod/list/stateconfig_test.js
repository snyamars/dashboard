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

import {resolvePodList} from 'pod/list/stateconfig';
import podListModule from 'pod/module';

describe('StateConfig for pod list', () => {
  /** @type {!common/dataselect/dataselect_service.DataSelectService} */
  let kdDataSelectService;

  beforeEach(() => {
    angular.mock.module(podListModule.name);
    angular.mock.inject((_kdDataSelectService_) => {
      kdDataSelectService = _kdDataSelectService_;
    });
  });

  it('should resolve pods with namespace', angular.mock.inject(($q) => {
    let promise = $q.defer().promise;

    let resource = jasmine.createSpyObj('$resource', ['get']);
    resource.get.and.callFake(function() {
      return {$promise: promise};
    });

    let actual = resolvePodList(resource, {namespace: 'foo'}, kdDataSelectService);

    expect(resource.get).toHaveBeenCalledWith(kdDataSelectService.getDefaultResourceQuery('foo'));
    expect(actual).toBe(promise);
  }));

  it('should resolve pods with no namespace', angular.mock.inject(($q) => {
    let promise = $q.defer().promise;

    let resource = jasmine.createSpyObj('$resource', ['get']);
    resource.get.and.callFake(function() {
      return {$promise: promise};
    });

    let actual = resolvePodList(resource, {}, kdDataSelectService);

    expect(resource.get).toHaveBeenCalledWith(kdDataSelectService.getDefaultResourceQuery(''));
    expect(actual).toBe(promise);
  }));
});
