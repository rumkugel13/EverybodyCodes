package main

import (
	"os"
	"strings"
)

func ReadLines(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

type Point struct {
	row, col int
}

func Duplicate[T any](grid [][]T) [][]T {
	duplicate := make([][]T, len(grid))
	for i := range grid {
		duplicate[i] = make([]T, len(grid[i]))
		copy(duplicate[i], grid[i])
	}
	return duplicate
}

func DuplicateMap[T comparable, U any](source map[T]U) map[T]U {
	duplicate := make(map[T]U)
	for key, value := range source {
		duplicate[key] = value
	}
	return duplicate
}

func SumMap[T comparable, U int|float32|float64](source map[T]U) U {
	var sum U
	for _, value := range source {
		sum += value
	}
	return sum
}

func Mod(a, n int) int {
	return ((a % n) + n) % n
}