package main

import (
	"fmt"
	"math"
	"strings"
)

func quest11() {
	// sample := []string{
	// 	"A:B,C",
	// 	"B:C,A",
	// 	"C:A",
	// }

	input := ReadLines("input/q11_p1.txt")
	conversions := q11_map(input)
	counts := q11_count(conversions, "A", 4)
	sum := SumMap(counts)

	fmt.Println("Quest 11 Part 1:", sum)

	input = ReadLines("input/q11_p2.txt")
	conversions = q11_map(input)
	counts = q11_count(conversions, "Z", 10)
	sum = SumMap(counts)
	fmt.Println("Quest 11 Part 2:", sum)

	input = ReadLines("input/q11_p3.txt")
	conversions = q11_map(input)
	minVal, maxVal := math.MaxInt64, math.MinInt64
	for key := range conversions {
		counts = q11_count(conversions, key, 20)
		sum = SumMap(counts)
		minVal = min(minVal, sum)
		maxVal = max(maxVal, sum)
	}
	fmt.Println("Quest 11 Part 3:", maxVal-minVal)
}

func q11_count(conversions map[string][]string, start string, days int) map[string]int {
	counts := map[string]int{start: 1}
	for day := 0; day < days; day++ {
		duplicate := map[string]int{}
		for key, values := range conversions {
			amount := counts[key]
			for _, value := range values {
				duplicate[value] += amount
			}
		}
		counts = duplicate
	}
	return counts
}

func q11_map(input []string) map[string][]string {
	conversions := map[string][]string{}
	for _, line := range input {
		parts := strings.Split(line, ":")
		conversions[parts[0]] = strings.Split(parts[1], ",")
	}
	return conversions
}
