package main

import (
	"fmt"
	"strings"
)

func quest02() {
	// sample := []string{"WORDS:THE,OWE,MES,ROD,HER", "", "AWAKEN THE POWER ADORNED WITH THE FLAMES BRIGHT IRE"}
	input := ReadLines("input/q02_p1.txt")

	words := make(map[string]int)
	wordList := strings.Split(input[0], ":")[1]
	max := 0
	for _, word := range strings.Split(wordList, ",") {
		words[word] = 0
		if len(word) > max {
			max = len(word)
		}
	}

	count := 0
	for i := 0; i < len(input[2]); i++ {
		for l := 1; l <= max && i+l < len(input[2]); l++ {
			word := input[2][i : i+l]
			if _, ok := words[word]; ok {
				count++
			}
		}
	}

	fmt.Println("Quest 02 Part 1:", count)

	// sample := []string{"WORDS:THE,OWE,MES,ROD,HER", "",
	// 	"AWAKEN THE POWER ADORNED WITH THE FLAMES BRIGHT IRE",
	// 	"THE FLAME SHIELDED THE HEART OF THE KINGS",
	// 	"POWE PO WER P OWE R",
	// 	"THERE IS THE END"}
	input = ReadLines("input/q02_p2.txt")
	words = make(map[string]int)
	wordList = strings.Split(input[0], ":")[1]
	max = 0
	for _, word := range strings.Split(wordList, ",") {
		words[word] = 0
		words[Reverse(word)] = 0
		if len(word) > max {
			max = len(word)
		}
	}

	symbols := make(map[symbol]int)
	for lineNum := 2; lineNum < len(input); lineNum++ {
		line := input[lineNum]
		for word := range words {
			col := 0
			for index := strings.Index(line[col:], word); index != -1 && col+index < len(line); index = strings.Index(line[col:], word) {
				for i := index + col; i < index+col+len(word); i++ {
					symbols[symbol{lineNum, i}]++
				}
				col += index + 1
			}
		}
	}

	count = len(symbols)

	fmt.Println("Quest 02 Part 2:", count)

	// sample := []string{"WORDS:THE,OWE,MES,ROD,RODEO", "",
	// 	"HELWORLT",
	// 	"ENIGWDXL",
	// 	"TRODEOAL"}
	input = ReadLines("input/q02_p3.txt")
	words = make(map[string]int)
	wordList = strings.Split(input[0], ":")[1]
	max = 0
	for _, word := range strings.Split(wordList, ",") {
		words[word] = 0
		words[Reverse(word)] = 0
		if len(word) > max {
			max = len(word)
		}
	}

	symbols = make(map[symbol]int)
	grid := input[2:]

	for word := range words {
		for row := 0; row < len(grid); row++ {
			for col := 0; col < len(grid[0]); col++ {
				for char := 0; char < len(word); char++ {
					if grid[row][(col+char)%len(grid[0])] != word[char] {
						break
					}
					if char == len(word)-1 {
						for i := 0; i < len(word); i++ {
							symbols[symbol{row, (col + i) % len(grid[0])}]++
						}
					}
				}
				for char := 0; char < len(word) && row+char < len(grid); char++ {
					if grid[row+char][col] != word[char] {
						break
					}
					if char == len(word)-1 {
						for i := 0; i < len(word); i++ {
							symbols[symbol{row + i, col}]++
						}
					}
				}
			}
		}
	}

	count = len(symbols)

	fmt.Println("Quest 02 Part 3:", count)
}

type symbol struct {
	line, col int
}
