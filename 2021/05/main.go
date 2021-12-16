package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// File loading boilerplate
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Part 1: %v\n", DoOne(string(file)))
	log.Printf("Part 2: %v\n", DoTwo(string(file)))

}

type Pair struct {
	X int
	Y int
}

type Line struct {
	Start Pair
	End   Pair
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

// findMax will consider all coordinates and figure out the largest
// one so we can accurately dimension our grid
func findMax(i []int) int {
	var max int

	for _, n := range i {
		if n > max {
			max = n
		}
	}

	return max
}

func parseInput(s string) []Line {
	var lines []Line

	ss := strings.Split(s, "\n")

	for _, l := range ss {
		if l == "" {
			continue
		}
		f := strings.Fields(l)

		start := strings.Split(f[0], ",")

		x1 := Atoi(start[0])
		y1 := Atoi(start[1])

		end := strings.Split(f[2], ",")

		x2 := Atoi(end[0])
		y2 := Atoi(end[1])

		lines = append(lines, Line{
			Start: Pair{x1, y1},
			End:   Pair{x2, y2},
		})
	}

	return lines

}

func sortInt(x, y int) (int, int) {
	if x < y {
		return x, y
	}
	return y, x
}

func calculateScore(g [][]int) int {
	var score int
	for _, row := range g {
		for _, col := range row {
			if col > 1 {
				score++
			}
		}
	}

	return score
}

func makeGrid(lines []Line) [][]int {
	var xPoints, yPoints []int

	for i := range lines {
		xPoints = append(xPoints, lines[i].Start.X)
		xPoints = append(xPoints, lines[i].End.X)
	}
	for i := range lines {
		yPoints = append(yPoints, lines[i].Start.Y)
		yPoints = append(yPoints, lines[i].End.Y)
	}

	maxX := findMax(xPoints)
	maxY := findMax(yPoints)

	fmt.Printf("Making a %v x %v grid.\n", maxX, maxY)
	// Construct our grid
	grid := make([][]int, maxY+1)

	for i := range grid {
		grid[i] = make([]int, maxX+1)
	}

	return grid
}

func drawLines(lines []Line, grid [][]int, diag bool) {
	// Start incrementing values
	for _, l := range lines {

		// Drawing a vertical line
		if l.Start.X == l.End.X {
			begin, end := sortInt(l.Start.Y, l.End.Y)

			for i := begin; i < end+1; i++ {
				grid[i][l.Start.X] += 1
			}
		}

		// Drawing a horizontal line
		if l.Start.Y == l.End.Y {
			begin, end := sortInt(l.Start.X, l.End.X)

			for i := begin; i < end+1; i++ {
				grid[l.Start.Y][i] += 1
			}
		}

		// Draw diagonally
		if diag && l.Start.Y != l.End.Y && l.Start.X != l.End.X {
			fmt.Printf("Drawing line from %v to %v\n", l.Start, l.End)
			points := solveLine(l)
			for _, p := range points {
				grid[p.Y][p.X] += 1
			}
		}
	}
}

func solveLine(l Line) []Pair {
	var points []Pair

	slope := (l.End.Y - l.Start.Y) / (l.End.X - l.Start.X)

	b := -1 * ((slope * l.Start.X) - l.Start.Y)

	min, max := sortInt(l.Start.X, l.End.X)

	// fmt.Println("Diag from ", min, "to ", max)

	// Iterate through x range, calculate y values
	for x := min; x < max+1; x++ {
		y := slope*x + b
		pair := Pair{X: x, Y: y}

		points = append(points, pair)
	}

	return points

}

func DoOne(s string) int {
	lines := parseInput(s)

	grid := makeGrid(lines)

	drawLines(lines, grid, false)

	return calculateScore(grid)
}

func DoTwo(s string) int {
	lines := parseInput(s)

	grid := makeGrid(lines)

	drawLines(lines, grid, true)

	// Print the row. Useful for testing only
	// for _, row := range grid {
	// 	fmt.Println(row)
	// }

	return calculateScore(grid)
}
