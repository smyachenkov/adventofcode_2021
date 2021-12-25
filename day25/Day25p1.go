package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const empty = '.'
const east = '>'
const south = 'v'

type coord struct {
	row, col int
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var grid [][]int
	for scanner.Scan() {
		charLine := strings.Split(scanner.Text(), "")
		line := []int{}
		for _, c := range charLine {
			if c == "." {
				line = append(line, empty)
			} else if c == ">" {
				line = append(line, east)
			} else if c == "v" {
				line = append(line, south)
			}
		}
		grid = append(grid, line)
	}

	turns := 1
	for processTurn(grid) {
		turns++
	}
	fmt.Printf("Result: %d\n", turns)

}

func printGrid(grid [][]int) {
	rows := len(grid)
	cols := len(grid[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			fmt.Print(string(grid[row][col]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func processTurn(grid [][]int) bool {
	rows := len(grid)
	cols := len(grid[0])
	changed := false

	newNodes := map[coord]bool{}
	wasBefore := map[coord]bool{}

	// east
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if isProcessed(row, col, newNodes) {
				continue
			}
			if grid[row][col] == east {
				nextCol := col + 1
				if nextCol == cols {
					nextCol = 0
				}
				if grid[row][nextCol] == empty && !isProcessed(row, nextCol, wasBefore) {
					grid[row][col] = empty
					grid[row][nextCol] = east
					wasBefore[coord{row, col}] = true
					newNodes[coord{row, nextCol}] = true
					changed = true
				}
			}
		}
	}

	newNodes = map[coord]bool{}
	wasBefore = map[coord]bool{}

	// south
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == south {
				if isProcessed(row, col, newNodes) {
					continue
				}
				prevRow := row - 1
				if prevRow == -1 {
					prevRow = cols - 1
				}
				nextRow := row + 1
				if nextRow == rows {
					nextRow = 0
				}
				if grid[nextRow][col] == empty && !isProcessed(nextRow, col, wasBefore) {
					grid[row][col] = empty
					grid[nextRow][col] = south
					wasBefore[coord{row, col}] = true
					newNodes[coord{nextRow, col}] = true
				}
			}
		}
	}
	return changed
}

func isProcessed(row, col int, m map[coord]bool) bool {
	c := coord{
		row: row,
		col: col,
	}
	if val, ok := m[c]; !ok {
		return false
	} else {
		return val
	}
}
