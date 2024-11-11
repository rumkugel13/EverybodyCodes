package main

import "fmt"

func quest03() {
	// sample := []string{
	// 	"..........",
	// 	"..###.##..",
	// 	"...####...",
	// 	"..######..",
	// 	"..######..",
	// 	"...####...",
	// 	".........."}
	input := getLines("input/q03_p1.txt")
	grid := q3_grid(input)
	grid = q3_dig(grid)
	count := q3_count(grid)

	fmt.Println("Quest 03 Part 1:", count)

	input = getLines("input/q03_p2.txt")
	grid = q3_grid(input)
	grid = q3_dig(grid)
	count = q3_count(grid)

	fmt.Println("Quest 03 Part 2:", count)

	input = getLines("input/q03_p3.txt")
	grid = q3_grid(input)
	grid = q3_digdiagonal(grid)
	count = q3_count(grid)

	fmt.Println("Quest 03 Part 3:", count)
}

func q3_grid(input []string) [][]int {
	grid := make([][]int, len(input))
	for row := 0; row < len(input); row++ {
		grid[row] = make([]int, len(input[row]))
		for col := 0; col < len(input[row]); col++ {
			if input[row][col] == '#' {
				grid[row][col] = 1
			}
		}
	}
	return grid
}

func q3_dig(grid [][]int) [][]int {
	changed := true
	for changed {
		changed = false
		duplicate := make([][]int, len(grid))
		for i := range grid {
			duplicate[i] = make([]int, len(grid[i]))
			copy(duplicate[i], grid[i])
		}

		for row := 1; row < len(grid)-1; row++ {
			for col := 1; col < len(grid[row])-1; col++ {
				num := grid[row][col]
				if num == 0 {
					continue
				}
				if grid[row-1][col] == num &&
					grid[row][col-1] == num &&
					grid[row+1][col] == num &&
					grid[row][col+1] == num {
					duplicate[row][col] = num + 1
					changed = true
				}
			}
		}

		grid = duplicate
	}
	return grid
}

func q3_digdiagonal(grid [][]int) [][]int {
	changed := true
	for changed {
		changed = false
		duplicate := make([][]int, len(grid))
		for i := range grid {
			duplicate[i] = make([]int, len(grid[i]))
			copy(duplicate[i], grid[i])
		}
		for row := 1; row < len(grid)-1; row++ {
			for col := 1; col < len(grid[row])-1; col++ {
				num := grid[row][col]
				if num == 0 {
					continue
				}
				if grid[row-1][col] == num &&
					grid[row][col-1] == num &&
					grid[row+1][col] == num &&
					grid[row][col+1] == num &&
					grid[row-1][col-1] == num &&
					grid[row-1][col+1] == num &&
					grid[row+1][col-1] == num &&
					grid[row+1][col+1] == num {
					duplicate[row][col] = num + 1
					changed = true
				}
			}
		}
		grid = duplicate
	}
	return grid
}

func q3_count(grid [][]int) int {
	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			count += grid[row][col]
		}
	}
	return count
}
