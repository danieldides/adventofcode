package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// File loading boilerplate
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var fishes []int

	s := strings.Split(strings.Replace(string(file), "\n", "", -1), ",")
	for _, s := range s {
		v, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		fishes = append(fishes, v)
	}

	log.Printf("Part 1: %v\n", DoOne(fishes))
	log.Printf("Part 2: %v\n", DoTwo(fishes))

}

func DoOld(fish []int) int {
	for day := 0; day < 80; day++ {
		for i := 0; i < len(fish); i++ {
			if fish[i] == 0 {
				fish = append(fish, 9)
				fish[i] = 6
				continue
			}
			fish[i]--
		}
	}
	return len(fish)
}

func shift(a [9]int) (int, [9]int) {
	var b [9]int

	zeroes := a[0]

	for i := 1; i < len(a); i++ {
		b[i-1] = a[i]
	}

	return zeroes, b
}

func DoNew(fish []int, days int) int {

	var zeroes int
	var a [9]int

	a = [9]int{}

	// Initial population
	for _, f := range fish {
		a[f] += 1
	}

	for day := 0; day < days; day++ {
		zeroes, a = shift(a)

		a[6] += zeroes
		a[8] += zeroes

		total := 0
		for s := range a {
			total += a[s]
		}

		// fmt.Println(day, total)

	}

	total := 0
	for s := range a {
		total += a[s]
	}

	return total
}

func DoOne(fish []int) int {
	return DoNew(fish, 80)
}

func DoTwo(fish []int) int {
	return DoNew(fish, 256)
}
