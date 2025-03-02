/*
Copyright 2019 The Kubernetes Authors.

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

package controllerfetcher

import "context"

// FakeControllerFetcher should be used in test only. It returns exactly the same controllerKey
type FakeControllerFetcher struct{}

// FindTopMostWellKnownOrScalable returns the same key for that fake implementation and returns and error when the kind is Node
// See pkg/target/controller_fetcher/controller_fetcher.go:296 where the original implementation does the same.
func (f FakeControllerFetcher) FindTopMostWellKnownOrScalable(_ context.Context, controller *ControllerKeyWithAPIVersion) (*ControllerKeyWithAPIVersion, error) {
	if controller.Kind == "Node" {
		return nil, ErrNodeInvalidOwner
	}
	return controller, nil
}

// NilControllerFetcher is a fake ControllerFetcher which always returns 'nil'
type NilControllerFetcher struct{}

// FindTopMostWellKnownOrScalable always returns nil
func (f NilControllerFetcher) FindTopMostWellKnownOrScalable(_ context.Context, _ *ControllerKeyWithAPIVersion) (*ControllerKeyWithAPIVersion, error) {
	return nil, nil
}

var _ ControllerFetcher = &FakeControllerFetcher{}
