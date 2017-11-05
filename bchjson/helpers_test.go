// Copyright (c) 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package bchjson_test

import (
	"reflect"
	"testing"

	"github.com/bchsuite/bchd/bchjson"
)

// TestHelpers tests the various helper functions which create pointers to
// primitive types.
func TestHelpers(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		f        func() interface{}
		expected interface{}
	}{
		{
			name: "bool",
			f: func() interface{} {
				return bchjson.Bool(true)
			},
			expected: func() interface{} {
				val := true
				return &val
			}(),
		},
		{
			name: "int",
			f: func() interface{} {
				return bchjson.Int(5)
			},
			expected: func() interface{} {
				val := int(5)
				return &val
			}(),
		},
		{
			name: "uint",
			f: func() interface{} {
				return bchjson.Uint(5)
			},
			expected: func() interface{} {
				val := uint(5)
				return &val
			}(),
		},
		{
			name: "int32",
			f: func() interface{} {
				return bchjson.Int32(5)
			},
			expected: func() interface{} {
				val := int32(5)
				return &val
			}(),
		},
		{
			name: "uint32",
			f: func() interface{} {
				return bchjson.Uint32(5)
			},
			expected: func() interface{} {
				val := uint32(5)
				return &val
			}(),
		},
		{
			name: "int64",
			f: func() interface{} {
				return bchjson.Int64(5)
			},
			expected: func() interface{} {
				val := int64(5)
				return &val
			}(),
		},
		{
			name: "uint64",
			f: func() interface{} {
				return bchjson.Uint64(5)
			},
			expected: func() interface{} {
				val := uint64(5)
				return &val
			}(),
		},
		{
			name: "string",
			f: func() interface{} {
				return bchjson.String("abc")
			},
			expected: func() interface{} {
				val := "abc"
				return &val
			}(),
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.f()
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test #%d (%s) unexpected value - got %v, "+
				"want %v", i, test.name, result, test.expected)
			continue
		}
	}
}
