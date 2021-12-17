package main

import (
	"fmt"
	"testing"
)

var tt = []struct {
	In     interface{}
	Result int
}{
	{},
}

func TestDoOne(t *testing.T) {
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

func TestDoTwo(t *testing.T) {
	for i, test := range tt {
		testName := fmt.Sprintf("%v", i)
		t.Run(testName, func(t *testing.T) {
			result := DoTwo(test.In)
			if result != test.Result {
				t.Errorf("Invalid result. Expected: %v Got: %v", test.Result, result)
			}
		})
	}
}
