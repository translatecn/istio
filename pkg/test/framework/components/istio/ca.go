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

package istio

import (
	"context"
	"fmt"
	"sync"
	"time"

	authenticationv1 "k8s.io/api/authentication/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Cert struct {
	ClientCert, Key, RootCert []byte
}

// 7 days
var saTokenExpiration int64 = 60 * 60 * 24 * 7

func GetServiceAccountToken(c kubernetes.Interface, aud, ns, sa string) (string, error) {
	san := san(ns, sa)

	if got, f := cachedTokens.Load(san); f {
		t := got.(token)
		if t.expiration.After(time.Now().Add(time.Minute)) {
			return t.token, nil
		}
		// Otherwise, its expired, load a new one
	}
	rt, err := c.CoreV1().ServiceAccounts(ns).CreateToken(context.Background(), sa,
		&authenticationv1.TokenRequest{
			Spec: authenticationv1.TokenRequestSpec{
				Audiences:         []string{aud},
				ExpirationSeconds: &saTokenExpiration,
			},
		}, metav1.CreateOptions{})
	if err != nil {
		return "", err
	}
	exp := rt.Status.ExpirationTimestamp.Time
	cachedTokens.Store(san, token{rt.Status.Token, exp})
	return rt.Status.Token, nil
}

// map of SAN to jwt token. Used to avoid repetitive calls
var cachedTokens sync.Map

type token struct {
	token      string
	expiration time.Time
}

// NewCitadelClient create a CA client for Citadel.

func san(ns, sa string) string {
	return fmt.Sprintf("spiffe://%s/ns/%s/sa/%s", "cluster.local", ns, sa)
}

func FetchRootCert(c kubernetes.Interface) (string, error) {
	cm, err := c.CoreV1().ConfigMaps("istio-system").Get(context.TODO(), "istio-ca-root-cert", metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return cm.Data["root-cert.pem"], nil
}
