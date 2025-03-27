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

package jwt

import (
	"istio.io/istio/pkg/test/framework/components/namespace"
	"istio.io/istio/pkg/test/framework/resource"
)

// Server for JWT tokens.
type Server interface {
	Namespace() namespace.Instance
	FQDN() string
	HTTPPort() int
	HTTPSPort() int
	JwksURI() string
}

// New creates a new JWT Server.
func New(ctx resource.Context, ns namespace.Instance) (Server, error) {
	return newKubeServer(ctx, ns)
}

// NewOrFail calls New and fails if an error occurs.

// Setup is a utility function for configuring a jwt Server.
