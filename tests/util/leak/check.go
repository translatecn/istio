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

// leak checks for goroutine leaks in tests
// This is (heavily) inspired by https://github.com/grpc/grpc-go/blob/master/internal/leakcheck/leakcheck.go
// and https://github.com/fortytw2/leaktest
package leak

import (
	"time"
)

var goroutinesToIgnore = []string{
	// "global" goroutines we always initialize. Maybe we shouldn't always initialize these, but for now every
	// test fails with these
	"k8s.io/klog/v2.(*loggingT).flushDaemon",      // k8s logging
	"go.opencensus.io/stats/view.(*worker).start", // metrics runs on init. We are *almost* off opencensus, but transitively import it.

	// goroutines for test
	"testing.Main(",
	"testing.tRunner(",
	"testing.(*M).",

	// go runtime
	"runtime.goexit",
	"created by runtime.gc",
	"runtime.MHeap_Scavenger",
	"signal.signal_recv",
	"sigterm.handler",
	"runtime_mcall",

	// created by leak checker
	"created by runtime/trace.Start",
	"interestingGoroutines",

	// This is not technically required. However, its a loop that is outside our control that runs every 500ms
	// so we skip it to avoid delayed tests
	"workqueue.(*Type).updateUnfinishedWorkLoop",
}

// TestingM is the minimal subset of testing.M that we use.
type TestingM interface {
	Run() int
}

type TestingTB interface {
	Cleanup(func())
	Errorf(format string, args ...any)
}

var gracePeriod = time.Second * 5

// Check adds a check to a test to ensure there are no leaked goroutines
// To use, simply call leak.Check(t) at the start of a test; Do not call it in defer.
// It is recommended to call this as the first step, as Cleanup is called in LIFO order; this ensures any
// Cleanup's called in the test happen first.
// Any existing goroutines before the test starts are filtered out. This ensures a single test failing doesn't
// cause all future tests to fail. However, it is still possible another test influences the result when t.Parallel is used.
// Where possible, CheckMain is preferred.

// CheckMain asserts that no goroutines are leaked after a test package exits.
// This can be used with the following code:
//
//	func TestMain(m *testing.M) {
//	    leak.CheckMain(m)
//	}
//
// Failures here are scoped to the package, not a specific test. To determine the source of the failure,
// you can use the tool `go test -exec $PWD/tools/go-ordered-test ./my/package`. This runs each test individually.
// If there are some tests that are leaky, you the Check method can be used on individual tests.

// MustGarbageCollect asserts that an object was garbage collected by the end of the test.
// The input must be a pointer to an object.

type goroutine struct {
	id    uint64
	stack string
}

type goroutineByID []*goroutine

func (g goroutineByID) Len() int           { return len(g) }
func (g goroutineByID) Less(i, j int) bool { return g[i].id < g[j].id }
func (g goroutineByID) Swap(i, j int)      { g[i], g[j] = g[j], g[i] }

// interestingGoroutines returns all goroutines we care about for the purpose
// of leak checking. It excludes testing or runtime ones.
