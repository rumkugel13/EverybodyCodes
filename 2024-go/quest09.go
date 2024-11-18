package main

import (
	"fmt"
	"math"
	"strconv"
)

func quest09() {
	// sample := []string{"2", "4", "7", "16"}
	input := ReadLines("input/q09_p1.txt")

	stamps_available := []int{10, 5, 3, 1}
	sum := 0

	for _, line := range input {
		brightness, _ := strconv.Atoi(line)
		for i := 0; i < len(stamps_available); i++ {
			for brightness-stamps_available[i] >= 0 {
				sum++
				brightness -= stamps_available[i]
			}
		}
	}

	fmt.Println("Quest 09 Part 1:", sum)

	// sample := []string{"33", "41", "55", "99"}
	input = ReadLines("input/q09_p2.txt")

	stamps_available = []int{30, 25, 24, 20, 16, 15, 10, 5, 3, 1}
	sum = 0
	for _, line := range input {
		brightness, _ := strconv.Atoi(line)
		sum += q9_beetles(brightness, stamps_available, map[string]int{})
	}
	fmt.Println("Quest 09 Part 2:", sum)

	fmt.Println("Quest 09 Part 3:", "Not implemented yet")
}

func q9_beetles(brightness int, stamps_available []int, stash map[string]int) int {
	if brightness == 0 {
		return 0
	}

	if brightness < 0 || len(stamps_available) == 0 {
		return math.MaxInt32
	}

	hash := q9_hash(brightness, stamps_available)
	if num, found := stash[hash]; found {
		return num
	}

	withFirstStamp := q9_beetles(brightness-stamps_available[0], stamps_available, stash) + 1
	withoutFirstStamp := q9_beetles(brightness, stamps_available[1:], stash)
	stash[hash] = min(withFirstStamp, withoutFirstStamp)
	return stash[hash]
}

func q9_hash(brightness int, stamps []int) string {
	hash := strconv.Itoa(brightness) + "_"
	for _, stamp := range stamps {
		hash += strconv.Itoa(stamp)
	}
	return hash
}
