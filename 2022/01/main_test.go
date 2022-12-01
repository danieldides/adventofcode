package main

import (
	"fmt"
	"testing"
)

var tt = []struct {
	In       []string
	Elf      int
	Weight   int
	TopThree int
}{
	{
		In: []string{
			"1000",
			"2000",
			"3000",
			"",
			"4000",
			"",
			"5000",
			"6000",
			"",
			"7000",
			"8000",
			"9000",
			"",
			"10000",
			"",
		},
		Elf:      4,
		Weight:   24000,
		TopThree: 45000,
	},
}

func TestDoOne(t *testing.T) {
	for i, test := range tt {
		testName := fmt.Sprintf("%v", i)
		t.Run(testName, func(t *testing.T) {
			e, w := DoOne(test.In)
			if e != test.Elf || w != test.Weight {
				t.Errorf("Invalid. Expected: %v from %v Got: %v from %v", test.Weight, test.Elf, w, e)
			}
		})
	}

}
func TestDoTwo(t *testing.T) {
	for i, test := range tt {
		testName := fmt.Sprintf("%v", i)
		t.Run(testName, func(t *testing.T) {
			w := DoTwo(test.In)
			if w != test.TopThree {
				t.Errorf("Invalid. Expected: %v Got: %v ", test.TopThree, w)
			}
		})
	}
}
