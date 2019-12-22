package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/juliangruber/go-intersect"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(file), "\n")
	wire1 := strings.Split(parts[0], ",")
	wire2 := strings.Split(parts[1], ",")
	ret := Compute(wire1, wire2)
	log.Printf("Part one: %v\n", ret)
}

// Compute takes in the input and gives us the processed output
func Compute(wire1, wire2 []string) int {
	visited1 := drive(wire1)
	visited2 := drive(wire2)

	intersection := intersect.Hash(visited1, visited2).([]interface{})

	var points []loc
	for _, point := range intersection {
		points = append(points, point.(loc))
	}

	var sums []int
	for _, p := range points {
		sum := abs(p.x) + abs(p.y)
		sums = append(sums, sum)
	}

	min := sums[0]
	for _, sum := range sums {
		if sum < min {
			min = sum
		}
	}

	return min
}

type loc struct {
	x int
	y int
}

// Drive every point in path and return a list of all the X/Y coordinates
func drive(path []string) []loc {

	x := 0
	y := 0

	var visited []loc

	for _, point := range path {

		cmd := string(point[0])
		rest := string(point[1:])
		distance, _ := strconv.Atoi(rest)

		switch cmd {
		case "U":
			for i := 0; i < distance; i++ {
				y++
				visited = append(visited, loc{x, y})
			}
		case "R":
			for i := 0; i < distance; i++ {
				x++
				visited = append(visited, loc{x, y})
			}
		case "D":
			for i := 0; i < distance; i++ {
				y--
				visited = append(visited, loc{x, y})
			}
		case "L":
			for i := 0; i < distance; i++ {
				x--
				visited = append(visited, loc{x, y})
			}

		}
		visited = append(visited, loc{x, y})
	}
	return visited
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
