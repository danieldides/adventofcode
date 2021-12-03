package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type Command struct {
	direction string
	value     int
}

type Ship struct {
	position int
	depth    int
}

func main() {
	// File loading boilerplate
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var in []Command

	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			fmt.Println(parts)
			continue
		}
		c := Command{
			direction: parts[0],
			value:     Atoi(parts[1]),
		}
		in = append(in, c)
	}

	log.Printf("Part 1: %v\n", DoOne(in))
	// log.Printf("Part 2: %v\n", DoTwo(in))

}

func findPosition(in []Command) int {
	ship := Ship{
		position: 0,
		depth:    0,
	}

	for _, command := range in {

		switch command.direction {
		case "forward":
			ship.position += command.value
		case "down":
			ship.depth += command.value
		case "up":
			ship.depth -= command.value
		}
	}

	fmt.Println(ship)

	return ship.position * ship.depth
}

// Part 1
func DoOne(in []Command) int {
	return findPosition(in)
}

// Part 2
func DoTwo(in []Command) int {
	return findPosition(in)
}
