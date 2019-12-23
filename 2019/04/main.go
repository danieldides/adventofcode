package main

import (
	"log"
	"strconv"
)

func main() {
	start := 172930
	end := 683082

	n := Compute(start, end)
	log.Printf("Part one: %v\n", n)
}

// Compute takes in the input and gives us the processed output
func Compute(start, end int) int {
	var nums []int

	for idx := start; idx < end; idx++ {
		if !DuplicateNums(idx) {
			continue
		}
		if !IncreasingNums(idx) {
			continue
		}
		nums = append(nums, idx)
	}
	return len(nums)
}

// DuplicateNums determines if two adjacent numbers are the same
func DuplicateNums(n int) bool {
	s := strconv.Itoa(n)
	return s[0] == s[1] || s[1] == s[2] || s[2] == s[3] || s[3] == s[4] || s[4] == s[5]
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
