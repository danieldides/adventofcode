package main

import (
	"io/ioutil"
	"log"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	ret := Compute(string(file))
	_ = ret
}

// Compute takes in the input and gives us the processed output
func Compute(input string) string {
	return "out"
}
