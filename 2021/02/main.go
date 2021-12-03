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
	aim      int
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
			fmt.Println("Bad input: ", parts)
			continue
		}
		c := Command{
			direction: parts[0],
			value:     Atoi(parts[1]),
		}
		in = append(in, c)
	}

	log.Printf("Part 1: %v\n", DoOne(in))
	log.Printf("Part 2: %v\n", DoTwo(in))

}

func DoOne(in []Command) int {
	ship := Ship{
		position: 0,
		depth:    0,
		aim:      0,
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

	return ship.position * ship.depth
}

func DoTwo(in []Command) int {
	ship := Ship{
		position: 0,
		depth:    0,
		aim:      0,
	}

	for _, command := range in {

		switch command.direction {
		case "forward":
			ship.position += command.value
			ship.depth += ship.aim * command.value
		case "down":
			ship.aim += command.value
		case "up":
			ship.aim -= command.value
		}
	}

	return ship.position * ship.depth
}
