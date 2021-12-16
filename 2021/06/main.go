package main

import (
	"fmt"
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

func DoOne(fish []int) int {
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

func sum(m map[int]int) int {
	var result int

	for _, v := range m {
		result += v
	}

	return result
}

func DoTwo(fish []int) int {
	// Map of age: number of fish at that age:w
	ages := [9]int{}

	// Initial population
	for _, f := range fish {
		ages[f] += 1
	}

	for day := 0; day < 18; day++ {
		fmt.Println("Day: ", day, sum(ages))
	}

	fmt.Println(ages)
	return len(fish)
}
