package main

import (
	"bufio"
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

	sum := 0
	for _, crabPos := range state {
		sum += crabPos
	}
	mean := sum / len(state)
	result := 0
	for _, crabPos := range state {
		result += fuelToTravel(abs(crabPos - mean))
	}
	println(result)
}

func fuelToTravel(distance int) int {
	return int(0.5 * float64(distance) * float64(1+distance))
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
