package main

import (
	"fmt"
	"strings"
)

func quest01() {
	// sample01 := "ABBAC"
	input := ReadLines("input/q01_p1.txt")[0]

	sum := 0
	for i := 0; i < len(input); i++ {
		if input[i] == 'A' {
			sum += 0
		} else if input[i] == 'B' {
			sum += 1
		} else if input[i] == 'C' {
			sum += 3
		}
	}

	fmt.Println("Quest 01 Part 1:", sum)

	input = ReadLines("input/q01_p2.txt")[0]
	sum = 0
	potions := map[byte]int{
		'A': 0,
		'B': 1,
		'C': 3,
		'D': 5,
		'x': 0,
	}

	for i := 0; i < len(input); i += 2 {
		if input[i] == 'x' || input[i+1] == 'x' {
			sum += potions[input[i]] + potions[input[i+1]]
		} else {
			sum += potions[input[i]] + potions[input[i+1]] + 2
		}
	}

	fmt.Println("Quest 01 Part 2:", sum)

	// sample03 := "xBxAAABCDxCC"
	input = ReadLines("input/q01_p3.txt")[0]
	sum = 0

	for i := 0; i < len(input); i += 3 {
		count := strings.Count(input[i:i+3], "x")

		if count == 2 {
			sum += potions[input[i]] + potions[input[i+1]] + potions[input[i+2]]
		} else if count == 1 {
			sum += potions[input[i]] + potions[input[i+1]] + potions[input[i+2]] + 2
		} else if count == 0 {
			sum += potions[input[i]] + potions[input[i+1]] + potions[input[i+2]] + 6
		}
	}
	fmt.Println("Quest 01 Part 3:", sum)
}
