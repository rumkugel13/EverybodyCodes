package main

import (
	"fmt"
	// "math"
	"strconv"
	"strings"
)

func quest14() {
	// sample := []string{"U5,R3,D2,L5,U4,R5,D2"}
	input := ReadLines("input/q14_p1.txt")
	growth_plan := strings.Split(input[0], ",")
	segments := map[Point3]int{}
	q14_grow(growth_plan, segments)

	maxHeight := q14_height(segments)

	fmt.Println("Quest 14 Part 1:", maxHeight)

	// sample := []string{
	// 	"U5,R3,D2,L5,U4,R5,D2",
	// 	"U6,L1,D2,R3,U2,L1",
	// }
	input = ReadLines("input/q14_p2.txt")
	segments = map[Point3]int{}
	for _, line := range input {
		growth_plan = strings.Split(line, ",")
		q14_grow(growth_plan, segments)
	}

	fmt.Println("Quest 14 Part 2:", len(segments))

	// sample := []string{
	// 	"U5,R3,D2,L5,U4,R5,D2",
	// 	"U6,L1,D2,R3,U2,L1",
	// }
	// sample2 := []string{
	// 	"U20,L1,B1,L2,B1,R2,L1,F1,U1",
	// 	"U10,F1,B1,R1,L1,B1,L1,F1,R2,U1",
	// 	"U30,L2,F1,R1,B1,R1,F2,U1,F1",
	// 	"U25,R1,L2,B1,U1,R2,F1,L2",
	// 	"U16,L1,B1,L1,B3,L1,B1,F1",
	// }
	// input = ReadLines("input/q14_p3.txt")
	// segments = map[Point3]int{}
	// leaves := map[Point3]int{}
	// for _, line := range input {
	// 	growth_plan = strings.Split(line, ",")
	// 	leaf := q14_grow(growth_plan, segments)
	// 	leaves[leaf] = 0
	// }

	// trunk := Point3{0, 0, 0}
	// segments[trunk] = 1
	// minMurk := math.MaxInt32
	// for trunk.y < q14_height(segments) {
	// 	murk := q14_distance_to_leaves(trunk, leaves)
	// 	if murk < minMurk {
	// 		minMurk = murk
	// 		// fmt.Println(murk, trunk)
	// 	}
	// 	trunk.y++
	// }

	// fmt.Println("Quest 14 Part 3:", minMurk)
}

// func q14_distance_to_leaves(segment Point3, leaves map[Point3]int) int {
// 	distance := 0
// 	for leaf := range leaves {
// 		distance += q14_distance(segment, leaf)
// 	}
// 	return distance
// }

// func q14_distance(p1, p2 Point3) int {
// 	return abs(p1.x-p2.x) + abs(p1.y-p2.y) + abs(p1.z-p2.z)
// }

func q14_height(segments map[Point3]int) int {
	height := 0
	for point := range segments {
		if point.y > height {
			height = point.y
		}
	}
	return height
}

func q14_grow(plan []string, segments map[Point3]int) Point3 {
	pos := Point3{0, 0, 0}

	for _, step := range plan {
		dir := step[0]
		dist, _ := strconv.Atoi(step[1:])
		for i := 0; i < dist; i++ {
			switch dir {
			case 'U':
				pos.y++
			case 'D':
				pos.y--
			case 'R':
				pos.x++
			case 'L':
				pos.x--
			case 'F':
				pos.z++
			case 'B':
				pos.z--
			}
			segments[pos] = 1
		}
	}
	return pos
}

type Point3 struct {
	x, y, z int
}
