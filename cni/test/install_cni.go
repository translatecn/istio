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

package install

const (
	cniConfSubDir    = "/testdata/pre/"
	k8sSvcAcctSubDir = "/testdata/k8s_svcacct/"

	defaultFileMode = 0o644
)

// Removes Istio CNI's config from the CNI config file

// populateTempDirs populates temporary test directories with golden files and
// other related configuration.

// create an install server instance and run it, blocking until it gets terminated
// via context cancellation

// checkResult checks if resultFile is equal to expectedFile at each tick until timeout

// compareConfResult does a string compare of 2 test files.

// checkBinDir verifies the presence/absence of test files.

// checkTempFilesCleaned verifies that all temporary files have been cleaned up

// doTest sets up necessary environment variables, runs the Docker installation
// container and verifies output file correctness.

// RunInstallCNITest sets up temporary directories and runs the test.
//
// Doing a go test install_cni.go by itself will not execute the test as the
// file doesn't have a _test.go suffix, and this func doesn't start with a Test
// prefix. This func is only meant to be invoked programmatically. A separate
// install_cni_test.go file exists for executing this test.
