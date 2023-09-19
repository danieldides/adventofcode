package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type assignment struct {
	start int
	end   int
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func createAssignemnts(a, b string) (assignment, assignment) {

	partsA := strings.Split(a, "-")
	first := assignment{
		start: atoi(partsA[0]),
		end:   atoi(partsA[1]),
	}

	partsB := strings.Split(b, "-")
	second := assignment{
		start: atoi(partsB[0]),
		end:   atoi(partsB[1]),
	}

	return first, second
}

// fullyContains checks if two elf assignments fully encapsulate each other,
// e.g.
//
//  2345678. 2-8
//     .34567..  3-7
//
// To do this we'll check the start and end stops for each assignment, noting if the start
// of one assignment is less than the other while the end of one is greater than the other
// (a[start] < b[start] && a[end] > b[end]) || b[start] < a[start] && b[end] > a[end]
func fullyContains(a, b string) bool {
	first, second := createAssignemnts(a, b)

	return (first.start <= second.start && first.end >= second.end) || (second.start <= first.start) && (second.end >= first.end)
}

// assignmentsOverlap walks the ranges from start->end for each assignment, and verifies
// that no number in first assignment is ever found in the second assignment
func assignmentsOverlap(a, b string) bool {
	first, second := createAssignemnts(a, b)

	for i := first.start; i < first.end+1; i++ {
		for j := second.start; j < second.end+1; j++ {
			if i == j {
				return true
			}
		}
	}
	return false
}

func main() {

	// Read in the file input
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	partOne := 0
	partTwo := 0

	// Process the file line-by-line
	for _, line := range lines {
		// Split each line on the comma to break the sections into first and second assignments
		parts := strings.Split(line, ",")

		first := parts[0]
		second := parts[1]

		// Compare the start and stop ranges of each assignment to see if one ecapsulates the other
		if fullyContains(first, second) {
			partOne++
		}

		// Check if there's any overlap at all
		if assignmentsOverlap(first, second) {
			partTwo++
		}

	}

	fmt.Println("Number of assignments that fully contain one another:", partOne)
	fmt.Println("Number of assignments that overlap at all:", partTwo)

}
