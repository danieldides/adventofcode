package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var alphabet = makeAlphabet()

func main() {

	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(f), "\n")
	// Part 1
	score := 0
	for _, line := range lines {
		score += findPriority(findLetter(line))
	}
	fmt.Println("Part one: ", score)

	// Part 2
	score = 0
	for i := 0; i < len(lines)-1; i += 3 {
		score += findPriority(findItem([]line{line(lines[i]), line(lines[i+1]), line(lines[i+2])}))
	}

	fmt.Println("Part two: ", score)
}

func splitString(s string) (string, string) {
	half := len(s) / 2
	return s[:half], s[half:]

}

// findLetter finds the duplicate letter in either half of a given
// rucksack manifest
//
// e.g.:
//
//	The first rucksack contains the items vJrwpWtwJgWrhcsFMMfFFhFp,
//	which means its first compartment contains the items vJrwpWtwJgWr,
//	while the second compartment contains the items hcsFMMfFFhFp.
//	The only item type that appears in both compartments is lowercase p.
func findLetter(manifest string) string {
	var letter string

	first, second := splitString(manifest)

	for _, f := range first {
		for _, s := range second {
			if f == s {
				letter = string(s)
			}
		}
	}

	return letter
}

func makeAlphabet() map[string]int {
	alphabet := make(map[string]int)

	for i := 'a'; i <= 'z'; i++ {
		alphabet[string(i)] = int(i - 'a' + 1)
	}

	for i := 'A'; i <= 'Z'; i++ {
		alphabet[string(i)] = int(i - 'A' + 27)
	}

	return alphabet
}

func findPriority(letter string) int {
	return alphabet[letter]
}

type line string

func (l line) Contains(r rune) bool {
	for _, char := range l {
		if char == r {
			return true
		}
	}
	return false
}

// findItem takes in chunks of three lines and returns the common letter
// among all three lines
func findItem(lines []line) string {

	for _, letter := range lines[0] {
		score := 0
		if lines[1].Contains(letter) {
			score++
		}

		if lines[2].Contains(letter) {
			score++
		}

		if score == 2 {
			return string(letter)
		}
	}
	return ""
}
