package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	grid := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			row = append(row, n)
		}
		grid = append(grid, row)
	}

	size := len(grid)

	// min risk to reach [last][last] from it position
	riskGrid := make([][]int, size)
	for i := 0; i < size; i++ {
		riskGrid[i] = make([]int, size)
	}

	riskGrid[size-1][size-1] = grid[size-1][size-1]

	for row := size - 1; row >= 0; row-- {
		for col := size - 1; col >= 0; col-- {
			if row == size-1 && col == size-1 {
				continue
			}
			// right
			right := math.MaxInt64
			if col < size-1 {
				right = riskGrid[row][col+1]
			}
			// bot
			bot := math.MaxInt64
			if row < size-1 {
				bot = riskGrid[row+1][col]
			}
			riskGrid[row][col] = min(right, bot) + grid[row][col]
		}
	}

	result := riskGrid[0][0] - grid[0][0]

	fmt.Printf("Result: %d", result)
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
