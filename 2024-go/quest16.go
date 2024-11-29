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

	input = ReadLines("input/q16_p2.txt")
	positions, faces = q16_parse(input)

	lens := make([]int, len(positions))
	for i := range positions {
		lens[i] = len(faces[i])
	}
	lcm := LCMSlice(lens)

	cyclePositions := make([]int, len(positions))
	totalCoins := 0
	roundCoins := []int{}
	for i := 0; i < lcm; i++ {
		for j, p := range positions {
			cyclePositions[j] = (cyclePositions[j] + p) % len(faces[j])
		}
		coins := q16_coins(faces, cyclePositions)
		totalCoins += coins
		roundCoins = append(roundCoins, coins)
	}

	m := 202420242024 / lcm
	rest := 202420242024 % lcm
	totalCoins *= m
	for i := range rest {
		totalCoins += roundCoins[i]
	}

	fmt.Println("Quest 16 Part 2:", totalCoins)

	fmt.Println("Quest 16 Part 3:", "Not implemented yet")
}

func q16_coins(faces [][]string, positions []int) int {
	symbols := map[rune]int{}
	for i, faceList := range faces {
		for c, symbol := range faceList[positions[i]] {
			if c != 1 {
				symbols[symbol]++
			}
		}
	}

	coins := 0
	for _, count := range symbols {
		if count >= 3 {
			coins += count - 2
		}
	}
	return coins
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
