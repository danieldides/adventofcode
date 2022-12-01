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
	_ = lines

	e, w := DoOne(lines)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Part 1: %v by %v\n", e, w)

	e = DoTwo(lines)

	log.Printf("Part 2: %v\n", e)

}

func findMax(ii []int) (int, int) {
	var elf, max int

	for i, v := range ii {
		if v > max {
			elf = i
			max = v
		}
	}

	return elf + 1, max
}

func findTopThree(ii []int) int {
	var c int

	sort.Ints(ii)
	sort.Sort(sort.Reverse(sort.IntSlice(ii)))

	for i := 0; i < 3; i++ {
		c += ii[i]
	}

	return c
}

// parseElves takes the input, parses the lines, and then
// returns a slice of total weights each elf is carrying
func parseElves(lines []string) []int {
	var elf, weight int
	var elves []int

	for _, line := range lines {
		w, err := strconv.Atoi(line)
		weight += w
		if err != nil {
			elf++
			elves = append(elves, weight)
			weight = 0
			continue
		}
	}

	return elves
}

func DoOne(lines []string) (int, int) {
	var e, w int
	elves := parseElves(lines)
	e, w = findMax(elves)
	return e, w
}

func DoTwo(lines []string) int {
	elves := parseElves(lines)
	e := findTopThree(elves)
	return e
}
