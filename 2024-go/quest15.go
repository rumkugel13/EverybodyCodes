package main

import (
	"fmt"
	"math"
	"strconv"
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
	pois, _ := q15_pois(input)
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
	result := q15_collect(start, start, pois, unique, q15_distances(input, start, pois), map[string]int{})

	fmt.Println("Quest 15 Part 2:", result)

	// input = ReadLines("input/q15_p3.txt")
	// start = q15_start(input)
	// pois, unique = q15_pois(input)
	// result = q15_collect(start, start, pois, unique, q15_distances(input, start, pois), map[string]int{})

	fmt.Println("Quest 15 Part 3:", "Not implemented yet")
}

func q15_collect(current, start Point, pois map[Point]byte, herbs uint64, distances map[int]int, collectStash map[string]int) int {
	cacheKey := strconv.FormatUint(herbs, 10) + "-" + strconv.Itoa(current.row*256+current.col)
	if d, ok := collectStash[cacheKey]; ok {
		return d
	}

	minDist := math.MaxInt32
	for herb := 'A'; herb <= 'Z'; herb++ {
		idx := (herb - 'A')
		if (herbs & (1 << idx)) == 0 {
			continue
		}

		herbs &= ^(1 << idx)

		for poi, h := range pois {
			if h == byte(herb) {
				distKey := q15_hash(current, poi)
				distance := distances[distKey]

				if herbs == 0 {
					distance += distances[q15_hash(poi, start)]
				} else {
					distance += q15_collect(poi, start, pois, herbs, distances, collectStash)
				}
				minDist = min(minDist, distance)
			}
		}

		herbs |= (1 << idx)
	}

	collectStash[cacheKey] = minDist
	return minDist
}

func q15_pois(grid []string) (map[Point]byte, uint64) {
	pois := map[Point]byte{}
	unique := uint64(0)
	for row, line := range grid {
		for col, char := range line {
			if char != '#' && char != '.' && char != '~' {
				pois[Point{row, col}] = byte(char)
				unique |= (1 << (char - 'A'))
			}
		}
	}
	return pois, unique
}

func q15_hash(a, b Point) int {
	if a.row < b.row || (a.row == b.row && a.col < b.col) {
		return (a.row*256+a.col)<<16 | (b.row*256 + b.col)
	} else {
		return (b.row*256+b.col)<<16 | (a.row*256 + a.col)
	}
}

func q15_distances(grid []string, start Point, pois map[Point]byte) map[int]int {
	distances := map[int]int{}
	pois[start] = 0
	for poi := range pois {
		for poi2 := range pois {
			if poi != poi2 {
				distKey := q15_hash(poi, poi2)
				if _, ok := distances[distKey]; !ok {
					distances[distKey] = q15_distance(grid, poi, poi2)
				}
			}
		}
	}
	return distances
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
