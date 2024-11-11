package main

import (
	"fmt"
	"math"
	"strconv"
)

func quest04() {
	// sample := []string{"3", "4", "7", "8"}
	input := ReadLines("input/q04_p1.txt")

	nails, min := q4_nails(input)

	sum := 0
	for _, nail := range nails {
		sum += nail - min
	}

	fmt.Println("Quest 04 Part 1:", sum)

	input = ReadLines("input/q04_p2.txt")
	nails, min = q4_nails(input)

	sum = 0
	for _, nail := range nails {
		sum += nail - min
	}

	fmt.Println("Quest 04 Part 2:", sum)

	input = ReadLines("input/q04_p3.txt")
	nails, _ = q4_nails(input)

	min = math.MaxInt
	for i := 0; i < len(nails); i++ {
		sum = 0
		for _, nail := range nails {
			sum += abs(nail - nails[i])
		}
		if sum < min {
			min = sum
		}
	}

	fmt.Println("Quest 04 Part 3:", min)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func q4_nails(input []string) ([]int, int) {
	nails := make([]int, len(input))
	min := math.MaxInt
	for i := 0; i < len(input); i++ {
		nails[i], _ = strconv.Atoi(input[i])
		if nails[i] < min {
			min = nails[i]
		}
	}
	return nails, min
}
