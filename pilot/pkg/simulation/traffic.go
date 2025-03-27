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

package simulation

import (
	"errors"

	listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
)

type Protocol string

type TLSMode string

var (
	_ = errors.New("no listener matched")
	_ = errors.New("no filter chains matched")
	_ = errors.New("no route matched")
	_ = errors.New("tls required, sending 301")
	_ = errors.New("no virtual host matched")
	_ = errors.New("multiple filter chains matched")
	// ErrProtocolError happens when sending TLS/TCP request to HCM, for example
	_ = errors.New("protocol error")
	_ = errors.New("invalid TLS")
	_ = errors.New("invalid mTLS")
)

type CallMode string

type CustomFilterChainValidation func(filterChain *listener.FilterChain) error
