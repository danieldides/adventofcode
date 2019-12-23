package main

import "testing"

// TestDuplicateNums verifies theres at least 1 set of repeating numbers
// that is not part of a larger set.
func TestDuplicateNums(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{123444, false},
		{111111, true},
		{110123, true},
		{123456, false},
		{124456, true},
		{111122, true},
	}

	for _, test := range tests {
		out := DuplicateNums(test.in)
		if out != test.out {
			t.Errorf("Fail duplicate numbers! Input: %v Expected: %v Got: %v",
				test.in, test.out, out)
		}
		break
	}
}

func TestIncreasingNums(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{223450, false},
		{123456, true},
		{987654, false},
	}

	for _, test := range tests {
		out := IncreasingNums(test.in)
		if out != test.out {
			t.Errorf("Fail increasing numbers! Input: %v Expected: %v Got: %v", test.in, test.out, out)
		}
	}
}
