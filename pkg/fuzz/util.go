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

package fuzz

import (
	"bytes"
	"strings"

	fuzzheaders "github.com/AdaLogics/go-fuzz-headers"

	"istio.io/istio/pkg/test"
)

const panicPrefix = "go-fuzz-skip: "

// Helper is a helper struct for fuzzing
type Helper struct {
	cf *fuzzheaders.ConsumeFuzzer
	t  test.Failer
}

type Validator interface {
	// FuzzValidate returns true if the current struct is valid for fuzzing.
	FuzzValidate() bool
}

// Fuzz is a wrapper around:
//
//	 fuzz.BaseCases(f)
//		f.Fuzz(func(...) {
//		   defer fuzz.Finalizer()
//		}
//
// To avoid needing to call BaseCases and Finalize everywhere.

// Finalize works around an issue in the oss-fuzz logic that doesn't allow using Skip()
// Instead, we send a panic which we handle and treat as skip.
// https://github.com/AdamKorcz/go-118-fuzz-build/issues/6
func Finalize() {
	if r := recover(); r != nil {
		if s, ok := r.(string); ok {
			if strings.HasPrefix(s, panicPrefix) {
				return
			}
		}
		panic(r)
	}
}

// New creates a new fuzz.Helper, capable of generating more complex types
func New(t test.Failer, data []byte) Helper {
	return Helper{cf: fuzzheaders.NewConsumer(data), t: t}
}

// Struct generates a Struct. Validation patterns can be passed in - if any return false, the fuzz case is skipped.
// Additionally, if the T implements Validator, it will implicitly be used.

// Slice generates a slice of Structs

// BaseCases inserts a few trivial test cases to do a very brief sanity check of a test that relies on []byte inputs
func BaseCases(f test.Fuzzer) {
	for _, c := range [][]byte{
		{},
		[]byte("."),
		bytes.Repeat([]byte("."), 1000),
	} {
		f.Add(c)
	}
}

// T Returns the underlying test.Failer. Should be avoided where possible; in oss-fuzz many functions do not work.
func (h Helper) T() test.Failer {
	return h.t
}

type mutateCtx struct {
	t        test.Failer
	curDepth int
	maxDepth int
}

// MutateStruct modify the field value of the structure.
// It is mainly used to check the correctness of the deep copy.

// DeepCopySlow is a general deep copy method that guarantees the correctness of deep copying,
// but may be very slow. Here, it is only used for testing.
// Note: this function not support copy struct private filed.
