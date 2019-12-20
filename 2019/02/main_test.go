package main

import (
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, test := range tests {
		out := Parse(test.in)
		if !assertEqual(test.out, out) {
			t.Errorf("Fail! Input: %v Expected: %v Got: %v", test.in, test.out, out)
		}
	}
}

func assertEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
