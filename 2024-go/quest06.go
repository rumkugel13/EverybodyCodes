package main

import (
	"fmt"
	"strings"
)

func quest06() {
	fmt.Println("Quest 06 Part 1:", 0)
	// sample := []string{
	// 	"RR:A,B,C",
	// 	"A:D,E",
	// 	"B:F,@",
	// 	"C:G,H",
	// 	"D:@",
	// 	"E:@",
	// 	"F:@",
	// 	"G:@",
	// 	"H:@",
	// }

	input := ReadLines("input/q06_p1.txt")
	graph := q6_graph(input)
	paths := q6_paths(graph, "RR")
	counts := q6_counts(paths)

	result := ""
	for i := range counts {
		if len(counts[i]) == 1 {
			result = strings.ReplaceAll(counts[i][0], ",", "")
			break
		}
	}

	fmt.Println("Quest 06 Part 1:", result)

	input = ReadLines("input/q06_p2.txt")
	graph = q6_graph(input)
	paths = q6_paths(graph, "RR")
	counts = q6_counts(paths)

	result = ""
	for i := range counts {
		if len(counts[i]) == 1 {
			split := strings.Split(counts[i][0], ",")
			for _, s := range split {
				result += s[0:1]
			}
			break
		}
	}

	fmt.Println("Quest 06 Part 2:", result)

	input = ReadLines("input/q06_p3.txt")
	graph = q6_graph(input)
	paths = q6_paths(graph, "RR")
	counts = q6_counts(paths)

	result = ""
	for i := range counts {
		if len(counts[i]) == 1 {
			split := strings.Split(counts[i][0], ",")
			for _, s := range split {
				result += s[0:1]
			}
			break
		}
	}

	fmt.Println("Quest 06 Part 3:", result)
}

func q6_counts(paths []string) map[int][]string {
	counts := map[int][]string{}
	for _, path := range paths {
		count := len(path)
		counts[count] = append(counts[count], path)
	}
	return counts
}

func q6_paths(graph map[string][]string, s1 string) []string {
	var paths []string
	if s1 == "@" {
		return paths
	}

	type Node struct {
		value string
		path  string
	}

	visited := make(map[string]bool)
	queue := []Node{{s1, s1}}
	visited[s1] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, next := range graph[current.value] {
			if next == "@" {
				paths = append(paths, current.path+",@")
			} else if !visited[next] {
				visited[next] = true
				queue = append(queue, Node{next, current.path + "," + next})
			}
		}
	}

	return paths
}

func q6_graph(input []string) map[string][]string {
	graph := map[string][]string{}
	for _, line := range input {
		split := strings.Split(line, ":")
		graph[split[0]] = strings.Split(split[1], ",")
	}
	return graph
}
