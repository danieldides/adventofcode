package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func main() {
	// File loading boilerplate
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var in []string

	lines := strings.Split(string(file), "\n")

	for _, line := range lines {
		in = append(in, line)
	}

	log.Printf("Part 1: %v\n", DoOne(in))
	// log.Printf("Part 2: %v\n", DoTwo(in))
}

func parseLine(in string) []int {
	var out []int
	s := strings.Split(in, "")
	for _, c := range s {
		out = append(out, Atoi(c))
	}
	return out
}

func findMax(in [][]int) int {
	max := 0
	for _, s := range in {
		if len(s) > max {
			max = len(s)
		}
	}

	return max
}

func sliceToBinary(in []int) (string, string) {
	var g, e string
	for i := range in {
		g += strconv.Itoa(in[i])
	}
	for i := range in {
		if in[i] == 1 {
			e += strconv.Itoa(0)
		}
		if in[i] == 0 {
			e += strconv.Itoa(1)
		}
	}

	return g, e
}

func DoOne(in []string) int {
	var reports [][]int

	for _, s := range in {
		l := parseLine(s)
		if len(l) > 0 {
			reports = append(reports, l)
		}
	}

	var common []int

	// Actual algo
	max := findMax(reports)
	for i := 0; i < max; i++ {
		sum := 0
		for _, report := range reports {
			sum += report[i]
		}
		common = append(common, sum)
	}

	var result []int

	for _, bit := range common {
		b := 0
		if bit >= len(reports)/2 {
			b = 1
		}
		result = append(result, b)
	}

	g, e := sliceToBinary(result)

	gamma, err := strconv.ParseInt(g, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilon, err := strconv.ParseInt(e, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(gamma * epsilon)

}

func DoTwo(in []string) int {
	return 0
}
