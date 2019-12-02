package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

// Fuel required to launch a given module is based on its mass.
// Specifically, to find the fuel required for a module, take its mass, divide by three,
// round down, and subtract 2.
func calculateMass(mass int) int {

	return int(math.Floor(float64(mass)/3)) - 2
}

func partOne() {
	// Tests
	log.Println("Test values")
	log.Println(calculateMass(12))
	log.Println(calculateMass(14))
	log.Println(calculateMass(1969))
	log.Println(calculateMass(100756))

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		val = calculateMass(val)
		total = total + val
	}

	log.Printf("Final value: %v\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func partTwo() {
	log.Println("Test values")
}

func main() {
	// partOne()
	partTwo()
}
