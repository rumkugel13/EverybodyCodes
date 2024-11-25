package main

import "fmt"

func quest15() {
	// sample := []string{
	// 	"#####.#####",
	// 	"#.........#",
	// 	"#.######.##",
	// 	"#.........#",
	// 	"###.#.#####",
	// 	"#H.......H#",
	// 	"###########",
	// }
	input := ReadLines("input/q15_p1.txt")
	start := q15_start(input)
	distance := q15_distance(input, start, 'H')

	fmt.Println("Quest 15 Part 1:", distance * 2)
}

func q15_distance(grid []string, start Point, end byte) int {
	queue := []Point{start}
	distances := map[Point]int{start: 0}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		curDist := distances[current]

		if grid[current.row][current.col] == end {
			return curDist
		}

		for _, dir := range Directions {
			next := Point{current.row + dir.row, current.col + dir.col}
			if !InsideGrid(grid, next) || grid[next.row][next.col] == '#' {
				continue
			}

			if _, exists := distances[next]; !exists {
				distances[next] = curDist + 1
				queue = append(queue, next)
			}
		}
	}
	return 0
}

func q15_start(grid []string) Point {
	for col, char := range grid[0] {
		if char == '.' {
			return Point{0, col}
		}
	}
	return Point{}
}
