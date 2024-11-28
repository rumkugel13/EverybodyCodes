package main

import (
	"fmt"
	"math"
)

func quest18() {
	// sample := []string{
	// 	"##########",
	// 	"..#......#",
	// 	"#.P.####P#",
	// 	"#.#...P#.#",
	// 	"##########",
	// }
	input := ReadLines("input/q18_p1.txt")
	start := Point{1, 0}

	result := q18_times(input, []Point{start}, q18_total_palmtrees(input))
	fmt.Println("Quest 18 Part 1:", result[len(result)-1])

	input = ReadLines("input/q18_p2.txt")
	starts := []Point{{1, 0}, {len(input) - 2, len(input[0]) - 1}}
	result = q18_times(input, starts, q18_total_palmtrees(input))
	fmt.Println("Quest 18 Part 2:", result[len(result)-1])

	input = ReadLines("input/q18_p3.txt")
	minSum := math.MaxInt32
	totalPalmtrees := q18_total_palmtrees(input)
	for r, row := range input {
		for c, cell := range row {
			if cell == '.' {
				times := q18_times(input, []Point{{r, c}}, totalPalmtrees)
				sum := SumSlice(times)
				if sum < minSum {
					minSum = sum
				}
			}
		}
	}
	fmt.Println("Quest 18 Part 3:", minSum)
}

func q18_total_palmtrees(grid []string) int {
	total := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == 'P' {
				total++
			}
		}
	}
	return total
}

func q18_times(grid []string, starts []Point, totalP int) []int {
	queue := []Point{}
	visited := map[Point]int{}
	distances := []int{}

	for _, start := range starts {
		visited[start] = 0
		queue = append(queue, start)
	}

	reachedP := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		distance := visited[current]
		if grid[current.row][current.col] == 'P' {
			reachedP++
			distances = append(distances, distance)
			if reachedP == totalP {
				return distances
			}
		}

		for _, dir := range Directions {
			next := Point{current.row + dir.row, current.col + dir.col}
			if !InsideGrid(grid, next) || grid[next.row][next.col] == '#' {
				continue
			}

			if _, exists := visited[next]; !exists {
				visited[next] = distance + 1
				queue = append(queue, next)
			}
		}
	}
	return distances
}
