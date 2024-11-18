package main

import "fmt"

func quest10() {
	// sample := []string{
	// 	"**PCBS**",
	// 	"**RLNW**",
	// 	"BV....PT",
	// 	"CR....HZ",
	// 	"FL....JW",
	// 	"SG....MN",
	// 	"**FTZV**",
	// 	"**GMJH**",
	// }
	input := ReadLines("input/q10_p1.txt")
	result := q10_runes(input)

	fmt.Println("Quest 10 Part 1:", result)

	input = ReadLines("input/q10_p2.txt")

	power := 0
	for row := 0; row < len(input); row += 9 {
		for col := 0; col < len(input[row]); col += 9 {
			grid := []string{}
			for i := 0; i < 8; i++ {
				grid = append(grid, input[row+i][col:col+8])
			}
			runes := q10_runes(grid)
			power += q10_power(runes)
		}
	}

	fmt.Println("Quest 10 Part 2:", power)

	fmt.Println("Quest 10 Part 3:", "Not implemented yet")
}

func q10_runes(grid []string) string {
	result := ""
	for row := 2; row < len(grid)-2; row++ {
		for col := 2; col < len(grid[row])-2; col++ {
			row_symbols := grid[row][0:2] + grid[row][len(grid[row])-2:]
			col_symbols := string(grid[0][col]) + string(grid[1][col]) + string(grid[len(grid)-2][col]) + string(grid[len(grid)-1][col])
		outer:
			for i := 0; i < len(row_symbols); i++ {
				for j := 0; j < len(col_symbols); j++ {
					if row_symbols[i] == col_symbols[j] {
						result += string(row_symbols[i])
						break outer
					}
				}
			}
		}
	}
	return result
}

func q10_power(runes string) int {
	result := 0
	for i, r := range runes {
		result += int(r-'A'+1) * (i + 1)
	}
	return result
}
