package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	state := []int{}
	fields := strings.Split(scanner.Text(), ",")
	for i := range fields {
		n, _ := strconv.Atoi(fields[i])
		state = append(state, n)
	}

	sort.Slice(state, func(a, b int) bool {
		return state[a] < state[b]
	})

	var median int
	if len(state)%2 == 0 {
		median = int(math.Ceil(float64(state[len(state)/2-1])))
	} else {
		median = state[(len(state)-1)/2]
	}

	result := 0
	for _, crabPos := range state {
		result += int(math.Abs(float64(median - crabPos)))
	}
	println(result)
}
