package main

import (
	"fmt"
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
		fmt.Println(c, cost)
		fuel += cost
	}

	fmt.Println(fuel)

	return fuel
}

func DoTwo(c []int) int {
	return 0
}
