package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// File loading boilerplate
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")
	cs := strings.Split(lines[0], ",")

	var crabs []int

	for _, c := range cs {
		v, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal(err)
		}
		crabs = append(crabs, v)
	}

	log.Printf("Part 1: %v\n", DoOne(crabs))
	log.Printf("Part 2: %v\n", DoTwo(crabs))

}

func median(cc []int) int {
	sort.Ints(cc)
	return cc[len(cc)/2]

}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func DoOne(cc []int) int {
	m := median(cc)

	fuel := 0
	for c := range cc {
		cost := abs(cc[c], m)
		fuel += cost
	}

	return fuel
}

func gaussSum(n int) int {
	return (n * (n + 1)) / 2
}

func findDistance(x, y int) int {
	return gaussSum(abs(x, y))
}

func DoTwo(cc []int) int {
	sort.Ints(cc)

	var distances []int

	// Go through all the possible values between the biggest
	// and the smallest, and test the distances
	for i := cc[0]; i <= cc[len(cc)-1]; i++ {
		var distance int
		for _, c := range cc {
			distance += findDistance(c, i)
		}
		distances = append(distances, distance)
	}

	min := 0
	for i, d := range distances {
		if i == 0 || d < min {
			min = d
		}
	}

	return min
}
