package main

import (
	"os"
	"strings"
)

/// ReadLines reads the contents of the specified file and returns a slice of strings,
/// where each string represents a line of text from the file.
///
/// If an error occurs while reading the file, the function will panic with the error.
func ReadLines(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}

/// Reverse returns a new string with the characters of the input string in reverse order.
///
/// This function takes a string as input and returns a new string with the same
/// characters but in reverse order. It is useful for reversing the order of
/// characters in a string.
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


/// Duplicate creates a deep copy of the given 2D slice.
///
/// The function takes a 2D slice of any type T as input and returns a new 2D slice
/// with the same dimensions and elements as the input slice. The elements are
/// copied by value, so the new slice is independent of the original.
func Duplicate[T any](grid [][]T) [][]T {
	duplicate := make([][]T, len(grid))
	for i := range grid {
		duplicate[i] = make([]T, len(grid[i]))
		copy(duplicate[i], grid[i])
	}
	return duplicate
}
