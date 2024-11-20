package main

import "fmt"

func quest12() {
	// sample := []string{
	// 	".............",
	// 	".C...........",
	// 	".B......T....",
	// 	".A......T.T..",
	// 	"=============",
	// }
	input := ReadLines("input/q12_p1.txt")
	catapults, targets := q12_parse(input)

	sum := q12_calc(catapults, targets, 10)
	fmt.Println("Quest 12 Part 1:", sum)

	// sample := []string{
	// 	".............",
	// 	".C...........",
	// 	".B......H....",
	// 	".A......T.H..",
	// 	"=============",
	// }
	input = ReadLines("input/q12_p2.txt")
	catapults, targets = q12_parse(input)
	sum = q12_calc(catapults, targets, 50)

	fmt.Println("Quest 12 Part 2:", sum)

	fmt.Println("Quest 12 Part 3:", "Not implemented yet")
}

func q12_calc(catapults []Point, targets map[Point]byte, maxRange int) int {
	sum := 0
	for target, rock := range targets {
		multiplier := 1
		if rock == 'H' {
			multiplier = 2
		}
	outer:
		for _, catapult := range catapults {
			for power := range maxRange {
				if q12_hit(catapult, target, power+1) {
					sum += catapult.row * (power + 1) * multiplier
					break outer
				}
			}
		}
	}
	return sum
}

func q12_hit(catapult Point, target Point, power int) bool {
	projectile := Point(catapult)

	projectile.row += power
	projectile.col += power

	projectile.col += power

	for projectile.row > 1 {
		projectile.row--
		projectile.col++
		if projectile == target {
			return true
		}
	}

	return false
}

func q12_parse(input []string) ([]Point, map[Point]byte) {
	catapults := []Point{}
	targets := map[Point]byte{}
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			if input[row][col] == 'A' || input[row][col] == 'B' || input[row][col] == 'C' {
				catapults = append(catapults, Point{len(input) - 1 - row, col})
			} else if input[row][col] == 'T' || input[row][col] == 'H' {
				targets[Point{len(input) - 1 - row, col}] = input[row][col]
			}
		}
	}
	return catapults, targets
}
