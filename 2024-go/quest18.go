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

	result := q18_times(input, []Point{start}, len(q18_palmtrees(input)))
	fmt.Println("Quest 18 Part 1:", result[len(result)-1])

	input = ReadLines("input/q18_p2.txt")
	starts := []Point{{1, 0}, {len(input) - 2, len(input[0]) - 1}}
	result = q18_times(input, starts, len(q18_palmtrees(input)))
	fmt.Println("Quest 18 Part 2:", result[len(result)-1])

	input = ReadLines("input/q18_p3.txt")
	palmtress := q18_palmtrees(input)
	distances := map[Point]int{}
	for _, tree := range palmtress {
		times := q18_all_times(input, tree)
		for point, time := range times {
			distances[point] += time
		}
	}

	minSum := math.MaxInt32
	for point, distance := range distances {
		if input[point.row][point.col] == '.' && distance < minSum {
			minSum = distance
		}
	}

	fmt.Println("Quest 18 Part 3:", minSum)
}

func q18_palmtrees(grid []string) []Point {
	palmtrees := []Point{}
	for r, row := range grid {
		for c, cell := range row {
			if cell == 'P' {
				palmtrees = append(palmtrees, Point{r, c})
			}
		}
	}
	return palmtrees
}

func q18_all_times(grid []string, start Point) map[Point]int {
	visited := map[Point]int{}
	queue := []Point{start}
	visited[start] = 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		distance := visited[current]
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
	return visited
}

func q18_times(grid []string, starts []Point, totalP int) []int {
	queue := []Point{}
	visited := map[Point]int{}
	distances := []int{}

	for _, start := range starts {
		visited[start] = 0
		queue = append(queue, start)
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		distance := visited[current]
		if grid[current.row][current.col] == 'P' {
			distances = append(distances, distance)
			if len(distances) == totalP {
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
