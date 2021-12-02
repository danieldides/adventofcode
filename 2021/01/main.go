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

}

// Part 1
func DoOne(in []int) int {
	var result int

	for i := 1; i < len(in); i++ {
		if in[i] >= in[i-1] {
			result++
		}
	}
	return result
}
