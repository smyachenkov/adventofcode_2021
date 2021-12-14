package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	templateLine := scanner.Text()

	template := []rune{}
	for _, v := range templateLine {
		template = append(template, v)
	}
	scanner.Scan()

	insertionRules := map[string]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		rule := strings.Split(line, " -> ")
		insertionRules[rule[0]] = rune(rule[1][0])
	}

	for i := 0; i < 10; i++ {
		for c := 1; c < len(template); c++ {
			seq := string(template[c-1]) + string(template[c])
			if val, ok := insertionRules[seq]; ok {
				template = append(template[:c+1], template[c:]...)
				template[c] = val
				c++
			}
		}
	}

	frequency := map[rune]int{}
	for _, c := range template {
		if _, ok := frequency[c]; !ok {
			frequency[c] = 0
		}
		frequency[c]++
	}

	maxFrequency := 0
	minFrequency := math.MaxInt64
	for _, c := range frequency {
		maxFrequency = max(maxFrequency, c)
		minFrequency = min(minFrequency, c)
	}
	fmt.Printf("Result: %d", maxFrequency-minFrequency)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
