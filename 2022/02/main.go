package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Read input.txt
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error reading input")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	partOneScore := 0
	partTwoScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) > 2 {
			log.Fatalf("only expecting two parts. line: %s", line)
		}
		theirs := parts[0]
		yours := parts[1]

		// Part 1
		score, err := calculateDuel(yours, theirs)
		if err != nil {
			log.Fatal(err)
		}
		partOneScore += score

		score, err = determineMove(yours, theirs)
		if err != nil {
			log.Fatal(err)
		}
		partTwoScore += score
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Part 1:", partOneScore)
	log.Println("Part 2:", partTwoScore)
}

const (
	rock     = 1
	paper    = 2
	scissors = 3

	loss = 0
	draw = 3
	win  = 6
)

// calculateDuel calculates the resulting score given the two inputs
// Their options:
//   - A - Rock
//   - B - Paper
//   - C - Scissors
//
// Your options:
//   - X - Rock			(1 point)
//   - Y - Paper		(2 points)
//   - Z - Scissors		(3 points)
//
// Other points:
//   - Loss	0 points
//   - Draw	3 points
//   - Win	6 points
func calculateDuel(yours, theirs string) (int, error) {
	score := 0
	switch theirs {
	case "A":
		switch yours {
		case "X":
			score = rock + draw
		case "Y":
			score = paper + win
		case "Z":
			score = scissors + loss
		default:
			return score, errors.New("invalid yours input")
		}
	case "B":
		switch yours {
		case "X":
			score = rock + loss
		case "Y":
			score = paper + draw
		case "Z":
			score = scissors + win
		default:
			return score, errors.New("invalid yours input")
		}

	case "C":
		switch yours {
		case "X":
			score = rock + win
		case "Y":
			score = paper + loss
		case "Z":
			score = scissors + draw
		default:
			return score, errors.New("invalid yours input")
		}

	default:
		return score, errors.New(fmt.Sprintf(`invalid theirs input %s`, theirs))
	}
	return score, nil
}

// determineMove determines the correct move given their move and the
// predetermined result
//
// Their options:
//   - A - Rock
//   - B - Paper
//   - C - Scissors
//
// Your options:
//   - X - Lose
//   - Y - Draw
//   - Z - Win
func determineMove(yours, theirs string) (int, error) {
	score := 0
	switch theirs {
	case "A":
		switch yours {
		case "X":
			score = loss + scissors
		case "Y":
			score = draw + rock
		case "Z":
			score = win + paper
		default:
			return score, errors.New("invalid yours input")
		}
	case "B":
		switch yours {
		case "X":
			score = loss + rock
		case "Y":
			score = draw + paper
		case "Z":
			score = win + scissors
		default:
			return score, errors.New("invalid yours input")
		}

	case "C":
		switch yours {
		case "X":
			score = loss + paper
		case "Y":
			score = draw + scissors
		case "Z":
			score = win + rock
		default:
			return score, errors.New("invalid yours input")
		}

	default:
		return score, errors.New(fmt.Sprintf(`invalid theirs input %s`, theirs))
	}

	return score, nil
}
