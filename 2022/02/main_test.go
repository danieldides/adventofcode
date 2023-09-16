package main

import "testing"

func TestPartOne(t *testing.T) {
	tt := []struct {
		name      string
		yourMove  string
		theirMove string
		score     int
	}{
		{
			name:      "example1",
			theirMove: "A",
			yourMove:  "Y",
			score:     8,
		},
		{
			name:      "example2",
			theirMove: "B",
			yourMove:  "X",
			score:     1,
		},
		{
			name:      "example3",
			theirMove: "C",
			yourMove:  "Z",
			score:     6,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			score, err := calculateDuel(tc.yourMove, tc.theirMove)
			if err != nil {
				t.Errorf(err.Error())
			}

			if score != tc.score {
				t.Errorf("mismatch. Got %d expected %d, inputs: %s %s", score, tc.score, tc.yourMove, tc.theirMove)
			}
		})
	}

}

func TestPartTwo(t *testing.T) {
	tt := []struct {
		name      string
		yourMove  string
		theirMove string
		score     int
	}{
		{
			name:      "example1",
			theirMove: "A",
			yourMove:  "Y",
			score:     4,
		},
		{
			name:      "example2",
			theirMove: "B",
			yourMove:  "X",
			score:     1,
		},
		{
			name:      "example3",
			theirMove: "C",
			yourMove:  "Z",
			score:     7,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			score, err := determineMove(tc.yourMove, tc.theirMove)
			if err != nil {
				t.Errorf(err.Error())
			}

			if score != tc.score {
				t.Errorf("mismatch. Got %d expected %d, inputs: %s %s", score, tc.score, tc.yourMove, tc.theirMove)
			}
		})
	}

}
