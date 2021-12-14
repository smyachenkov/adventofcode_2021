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

	pairs := map[string]int{}
	for i := 1; i < len(template); i++ {
		key := string(template[i-1]) + string(template[i])
		if _, ok := pairs[key]; !ok {
			pairs[key] = 0
		}
		pairs[key]++
	}

	rules := map[string]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		rule := strings.Split(line, " -> ")
		rules[rule[0]] = rune(rule[1][0])
	}

	frequency := map[rune]int{}
	for _, c := range template {
		if _, ok := frequency[c]; !ok {
			frequency[c] = 0
		}
		frequency[c]++
	}
	for i := 0; i < 40; i++ {
		newPairs := map[string]int{}
		for k, v := range pairs {
			newPairs[k] = v
		}
		for pair := range pairs {
			seq := string(pair[0]) + string(pair[1])
			if _, ok := rules[seq]; ok {
				inserted := rules[seq]
				diff := pairs[seq]
				frequency[inserted] = frequency[inserted] + diff

				leftPair := string(pair[0]) + string(inserted)
				rightPair := string(inserted) + string(pair[1])

				newPairs[seq] -= diff

				newPairs[leftPair] += diff
				newPairs[rightPair] += diff
			}
		}
		// remove zeroed pairs
		for k, v := range newPairs {
			if v == 0 {
				delete(newPairs, k)
			}
		}
		pairs = newPairs
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
