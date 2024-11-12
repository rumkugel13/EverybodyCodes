package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func quest05() {
	// sample := []string{
	// 	"2 3 4 5",
	// 	"3 4 5 2",
	// 	"4 5 2 3",
	// 	"5 2 3 4",
	// }

	input := ReadLines("input/q05_p1.txt")
	cols := q5_cols(input)
	for i := 0; i < 10; i++ {
		q5_round(&cols, i)
	}
	result := q5_number(cols)

	fmt.Println("Quest 05 Part 1:", result)

	// sample := []string{
	// 	"2 3 4 5",
	// 	"6 7 8 9",
	// }

	input = ReadLines("input/q05_p2.txt")
	cols = q5_cols(input)

	shouted := map[string]int{}
	round := 0
	shout := ""
	for {
		q5_round(&cols, round)
		shout = q5_number(cols)
		if shouted[shout] == 2023 {
			break
		}
		shouted[shout]++
		round++
	}

	num, _ := strconv.Atoi(shout)
	fmt.Println("Quest 05 Part 2:", num*(round+1))

	// sample := []string{
	// 	"2 3 4 5",
	// 	"6 7 8 9",
	// }

	input = ReadLines("input/q05_p3.txt")
	cols = q5_cols(input)
	shouted = map[string]int{}
	max := 0
	for {
		q5_round(&cols, round)
		shout = q5_number(cols)
		num, _ := strconv.Atoi(shout)
		shouted[shout]++
		round++
		if num > max {
			max = num
		}
		// note: arbitrary number, adjust if result is wrong
		if round > 2000000 {
			break
		}
	}

	fmt.Println("Quest 05 Part 3:", max)
}

func q5_cols(input []string) [][]int {
	cols := make([][]int, 4)
	for row := 0; row < len(input); row++ {
		split := strings.Fields(input[row])
		for col := 0; col < len(split); col++ {
			val, _ := strconv.Atoi(split[col])
			cols[col] = append(cols[col], val)
		}
	}
	return cols
}

func q5_round(cols *[][]int, round int) {
	clapper := (*cols)[round%len(*cols)][0]
	(*cols)[round%len(*cols)] = (*cols)[round%len(*cols)][1:]

	column := &(*cols)[(round+1)%len(*cols)]
	pos := (clapper - 1) % (len(*column) * 2)

	if len(*column) < pos {
		pos = Mod(clapper, len(*column)) + 1
	}

	*column = slices.Insert(*column, pos, clapper)
}

func q5_number(cols [][]int) string {
	var number string
	for i := 0; i < len(cols); i++ {
		char := strconv.Itoa(cols[i][0])
		number += char
	}
	return number
}
