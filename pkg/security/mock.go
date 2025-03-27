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

package security

import (
	"fmt"
	"sync"
)

type DirectSecretManager struct {
	items map[string]*SecretItem
	mu    sync.RWMutex
}

var _ SecretManager = &DirectSecretManager{}

func (d *DirectSecretManager) GenerateSecret(resourceName string) (*SecretItem, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	si, f := d.items[resourceName]
	if !f {
		return nil, fmt.Errorf("resource %v not found", resourceName)
	}
	return si, nil
}

func (d *DirectSecretManager) Set(resourceName string, secret *SecretItem) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if secret == nil {
		delete(d.items, resourceName)
	} else {
		d.items[resourceName] = secret
	}
}
