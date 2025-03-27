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

package assert

import (
	"reflect"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/testing/protocmp"

	"istio.io/istio/pkg/test"
)

var compareErrors = cmp.Comparer(func(x, y error) bool {
	switch {
	case x == nil && y == nil:
		return true
	case x != nil && y == nil:
		return false
	case x == nil && y != nil:
		return false
	case x != nil && y != nil:
		return x.Error() == y.Error()
	default:
		panic("unreachable")
	}
})

// cmpOptioner can be implemented to provide custom options that should be used when comparing a type.
// Warning: this is no recursive, unfortunately. So a type `A{B}` cannot rely on `B` implementing this to customize comparing `B`.
type cmpOptioner interface {
	CmpOpts() []cmp.Option
}

// opts gets the comparison opts for a type. This includes some defaults, but allows each type to explicitly append their own.
func opts[T any](a T) []cmp.Option {
	if o, ok := any(a).(cmpOptioner); ok {
		opts := append([]cmp.Option{}, cmpOpts...)
		opts = append(opts, o.CmpOpts()...)
		return opts
	}
	// if T is actually a slice (ex: []A), check that and get the opts for the element type (A).
	t := reflect.TypeOf(a)
	if t != nil && t.Kind() == reflect.Slice {
		v := reflect.New(t.Elem()).Elem().Interface()
		if o, ok := v.(cmpOptioner); ok {
			opts := append([]cmp.Option{}, cmpOpts...)
			opts = append(opts, o.CmpOpts()...)
			return opts
		}
	}
	return cmpOpts
}

var cmpOpts = []cmp.Option{protocmp.Transform(), cmpopts.EquateEmpty(), compareErrors}

// Compare compares two objects and returns and error if they are not the same.

// Equal compares two objects and fails if they are not the same.
func Equal[T any](t test.Failer, a, b T, context ...string) {
	t.Helper()
	if !cmp.Equal(a, b, opts(a)...) {
		cs := ""
		if len(context) > 0 {
			cs = " " + strings.Join(context, ", ") + ":"
		}
		t.Fatalf("found diff:%s %v\nLeft:  %v\nRight: %v", cs, cmp.Diff(a, b, opts(a)...), a, b)
	}
}

// EventuallyEqual compares repeatedly calls the fetch function until the result matches the expectation.

// Error asserts the provided err is non-nil

// NoError asserts the provided err is nil
func NoError(t test.Failer, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected no error but got: %v", err)
	}
}

// ChannelHasItem asserts a channel has an element within 5s and returns the element

// ChannelIsEmpty asserts a channel is empty for at least 20ms
