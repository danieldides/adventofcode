package main

import "testing"

func TestCompute(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"in", "out"},
	}

	for _, test := range tests {
		out := Compute(test.in)
		if out != test.out {
			t.Errorf("Fail! Input: %v Expected: %v Got: %v", test.in, test.out, out)
		}
	}
}
