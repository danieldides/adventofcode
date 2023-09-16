package main

import (
	"fmt"
	"testing"
)

var tt = []struct {
	manifest string
	letter   string
	score    int
}{
	{
		manifest: "vJrwpWtwJgWrhcsFMMfFFhFp",
		letter:   "p",
		score:    16,
	},
	{
		manifest: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		letter:   "L",
		score:    38,
	},
	{
		manifest: "PmmdzqPrVvPwwTWBwg",
		letter:   "P",
		score:    42,
	},
	{
		manifest: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		letter:   "v",
		score:    22,
	},
	{
		manifest: "ttgJtRGJQctTZtZT",
		letter:   "t",
		score:    20,
	},
	{
		manifest: "CrZsJsPPZsGzwwsLwLmpwMDw",
		letter:   "s",
		score:    19,
	},
}

func TestFindLetter(t *testing.T) {
	for _, tc := range tt {
		t.Run(tc.manifest, func(t *testing.T) {
			letter := findLetter(tc.manifest)

			if letter != tc.letter {
				t.Errorf("letter %s got %s", tc.letter, letter)
			}
		})
	}
}

func TestCalculatePriority(t *testing.T) {
	for _, tc := range tt {
		t.Run(tc.letter, func(t *testing.T) {
			score := findPriority(tc.letter)
			if score != tc.score {
				t.Errorf("expected %d, got %d", tc.score, score)
			}
		})
	}
}

func TestFindItem(t *testing.T) {
	tt := []struct {
		lines  []line
		letter string
	}{
		{
			lines:  []line{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"},
			letter: "r",
		},
		{
			lines:  []line{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"},
			letter: "Z",
		},
	}

	for idx, tc := range tt {
		t.Run(fmt.Sprintf("part_two_%v", idx), func(t *testing.T) {
			item := findItem(tc.lines)
			if item != tc.letter {
				t.Errorf("expected %s, got %s", tc.letter, item)
			}
		})
	}

}
