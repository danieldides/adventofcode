package main

import (
	"log"
	"strconv"
)

func main() {
	// From input
	start := 172930
	end := 683082

	n := Compute(start, end)
	log.Printf("Part one: %v\n", n)
}

// Compute takes in the input and gives us the processed output
func Compute(start, end int) int {
	var nums int

	for idx := start; idx < end; idx++ {
		if !DuplicateNums(idx) {
			continue
		}
		if !IncreasingNums(idx) {
			continue
		}
		nums++
	}
	return nums
}

// DuplicateNums determines if two adjacent numbers are the same
func DuplicateNums(n int) bool {
	s := strconv.Itoa(n)

	pair := false
	for i := 0; i < 5; i++ {

		if s[i] == s[i+1] {
			if (i-1 < 0 || s[i-1] != s[i]) && (i+2 >= len(s) || s[i+2] != s[i]) {
				pair = true
			}
		}
	}
	return pair
}

// IncreasingNums verifies that numbers only increase or stay the same
func IncreasingNums(n int) bool {
	s := strconv.Itoa(n)

	for i := 0; i < 5; i++ {
		a, _ := strconv.Atoi(string(s[i]))
		b, _ := strconv.Atoi(string(s[i+1]))
		if a > b {
			return false
		}
	}
	return true
}
