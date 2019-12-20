package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := openFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ret := Compute(line)
		_ = ret
	}
}

func openFile(input string) (*os.File, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Compute takes in the input and gives us the processed output
func Compute(input string) string {
	return "out"
}
