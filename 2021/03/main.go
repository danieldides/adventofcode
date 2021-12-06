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
	log.Printf("Part 2: %v\n", DoTwo(in))
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
			// e += strconv.Itoa(0)
			e += "0"
		}
		if in[i] == 0 {
			// e += strconv.Itoa(1)
			e += "1"
		}
	}

	return g, e
}

func RemoveIndex(s [][]int, index int) [][]int {
	return append(s[:index], s[index+1:]...)
}

func filterReportsOxy(reports [][]int, tie int) []int {
	var result []int

	mask := findMostSignificantBit(reports, tie)

	var s [][]int

	// Oxygen
	for i := 0; i < len(mask); i++ {
		mask = findMostSignificantBit(reports, tie)

		for _, report := range reports {
			// fmt.Printf("Comparing %v of mask %v to %v\n", i, mask, report)
			if report[i] == mask[i] {
				// fmt.Printf("\tSaving: %v\n", report)
				// s = RemoveIndex(reports, i)
				s = append(s, report)
			}
		}

		reports = s
		s = nil
		// fmt.Printf("Iteration %v. Reports (%v): %v\n", i, len(reports), reports)
		if len(reports) == 1 {
			// fmt.Println("never gonna give you up", tie)
			result = reports[0]
			break
		}
	}

	return result

}

func filterReportsCo2(reports [][]int, tie int) []int {
	var result []int

	mask := findLeastSignificantBit(reports, tie)

	var s [][]int

	// Oxygen
	for i := 0; i < len(mask); i++ {
		mask = findLeastSignificantBit(reports, tie)

		for _, report := range reports {
			// fmt.Printf("Comparing %v of mask %v to %v\n", i, mask, report)
			if report[i] == mask[i] {
				// fmt.Printf("\tSaving: %v\n", report)
				// s = RemoveIndex(reports, i)
				s = append(s, report)
			}
		}

		reports = s
		s = nil
		// fmt.Printf("Iteration %v. Reports (%v): %v\n", i, len(reports), reports)
		if len(reports) == 1 {
			// fmt.Println("never gonna give you up", tie)
			result = reports[0]
			break
		}
	}

	return result

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

func findMostSignificantBit(rr [][]int, tie int) []int {
	var result []int

	// Actual algo
	max := findMax(rr)
	for i := 0; i < max; i++ {
		var zeroes, ones, set int
		for _, r := range rr {
			if r[i] == 0 {
				zeroes++
			}
			if r[i] == 1 {
				ones++
			}

			// UGH this is horrible
			if zeroes > ones {
				set = 0
			}

			if ones > zeroes {
				set = 1
			}

			if ones == zeroes {
				set = tie
			}

		}
		result = append(result, set)

	}

	return result
}

func findLeastSignificantBit(rr [][]int, tie int) []int {
	var result []int

	// Actual algo
	max := findMax(rr)
	for i := 0; i < max; i++ {
		var zeroes, ones, set int
		for _, r := range rr {
			if r[i] == 0 {
				zeroes++
			}
			if r[i] == 1 {
				ones++
			}

			// UGH this is horrible
			if zeroes > ones {
				set = 1
			}

			if ones > zeroes {
				set = 0
			}

			if ones == zeroes {
				set = tie
			}

		}
		result = append(result, set)

	}

	return result
}

func DoTwo(in []string) int {
	var reports [][]int

	for _, s := range in {
		l := parseLine(s)
		if len(l) > 0 {
			reports = append(reports, l)
		}
	}

	oxy := filterReportsOxy(reports, 1)
	co2 := filterReportsCo2(reports, 0)

	// fmt.Println("Oxygen: ", oxy)
	// fmt.Println("CO2: ", co2)

	// fmt.Println(sliceToBinary(oxy))

	g, _ := sliceToBinary(oxy)
	e, _ := sliceToBinary(co2)

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
