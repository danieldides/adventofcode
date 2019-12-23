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
	distance, steps := Compute(wire1, wire2)
	log.Printf("Part one: %v\n", distance)
	log.Printf("Part two: %v\n", steps)
}

// Compute takes in the input and gives us the processed output
func Compute(wire1, wire2 []string) (int, int) {
	visited1, m1 := drive(wire1)
	visited2, m2 := drive(wire2)

	intersection := intersect.Hash(visited1, visited2).([]interface{})

	var points []loc
	for _, point := range intersection {
		points = append(points, point.(loc))
	}

	if len(points) == 0 {
		log.Fatal("Empty intersections")
	}

	// Find the shortest distance
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

	// Calculate the minimum steps at all intersections
	minSteps := m1[points[0]] + m1[points[1]]
	for _, point := range points {
		steps := m1[point] + m2[point]
		if steps < minSteps {
			minSteps = steps
		}
	}

	return min, minSteps
}

type loc struct {
	x int
	y int
}

// Drive every point in path and return a list of all the X/Y coordinates
func drive(path []string) ([]loc, map[loc]int) {
	x := 0
	y := 0
	steps := 0

	var visited []loc

	// Map a location to a numeber of steps
	m := make(map[loc]int)

	for _, point := range path {
		cmd := string(point[0])
		rest := string(point[1:])
		distance, _ := strconv.Atoi(rest)
		switch cmd {
		case "U":
			for i := 0; i < distance; i++ {
				steps++
				y++
				l := loc{x, y}
				visited = append(visited, l)
				m[l] = steps
			}
		case "R":
			for i := 0; i < distance; i++ {
				steps++
				x++
				l := loc{x, y}
				visited = append(visited, l)
				m[l] = steps
			}
		case "D":
			for i := 0; i < distance; i++ {
				steps++
				y--
				l := loc{x, y}
				visited = append(visited, l)
				m[l] = steps
			}
		case "L":
			for i := 0; i < distance; i++ {
				steps++
				x--
				l := loc{x, y}
				visited = append(visited, l)
				m[l] = steps
			}
		}
	}
	return visited, m
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
