package main

import (
	"fmt"
	"testing"
)

func TestDoOne(t *testing.T) {
	tt := []struct {
		In     []int
		Result int
	}{
		{
			In:     []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			Result: 7,
		},
	}

	for i, test := range tt {
		testName := fmt.Sprintf("%v", i)
		t.Run(testName, func(t *testing.T) {
			result := DoOne(test.In)
			if result != test.Result {
				t.Errorf("Invalid result. Expected: %v Got: %v", test.Result, result)
			}
		})
	}
}
