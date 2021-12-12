package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	graph := map[string][]string{}
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "-")
		from := row[0]
		to := row[1]
		if _, ok := graph[from]; !ok {
			graph[from] = []string{}
		}
		if _, ok := graph[to]; !ok {
			graph[to] = []string{}
		}
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	paths := [][]string{}
	path := []string{"start"}
	visit("start", path, graph, &paths, "")

	fmt.Printf("Total paths: %d", len(paths))
}

func visit(node string, path []string, graph map[string][]string, paths *[][]string, twiceVisitedSmall string) {
	if node == "end" {
		c := make([]string, len(path))
		copy(c, path)
		*paths = append(*paths, c)
		return
	}
	for _, nextNode := range graph[node] {
		if nextNode == "start" {
			continue
		}
		// ignore once visited small caves
		if isLowerCase(nextNode) && contains(path, nextNode) {
			continue
		}
		path = append(path, nextNode)
		visit(nextNode, path, graph, paths, twiceVisitedSmall)
		path = path[:len(path)-1]
	}
}

func contains(s []string, value string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

func isLowerCase(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}
