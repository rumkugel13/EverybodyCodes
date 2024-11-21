package main

import (
	"fmt"
	"math"
)

func quest13() {
	// sample := []string{
	// 	"#######",
	// 	"#6769##",
	// 	"S50505E",
	// 	"#97434#",
	// 	"#######",
	// }
	input := ReadLines("input/q13_p1.txt")
	time := q13_shortest_path(input, 'S', 'E')

	fmt.Println("Quest 13 Part 1:", time)

	input = ReadLines("input/q13_p2.txt")
	time = q13_shortest_path(input, 'S', 'E')

	fmt.Println("Quest 13 Part 2:", time)

	// sample := []string{
	// 	"SSSSSSSSSSS",
	// 	"S674345621S",
	// 	"S###6#4#18S",
	// 	"S53#6#4532S",
	// 	"S5450E0485S",
	// 	"S##7154532S",
	// 	"S2##314#18S",
	// 	"S971595#34S",
	// 	"SSSSSSSSSSS",
	// }
	input = ReadLines("input/q13_p3.txt")
	time = q13_shortest_path(input, 'E', 'S')

	fmt.Println("Quest 13 Part 3:", time)
}

func q13_shortest_path(input []string, start, end byte) int {
	startPoint := FindInGrid(input, start)
	queue := []Point{startPoint}
	times := map[Point]int{startPoint: 0}
	shortestPath := math.MaxInt64

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if input[current.row][current.col] == end {
			shortestPath = min(shortestPath, times[current])
			continue
		}

		currentHeight := q13_height(input, current)

		for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := Point{current.row + dir.row, current.col + dir.col}
			if !InsideGrid(input, next) || input[next.row][next.col] == '#' {
				continue
			}

			nextHeight := q13_height(input, next)

			delta := abs(nextHeight - currentHeight)
			cost := min(10-delta, delta)
			newTime := times[current] + cost + 1

			if existingTime, exists := times[next]; !exists || newTime < existingTime {
				times[next] = newTime
				queue = append(queue, next)
			}
		}
	}

	return shortestPath
}

func q13_height(input []string, pos Point) int {
	height := int(input[pos.row][pos.col] - '0')
	if input[pos.row][pos.col] == 'E' || input[pos.row][pos.col] == 'S' {
		height = 0
	}
	return height
}
