package main

import (
	"fmt"
	"testing"
)

var tt = []struct {
	first         string
	second        string
	fullyContains bool
	overlaps      bool
}{
	{
		first:         "2-4",
		second:        "6-8",
		fullyContains: false,
		overlaps:      false,
	},
	{
		first:         "2-3",
		second:        "4-5",
		fullyContains: false,
		overlaps:      false,
	},
	{
		first:         "5-7",
		second:        "7-9",
		fullyContains: false,
		overlaps:      true,
	},
	{
		first:         "2-8",
		second:        "3-7",
		fullyContains: true,
		overlaps:      true,
	},
	{
		first:         "6-6",
		second:        "4-6",
		fullyContains: true,
		overlaps:      true,
	},
	{
		first:         "2-6",
		second:        "4-8",
		fullyContains: false,
		overlaps:      true,
	},
}

func TestFullyContains(t *testing.T) {

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%s,%s", tc.first, tc.second), func(t *testing.T) {
			contains := fullyContains(tc.first, tc.second)
			if contains != tc.fullyContains {
				t.Errorf("expected %v, got %v", tc.fullyContains, contains)
			}
		})
	}

}

func TestOverlaps(t *testing.T) {
	for _, tc := range tt {
		t.Run(fmt.Sprintf("%s,%s", tc.first, tc.second), func(t *testing.T) {
			overlaps := assignmentsOverlap(tc.first, tc.second)
			if overlaps != tc.overlaps {
				t.Errorf("expected %v, got %v", tc.overlaps, overlaps)
			}
		})
	}

}
