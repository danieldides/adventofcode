package main

import (
	"io/ioutil"
	"log"
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

	log.Printf("Part 1: %v\n", DoOne())
	log.Printf("Part 2: %v\n", DoTwo())

}

func DoOne() int {
	return 0
}

func DoTwo() int {
	return 0
}
