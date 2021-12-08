package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Board [][]int

func prettyPrint(b []Board) {
	for i := range b {
		for y := range b[i] {
			fmt.Println(b[i][y])
		}
		fmt.Println()
	}
}

// Game considers the given board and all the inputs so far
// to see if we have a solution.
// This is really the algo
func check(board Board) bool {
	var win bool

	// Vertical win
	for y := range board {
		winning := true
		for x := range board {
			if board[x][y] != -1 {
				winning = false
				break
			}

		}
		if winning {
			// fmt.Println("you win vertically")
			// prettyPrint([]Board{board})
			return true
		}

	}

	// Horizontal win
	for y := range board {
		winning := true
		for x := range board {
			if board[y][x] != -1 {
				winning = false
				break
			}

		}
		if winning {
			// fmt.Println("you win horizontally")
			// prettyPrint([]Board{board})
			return true
		}

	}

	return win
}

func parseInput(in string) ([]int, []Board) {
	lines := strings.Split(in, "\n")

	var inputs []int
	ss := strings.Split(lines[0], ",")
	for _, s := range ss {
		val, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, val)
	}

	// Skip empty line between inputs and board
	var boards []Board
	var board Board

	for _, r := range lines[2:] {
		var row []int

		line := strings.Fields(r)
		if len(line) == 0 {
			// Gap between boards indicates we're at the end
			// of the previous board. Add it to our map and
			// reset our board
			boards = append(boards, board)
			board = Board{}
			continue
		}
		for _, l := range line {
			val, err := strconv.Atoi(l)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, val)
		}
		board = append(board, row)
	}

	return inputs, boards

}

func mark(in int, bb []Board) {
	// Mark the input as a -1 so we don't count it later
	// Can't use 0 because that's a valid number in the inputs
	for i := 0; i < len(bb); i++ {
		for y := 0; y < len(bb[i]); y++ {
			for x := 0; x < len(bb[i][y]); x++ {
				if bb[i][y][x] == in {
					bb[i][y][x] = -1
				}
			}
		}
	}
}

func calculateScore(in int, b Board) int {
	var score int

	for y := range b {
		for x := range b[y] {
			if b[y][x] != -1 {
				score += b[y][x]
			}
		}
	}
	return score * in
}

func SolveOne(inputs []int, boards []Board) int {
	var score int

	for _, input := range inputs {
		mark(input, boards)
		for _, board := range boards {
			if check(board) {
				return calculateScore(input, board)
			}
		}
	}

	return score

}

func SolveTwo(inputs []int, boards []Board) int {
	var score int

	for len(boards) > 0 {
	out:
		for _, input := range inputs {
			mark(input, boards)
			for i, board := range boards {
				if check(board) {
					score = calculateScore(input, board)
					boards = append(boards[:i], boards[i+1:]...)
					break out
				}
			}
		}
	}

	return score
}

func main() {
	// File loading boilerplate
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Part 1: %v\n", DoOne(string(file)))
	log.Printf("Part 2: %v\n", DoTwo(string(file)))

}

func DoOne(s string) int {
	inputs, boards := parseInput(s)
	return SolveOne(inputs, boards)
}

func DoTwo(s string) int {
	inputs, boards := parseInput(s)
	return SolveTwo(inputs, boards)
}
