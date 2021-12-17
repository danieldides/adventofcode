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
			In:     []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			Result: 37,
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

func TestDoTwo(t *testing.T) {
	tt := []struct {
		In     []int
		Result int
	}{
		{
			In:     []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			Result: 168,
		},
	}

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
