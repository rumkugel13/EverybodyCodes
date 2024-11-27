package main

import (
	"fmt"
)

func quest16() {
	// sample := []string{
	// 	"1,2,3",
	// 	"",
	// 	"^_^ -.- ^,-",
	// 	">.- ^_^ >.<",
	// 	"-_- -.- >.<",
	// 	"    -.^ ^_^",
	// 	"    >.>    ",
	// }
	input := ReadLines("input/q16_p1.txt")
	positions, faces := q16_parse(input)

	finalPositions := make([]int, len(positions))
	for i, p := range positions {
		finalPositions[i] = (0 + p*100) % len(faces[i])
	}

	result := ""
	for i, p := range finalPositions {
		result += faces[i][p] + " "
	}
	fmt.Println("Quest 16 Part 1:", result)
}

func q16_parse(input []string) ([]int, [][]string) {
	positions := CommaSepToIntArr(input[0])
	faces := [][]string{}
	for range positions {
		faces = append(faces, []string{})
	}

	for _, line := range input[2:] {
		for i := 0; i < len(line); i += 4 {
			if line[i] == ' ' {
				continue
			}
			faces[i/4] = append(faces[i/4], string(line[i:i+3]))
		}
	}
	return positions, faces
}
