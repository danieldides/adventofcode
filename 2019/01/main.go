package main

import (
	"bufio" "log"
	"math"
	"os"
	"strconv"
)

// Fuel required to launch a given module is based on its mass.
// Specifically, to find the fuel required for a module, take its mass, divide by three,
// round down, and subtract 2.
func calculateFuel(mass int) int {
	return int(math.Floor(float64(mass)/3)) - 2
}

// Fuel itself requires fuel just like a module - take its mass, divide by three, round down, and subtract 2.
// However, that fuel also requires fuel, and that fuel requires fuel, and so on.
// Any mass that would require negative fuel should instead be treated as if it requires zero fuel;
func calculateFuelCompound(mass, total int) int {
	// fmt.Printf("Called: %v, %v\n", mass, total)

	fuel := calculateFuel(mass)

	if fuel <= 0 {
		// fmt.Printf("Returning. Last call: %v, %v\n", mass, total)
		return total
	}

	total += fuel
	return calculateFuelCompound(fuel, total)
}

func partOne() (int, error) {

	file, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		val, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		val = calculateFuel(val)
		total = total + val
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return total, nil
}

func partTwo() (int, error) {

	file, err := os.Open("input.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		val, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		val = calculateFuelCompound(val, 0)
		total = total + val
	}

	return total, nil
}

func main() {
	log.Println("Part One")

	// Tests
	//log.Printf("Test 1: %v\n", calculateFuel(12))
	//log.Printf("Test 2: %v\n", calculateFuel(14))
	//log.Printf("Test 3: %v\n", calculateFuel(1969))
	//log.Printf("Test 4: %v\n", calculateFuel(100756))

	total, err := partOne()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Final: %v\n", total)

	log.Println("Part Two")

	// Tests
	// fmt.Printf("Test 1: %v\n", calculateFuelCompound(1969))
	// fmt.Printf("Test 2: %v\n", calculateFuelCompound(100756, 0))

	total, err = partTwo()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Final: %v\n", total)

}
