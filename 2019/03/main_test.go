package main

import (
	"testing"
)

func TestCompute(t *testing.T) {

	tests := []struct {
		wire1    []string
		wire2    []string
		distance int
		steps    int
	}{
		{
			wire1:    []string{"R8", "U5", "L5", "D3"},
			wire2:    []string{"U7", "R6", "D4", "L7"},
			distance: 6,
			steps:    30,
		},

		{
			wire1:    []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			wire2:    []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			distance: 159,
			steps:    610,
		},
		{
			wire1:    []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
			wire2:    []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			distance: 135,
			steps:    410,
		},
	}

	for i, test := range tests {
		distance, steps := Compute(test.wire1, test.wire2)
		if distance != test.distance {
			t.Errorf("Distance Fail! Input: %v Expected: %v Got: %v", i+1, test.distance, distance)
		}
		if steps != test.steps {
			t.Errorf("Steps Fail! Input: %v Expected: %v Got: %v", i+1, test.steps, steps)
		}
	}
}
