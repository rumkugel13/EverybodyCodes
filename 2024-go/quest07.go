package main

import (
	"fmt"
	"sort"
	"strings"
)

func quest07() {
	// sample := []string{
	// 	"A:+,-,=,=",
	// 	"B:+,=,-,+",
	// 	"C:=,-,+,+",
	// 	"D:=,=,=,+",
	// }

	input := ReadLines("input/q07_p1.txt")
	plans := q7_plans(input)

	powers := map[string]int{}
	for key := range plans {
		powers[key] = 10
	}

	total := map[string]int{}
	for round := 0; round < 10; round++ {
		next := q7_next(plans, powers, round)
		for key := range plans {
			total[key] += next[key]
			powers[key] = next[key]
		}
	}

	sorted := q7_sort(total)

	var result string
	for _, kv := range sorted {
		result = kv.Key + result
	}

	fmt.Println("Quest 07 Part 1:", result)

	// sample := []string{
	// 	"A:+,-,=,=",
	// 	"B:+,=,-,+",
	// 	"C:=,-,+,+",
	// 	"D:=,=,=,+",
	// }
	input = ReadLines("input/q07_p2.txt")
	plans = q7_plans(input)
	powers = map[string]int{}
	for key := range plans {
		powers[key] = 10
	}

	// sampletrack := []string{
	// 	"S+===",
	// 	"-   +",
	// 	"=+=-+",
	// }

	track := ReadLines("input/q07_p2_track.txt")
	linetrack := q7_track(track)

	total = map[string]int{}
	for round := 0; round < 10*len(linetrack); round++ {
		for key, plan := range plans {
			powers[key] = q7_next_track(plan, powers[key], round, linetrack)
			total[key] += powers[key]
		}
	}

	sorted = q7_sort(total)

	result = ""
	for _, kv := range sorted {
		result = kv.Key + result
	}

	fmt.Println("Quest 07 Part 2:", result)

	// sample := []string{"A:+,-,=,+,-,=,+,-,=,+,+"}
	input = ReadLines("input/q07_p3.txt")
	rival_plan := q7_plans(input)["A"]

	track = ReadLines("input/q07_p3_track.txt")
	linetrack = q7_track(track)

	// every 11 loops a new cycle begins, because LCM(11,340)=3740
	// and 3740/340=11, meaning after 11 loops the start of the track and the plans align
	// since 2024 is a multiple of 11, we can just run 11 loops instead of all of them
	// and multiply the result by 2024/11=184 (or just skip it since we only check if rival is ahead, not by how much)
	rounds := 11
	rival_power := 10
	rival_total := 0
	for round := 0; round < rounds*len(linetrack); round++ {
		rival_power = q7_next_track(rival_plan, rival_power, round, linetrack)
		rival_total += rival_power
	}

	winners := 0
	action_plans := q7_action_plans()
	for _, plan := range action_plans {
		current_power := 10
		current_total := 0
		for round := 0; round < rounds*len(linetrack); round++ {
			current_power = q7_next_track(plan, current_power, round, linetrack)
			current_total += current_power
		}
		if current_total > rival_total {
			winners++
		}
	}

	fmt.Println("Quest 07 Part 3:", winners)
}

func q7_action_plans() []string {
	result := []string{}
	var generate func(current string, plus, minus, equal int)
	generate = func(current string, plus, minus, equal int) {
		if len(current) == 11 {
			result = append(result, current)
			return
		}
		if plus < 5 {
			generate(current+"+", plus+1, minus, equal)
		}
		if minus < 3 {
			generate(current+"-", plus, minus+1, equal)
		}
		if equal < 3 {
			generate(current+"=", plus, minus, equal+1)
		}
	}
	generate("", 0, 0, 0)
	return result
}

// note: +1 since track starts with S
func q7_next_track(plan string, power int, round int, linetrack string) int {
	action := plan[round%len(plan)]
	if linetrack[(round+1)%len(linetrack)] == '+' {
		action = '+'
	} else if linetrack[(round+1)%len(linetrack)] == '-' {
		action = '-'
	}
	switch action {
	case '+':
		return power + 1
	case '-':
		return power - 1
	}
	return power
}

func q7_track(track []string) string {
	result := ""
	type Position struct {
		x, y int
	}
	visited := make(map[Position]bool)
	pos := Position{0, 0}
	width := len(track[0])
	height := len(track)
outer:
	for {
		result += string(track[pos.y][pos.x])
		visited[pos] = true

		// Try all possible directions
		for _, dir := range []Position{{1, 0}, {0, 1}, {0, -1}, {-1, 0}} {
			nextPos := Position{pos.x + dir.x, pos.y + dir.y}

			if nextPos.x == 0 && nextPos.y == 0 {
				break outer
			}
			if nextPos.x >= 0 && nextPos.x < width && nextPos.y >= 0 && nextPos.y < height &&
				track[nextPos.y][nextPos.x] != ' ' && !visited[Position{nextPos.x, nextPos.y}] {
				pos.x, pos.y = nextPos.x, nextPos.y
				break
			}
		}
	}

	return result
}

func q7_plans(input []string) map[string]string {
	plans := map[string]string{}
	for _, line := range input {
		split := strings.Split(line, ":")
		plans[split[0]] = strings.ReplaceAll(split[1], ",", "")
	}
	return plans
}

func q7_next(plans map[string]string, powers map[string]int, round int) map[string]int {
	next := map[string]int{}
	for key := range plans {
		next[key] = powers[key]
		switch plans[key][round%len(plans[key])] {
		case '+':
			next[key] += 1
		case '-':
			next[key] -= 1
		case '=':
			next[key] = powers[key]
		}
	}
	return next
}

type kv struct {
	Key   string
	Value int
}

func q7_sort(total map[string]int) []kv {
	var sorted []kv
	for k, v := range total {
		sorted = append(sorted, kv{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value < sorted[j].Value
	})
	return sorted
}
