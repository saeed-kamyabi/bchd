// Copyright (c) 2014 The bchsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package bchjson_test

import (
	"testing"

	"github.com/bchsuite/bchd/bchjson"
)

// TestErrorCodeStringer tests the stringized output for the ErrorCode type.
func TestErrorCodeStringer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   bchjson.ErrorCode
		want string
	}{
		{bchjson.ErrDuplicateMethod, "ErrDuplicateMethod"},
		{bchjson.ErrInvalidUsageFlags, "ErrInvalidUsageFlags"},
		{bchjson.ErrInvalidType, "ErrInvalidType"},
		{bchjson.ErrEmbeddedType, "ErrEmbeddedType"},
		{bchjson.ErrUnexportedField, "ErrUnexportedField"},
		{bchjson.ErrUnsupportedFieldType, "ErrUnsupportedFieldType"},
		{bchjson.ErrNonOptionalField, "ErrNonOptionalField"},
		{bchjson.ErrNonOptionalDefault, "ErrNonOptionalDefault"},
		{bchjson.ErrMismatchedDefault, "ErrMismatchedDefault"},
		{bchjson.ErrUnregisteredMethod, "ErrUnregisteredMethod"},
		{bchjson.ErrNumParams, "ErrNumParams"},
		{bchjson.ErrMissingDescription, "ErrMissingDescription"},
		{0xffff, "Unknown ErrorCode (65535)"},
	}

	// Detect additional error codes that don't have the stringer added.
	if len(tests)-1 != int(bchjson.TstNumErrorCodes) {
		t.Errorf("It appears an error code was added without adding an " +
			"associated stringer test")
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

// TestError tests the error output for the Error type.
func TestError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   bchjson.Error
		want string
	}{
		{
			bchjson.Error{Description: "some error"},
			"some error",
		},
		{
			bchjson.Error{Description: "human-readable error"},
			"human-readable error",
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.Error()
		if result != test.want {
			t.Errorf("Error #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}
