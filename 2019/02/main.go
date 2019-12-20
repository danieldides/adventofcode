package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(file), ",")
	code := make([]int, len(input))
	for i, s := range input {
		code[i], _ = strconv.Atoi(s)
	}

	// Part one shenanigans
	code[1] = 12
	code[2] = 2
	res := Parse(code)
	log.Printf("Part one: %v\n", res[0])

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			code[1] = noun
			code[2] = verb
			res = Parse(code)
			if res[0] == 19690720 {
				log.Printf("Part two: %v %v\n", noun, verb)
				break
			}
		}
	}

}

// Compute takes in the input and gives us the processed output
func Compute(input string) string {
	return "out"
}

func Parse(input []int) []int {
	memory := make([]int, len(input))
	copy(memory, input)

	ptr := 0
	for {
		instruction := memory[ptr]
		switch instruction {
		case 1:
			paramOne := memory[ptr+1]
			paramTwo := memory[ptr+2]
			address := memory[ptr+3]
			memory[address] = memory[paramOne] + memory[paramTwo]
			ptr += 4
		case 2:
			paramOne := memory[ptr+1]
			paramTwo := memory[ptr+2]
			address := memory[ptr+3]
			memory[address] = memory[paramOne] * memory[paramTwo]
			ptr += 4
		case 99:
			return memory
		default:
			return memory
		}
		if ptr == len(memory) {
			return memory
		}
	}
}
