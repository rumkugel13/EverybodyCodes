package main

import "fmt"

func quest19() {
	// sample := []string{
	// 	"LR",
	// 	"",
	// 	">-IN-",
	// 	"-----",
	// 	"W---<",
	// }
	input := ReadLines("input/q19_p1.txt")
	key := input[0]
	cells := q19_cells(input[2:])
	iteration := 0
	for row := 1; row < len(cells)-1; row++ {
		for col := 1; col < len(cells[0])-1; col++ {
			dir := key[iteration%len(key)]
			iteration++
			q19_rotate(cells, row, col, dir)
		}
	}

	result := q19_message(cells)
	fmt.Println("Quest 19 Part 1:", result)

	input = ReadLines("input/q19_p2.txt")
	key = input[0]
	cells = q19_cells(input[2:])
	for range 100 {
		iteration := 0
		for row := 1; row < len(cells)-1; row++ {
			for col := 1; col < len(cells[0])-1; col++ {
				dir := key[iteration%len(key)]
				iteration++
				q19_rotate(cells, row, col, dir)
			}
		}
	}

	result = q19_message(cells)
	fmt.Println("Quest 19 Part 2:", result)

	fmt.Println("Quest 19 Part 3:", "Not implemented yet")
}

func q19_message(cells [][]byte) string {
	result := ""
	start := false
	for row := 0; row < len(cells); row++ {
		for col := 0; col < len(cells[0]); col++ {
			if cells[row][col] == '<' {
				start = false
			}
			if start {
				result += string(cells[row][col])
			}
			if cells[row][col] == '>' {
				start = true
			}
		}
	}
	return result
}

func q19_rotate(cells [][]byte, row, col int, dir byte) {
	switch dir {
	case 'L':
		temp := cells[row-1][col-1]
		cells[row-1][col-1] = cells[row-1][col]
		cells[row-1][col] = cells[row-1][col+1]
		cells[row-1][col+1] = cells[row][col+1]
		cells[row][col+1] = cells[row+1][col+1]
		cells[row+1][col+1] = cells[row+1][col]
		cells[row+1][col] = cells[row+1][col-1]
		cells[row+1][col-1] = cells[row][col-1]
		cells[row][col-1] = temp
	case 'R':
		temp := cells[row-1][col-1]
		cells[row-1][col-1] = cells[row][col-1]
		cells[row][col-1] = cells[row+1][col-1]
		cells[row+1][col-1] = cells[row+1][col]
		cells[row+1][col] = cells[row+1][col+1]
		cells[row+1][col+1] = cells[row][col+1]
		cells[row][col+1] = cells[row-1][col+1]
		cells[row-1][col+1] = cells[row-1][col]
		cells[row-1][col] = temp
	}
}

func q19_cells(grid []string) [][]byte {
	cells := make([][]byte, len(grid))
	for i, row := range grid {
		cells[i] = []byte(row)
	}
	return cells
}
