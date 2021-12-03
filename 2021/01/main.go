package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func main() {
	// File loading boilerplate
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var in []int

	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		in = append(in, Atoi(line))
	}

	log.Printf("Part 1: %v\n", DoOne(in))
	log.Printf("Part 2: %v\n", DoTwo(in))

}

func countIncreases(in []int) int {
	var result int

	for i := 1; i < len(in); i++ {
		// fmt.Print(in[i])
		if in[i] > in[i-1] {
			// fmt.Print(" Increased\n")
			result++
			continue
		}
		// fmt.Print(" Decreased\n")
	}
	return result
}

// Part 1
func DoOne(in []int) int {
	return countIncreases(in)
}

// Part 2
func DoTwo(in []int) int {
	var v []int
	for i := 0; i < len(in)-2; i++ {
		x, y, z := in[i], in[i+1], in[i+2]
		sum := x + y + z
		// fmt.Println(x, y, z, sum)
		v = append(v, sum)
	}
	return countIncreases(v)
}
