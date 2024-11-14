package main

import (
	"fmt"
	"math"
	"strconv"
)

func quest08() {
	// sample := "13"
	input := ReadLines("input/q08_p1.txt")[0]

	available, _ := strconv.Atoi(input)
	height := int(math.Sqrt(float64(available)))
	width := height*2 + 1
	additional := (height * height) + 2*height + 1 - available
	result := width * additional

	fmt.Println("Quest 08 Part 1:", result)

	// sample := "3"
	input = ReadLines("input/q08_p2.txt")[0]

	priests, _ := strconv.Atoi(input)
	acolytes := 1111
	available = 20240000

	thickness := 1
	width = 1
	for available > 0 {
		available -= width * thickness
		thickness = thickness * priests % acolytes
		width += 2
	}

	result = (width - 2) * (-available)

	fmt.Println("Quest 08 Part 2:", result)

	// sample := "2"
	// input = sample//ReadLines("input/q08_p3.txt")[0]
	// priests, _ = strconv.Atoi(input)
	// acolytes = 5//10
	// available = 160//202400000

	// thickness = 1
	// width = 1
	// for available > 0 {
	// 	available -= width * thickness
	// 	available += priests * width * thickness % acolytes
	// 	thickness = thickness * priests % acolytes + acolytes
	// 	width += 2
	// }

	// result = -available
	// fmt.Println("Quest 08 Part 3:", result)
}
