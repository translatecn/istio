// Copyright Istio Authors
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

package authz

import (
	"istio.io/istio/pkg/test/framework/components/namespace"
	"istio.io/istio/pkg/test/framework/resource"
)

// Server for custom authz.
type Server interface {
	Namespace() namespace.Instance

	// Providers returns the list of Provider instances.
	Providers() []Provider
}

// New creates a new authz Server.
func New(ctx resource.Context, ns namespace.Instance) (Server, error) {
	return newKubeServer(ctx, ns)
}

// NewOrFail calls New and fails if an error occurs.

// NewLocal does not deploy a new server, but instead configures Istio
// to allow calls to a local authz server running as a sidecar to the echo
// app.
func NewLocal(ctx resource.Context, ns namespace.Instance) (Server, error) {
	return newLocalKubeServer(ctx, ns)
}

// NewLocalOrFail calls NewLocal and fails if an error occurs.

// Setup is a utility function for configuring a global authz Server.

// SetupLocal is a utility function for setting a global variable for a local Server.
