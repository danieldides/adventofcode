package main

import (
	"fmt"
	"testing"
)

func TestDoOne(t *testing.T) {
	tt := []struct {
		In     []Command
		Result int
	}{
		{
			In: []Command{
				{
					direction: "forward",
					value:     5,
				},
				{
					direction: "down",
					value:     5,
				},
				{
					direction: "forward",
					value:     8,
				},
				{
					direction: "up",
					value:     3,
				},
				{
					direction: "down",
					value:     8,
				},
				{
					direction: "forward",
					value:     2,
				},
			},
			Result: 150,
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
		In     []Command
		Result int
	}{
		{
			In: []Command{
				{
					direction: "forward",
					value:     5,
				},
				{
					direction: "down",
					value:     5,
				},
				{
					direction: "forward",
					value:     8,
				},
				{
					direction: "up",
					value:     3,
				},
				{
					direction: "down",
					value:     8,
				},
				{
					direction: "forward",
					value:     2,
				},
			},
			Result: 900,
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
