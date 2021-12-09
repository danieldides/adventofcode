package main

import (
	"fmt"
	"testing"
)

func TestDoOne(t *testing.T) {
	tt := []struct {
		In     interface{}
		Result int
	}{
		{},
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