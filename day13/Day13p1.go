package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	coords := [][]int{}
	foldInstructions := [][]string{}
	isInstructionsPart := false

	gridRows := 0
	gridCols := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isInstructionsPart = true
			continue
		}
		// parse grid or fold instructions
		if isInstructionsPart {
			c := strings.Replace(line, "fold along ", "", -1)
			s := strings.Split(c, "=")
			foldInstructions = append(foldInstructions, []string{s[0], s[1]})
		} else {
			c := strings.Split(line, ",")
			col, _ := strconv.Atoi(c[0])
			row, _ := strconv.Atoi(c[1])
			coords = append(coords, []int{col, row})
			gridCols = max(gridCols, col+1)
			gridRows = max(gridRows, row+1)
		}
	}

	// build grid
	grid := make([][]bool, gridRows)
	for row := 0; row < gridRows; row++ {
		grid[row] = make([]bool, gridCols)
	}
	for _, v := range coords {
		row := v[1]
		col := v[0]
		grid[row][col] = true
	}

	// process fold
	for i := 0; i < 1; i++ {
		foldInstruction := foldInstructions[i]
		if foldInstruction[0] == "x" {
			grid = foldVertical(grid)
		} else {
			grid = foldHorizontal(grid)
		}
	}

	// count result
	result := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] {
				result++
			}
		}
	}
	fmt.Printf("Result %d", result)
}

func foldHorizontal(grid [][]bool) [][]bool {
	result := make([][]bool, len(grid)/2)
	// first n/2 rows
	for row := 0; row < len(result); row++ {
		result[row] = make([]bool, len(grid[0]))
		copy(result[row], grid[row])
		mirrorRow := len(grid) - 1 - row
		for mCol, v := range grid[mirrorRow] {
			result[row][mCol] = v || result[row][mCol]
		}
	}
	return result
}

func foldVertical(grid [][]bool) [][]bool {
	result := make([][]bool, len(grid))
	for row := 0; row < len(result); row++ {
		result[row] = make([]bool, len(grid[0])/2)
		copy(result[row], grid[row])
		for col := 0; col < len(grid[0])/2; col++ {
			mCol := len(grid[0]) - 1 - col
			result[row][col] = result[row][col] || grid[row][mCol]
		}
	}
	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
