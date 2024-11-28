package main

import (
	"fmt"
	"slices"
)

func quest17() {
	// sample := []string{
	// 	"*...*",
	// 	"..*..",
	// 	".....",
	// 	".....",
	// 	"*.*..",
	// }
	input := ReadLines("input/q17_p1.txt")
	stars := q17_stars(input)
	distances := q17_distances(stars)
	slices.SortFunc(distances, func(a, b Pair) int {
		return a.distance - b.distance
	})
	
	result := q17_minimal_spanning_tree(distances) + len(stars)
	fmt.Println("Quest 17 Part 1:", result)

	input = ReadLines("input/q17_p2.txt")
	stars = q17_stars(input)
	distances = q17_distances(stars)
	slices.SortFunc(distances, func(a, b Pair) int {
		return a.distance - b.distance
	})
	result = q17_minimal_spanning_tree(distances) + len(stars)
	fmt.Println("Quest 17 Part 2:", result)

	fmt.Println("Quest 17 Part 3:", "Not implemented yet")
}

func q17_minimal_spanning_tree(pairs []Pair) int {
	parent := make(map[Point]Point)
	rank := make(map[Point]int)

	var find func(Point) Point
	find = func(p Point) Point {
		if parent[p] == (Point{}) {
			return p
		}
		if parent[p] != p {
			parent[p] = find(parent[p])
		}
		return parent[p]
	}

	union := func(p1, p2 Point) bool {
		root1 := find(p1)
		root2 := find(p2)

		if root1 == root2 {
			return false
		}

		if rank[root1] > rank[root2] {
			parent[root2] = root1
		} else if rank[root1] < rank[root2] {
			parent[root1] = root2
		} else {
			parent[root2] = root1
			rank[root1]++
		}
		return true
	}

	totalDistance := 0
	for _, pair := range pairs {
		if union(pair.start, pair.end) {
			totalDistance += pair.distance
		}
	}
	return totalDistance
}

type Pair struct {
	start, end Point
	distance   int
}

func q17_distances(stars []Point) []Pair {
	pairs := []Pair{}
	for i, start := range stars {
		for j, end := range stars {
			if i == j {
				continue
			}
			pairs = append(pairs, Pair{start, end, Distance(start, end)})
		}
	}
	return pairs
}

func q17_stars(grid []string) []Point {
	stars := []Point{}
	for row, line := range grid {
		for col, c := range line {
			if c == '*' {
				stars = append(stars, Point{row, col})
			}
		}
	}
	return stars
}
