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

package mock

import (
	"encoding/pem"
	"fmt"
	"sync/atomic"
	"time"

	"istio.io/istio/security/pkg/pki/util"
)

// CAClient is the mocked CAClient for testing.
type CAClient struct {
	SignInvokeCount uint64
	bundle          *util.KeyCertBundle
	certLifetime    time.Duration
	GeneratedCerts  [][]string // Cache the generated certificates for verification purpose.
	mockTrustAnchor bool
}

// NewMockCAClient creates an instance of CAClient. errors is used to specify the number of errors
// before CSRSign returns a valid response. certLifetime specifies the TTL for the newly issued workload cert.

func (c *CAClient) Close() {
	// CSRSign returns the certificate or errors depending on the settings.
}

func (c *CAClient) CSRSign(csrPEM []byte, certValidTTLInSec int64) ([]string, error) {
	atomic.AddUint64(&c.SignInvokeCount, 1)
	signingCert, signingKey, certChain, rootCert := c.bundle.GetAll()
	csr, err := util.ParsePemEncodedCSR(csrPEM)
	if err != nil {
		return nil, fmt.Errorf("csr sign error: %v", err)
	}
	subjectIDs := []string{"test"}
	certBytes, err := util.GenCertFromCSR(csr, signingCert, csr.PublicKey, *signingKey, subjectIDs, c.certLifetime, false)
	if err != nil {
		return nil, fmt.Errorf("csr sign error: %v", err)
	}

	block := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	}
	cert := pem.EncodeToMemory(block)

	ret := []string{string(cert), string(certChain), string(rootCert)}
	c.GeneratedCerts = append(c.GeneratedCerts, ret)
	return ret, nil
}

func (c *CAClient) GetRootCertBundle() ([]string, error) {
	if c.mockTrustAnchor {
		rootCertBytes := c.bundle.GetRootCertPem()
		return []string{string(rootCertBytes)}, nil
	}

	return []string{}, nil
}
