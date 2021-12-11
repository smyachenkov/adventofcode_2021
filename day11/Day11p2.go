package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var directions = [][]int{
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	grid := [][]int{}
	for scanner.Scan() {
		row := []int{}
		for _, c := range scanner.Text() {
			val, _ := strconv.Atoi(string(c))
			row = append(row, val)
		}
		grid = append(grid, row)
	}

	flashed := make([][]bool, len(grid))
	for i := range flashed {
		flashed[i] = make([]bool, len(grid[0]))
	}

	turn := 1

	for true {
		// empty flashes
		for row := 0; row < len(flashed); row++ {
			for col := 0; col < len(flashed[0]); col++ {
				flashed[row][col] = false
			}
		}

		flashedThisTurn := 0

		for row := 0; row < len(grid); row++ {
			for col := 0; col < len(grid[0]); col++ {
				if flashed[row][col] {
					continue
				}
				grid[row][col]++
				if grid[row][col] > 9 {
					// process chained flashes
					flashedThisTurn += processFlash(grid, flashed, row, col, 0)
				}
			}
		}
		fmt.Printf("Turn %d: %d flashed\n", turn, flashedThisTurn)
		if flashedThisTurn == len(grid)*len(grid[0]) {
			break
		}
		turn++
	}
}

func processFlash(grid [][]int, flashed [][]bool, row, col, result int) int {
	flashed[row][col] = true
	result++
	grid[row][col] = 0
	for _, d := range directions {
		nextRow := row + d[0]
		nextCol := col + d[1]
		if nextRow >= 0 && nextRow < len(grid) && nextCol >= 0 && nextCol < len(grid[0]) {
			if !flashed[nextRow][nextCol] {
				grid[nextRow][nextCol]++
				if grid[nextRow][nextCol] > 9 {
					result = processFlash(grid, flashed, nextRow, nextCol, result)
				}
			}
		}
	}
	return result
}
