package main

import (
	"testing"
)

func TestCompute(t *testing.T) {

	tests := []struct {
		wire1 []string
		wire2 []string
		out   int
	}{
		{
			wire1: []string{"R8", "U5", "L5", "D3"},
			wire2: []string{"U7", "R6", "D4", "L7"},
			out:   6,
		},

		{
			wire1: []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			wire2: []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			out:   159,
		},
		{
			wire1: []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
			wire2: []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			out:   135,
		},
	}

	for i, test := range tests {
		out := Compute(test.wire1, test.wire2)
		if out != test.out {
			t.Errorf("Fail! Input: %v Expected: %v Got: %v", i+1, test.out, out)
		}
	}
}
