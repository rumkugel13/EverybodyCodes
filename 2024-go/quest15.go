package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"maps"
)

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
	pois,_ := q15_pois(input)
	minDist := math.MaxInt32
	for poi := range pois {
		distance := q15_distance(input, start, poi)
		minDist = min(minDist, distance)
	}

	fmt.Println("Quest 15 Part 1:", minDist*2)

	// sample := []string{
	// 	"##########.##########",
	// 	"#...................#",
	// 	"#.###.##.###.##.#.#.#",
	// 	"#..A#.#..~~~....#A#.#",
	// 	"#.#...#.~~~~~...#.#.#",
	// 	"#.#.#.#.~~~~~.#.#.#.#",
	// 	"#...#.#.B~~~B.#.#...#",
	// 	"#...#....BBB..#....##",
	// 	"#C............#....C#",
	// 	"#####################",
	// }
	input = ReadLines("input/q15_p2.txt")
	start = q15_start(input)
	pois, unique := q15_pois(input)
	result := q15_collect(input, start, start, pois, unique, map[string]int{}, map[string]int{})

	fmt.Println("Quest 15 Part 2:", result)

	// input = ReadLines("input/q15_p3.txt")
	// start = q15_start(input)
	// pois, unique = q15_pois(input)
	// result = q15_collect(input, start, start, pois, unique, map[string]int{}, map[string]int{})

	// fmt.Println("Quest 15 Part 3:", result)
}

func q15_collect(grid []string, current, start Point, pois map[Point]byte, herbs map[byte]byte, distStash, collectStash map[string]int) int {
	if len(herbs) == 0 {
		distKey := strconv.Itoa(current.row) + "-" + strconv.Itoa(current.col) + "-" + strconv.Itoa(start.row) + "-" + strconv.Itoa(start.col)
		if d, ok := distStash[distKey]; ok {
			return d
		}
		dist := q15_distance(grid, current, start)
		distStash[distKey] = dist
		return dist
	}

	herbsKey := string(slices.Collect(maps.Keys(herbs)))
	cacheKey := herbsKey + "-" + strconv.Itoa(current.row) + "-" + strconv.Itoa(current.col)
	if d, ok := collectStash[cacheKey]; ok {
		return d
	}

	minDist := math.MaxInt32
	for _, herb := range herbs {
		delete(herbs, herb)

		for poi, h := range pois {
			if h == herb {
				distKey := strconv.Itoa(current.row) + "-" + strconv.Itoa(current.col) + "-" + strconv.Itoa(poi.row) + "-" + strconv.Itoa(poi.col)
				poiDist := 0
				if d, ok := distStash[distKey]; ok {
					poiDist = d
				} else {
					poiDist = q15_distance(grid, current, poi)
					distStash[distKey] = poiDist
				}

				collectDist := q15_collect(grid, poi, start, pois, herbs, distStash, collectStash)
				distance := poiDist + collectDist
				minDist = min(minDist, distance)
			}
		}
		herbs[herb] = herb
	}
	
	collectStash[cacheKey] = minDist
	return minDist
}

func q15_pois(grid []string) (map[Point]byte, map[byte]byte) {
	pois := map[Point]byte{}
	unique := map[byte]byte{}
	for row, line := range grid {
		for col, char := range line {
			if char != '#' && char != '.' && char != '~' {
				pois[Point{row, col}] = byte(char)
				unique[byte(char)] = byte(char)
			}
		}
	}
	return pois, unique
}

func q15_distance(grid []string, start Point, end Point) int {
	queue := []Point{start}
	distances := map[Point]int{start: 0}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		curDist := distances[current]

		if current == end {
			return curDist
		}

		for _, dir := range Directions {
			next := Point{current.row + dir.row, current.col + dir.col}
			if !InsideGrid(grid, next) || grid[next.row][next.col] == '#' || grid[next.row][next.col] == '~' {
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
