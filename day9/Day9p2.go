package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var directions = [][]int{
	{-1, 0}, // top
	{0, 1},  // right
	{1, 0},  // bot
	{0, -1}, // left
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	grid := [][]int{}
	lowPoints := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, v := range line {
			n, _ := strconv.Atoi(string(v))
			row = append(row, n)
		}
		grid = append(grid, row)
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			higherNeighbors := 0
			for _, d := range directions {
				neighbourRow := row + d[0]
				neighbourCol := col + d[1]
				if neighbourRow < 0 || neighbourRow == len(grid) || neighbourCol < 0 || neighbourCol == len(grid[0]) {
					higherNeighbors++
				} else {
					if grid[neighbourRow][neighbourCol] > grid[row][col] {
						higherNeighbors++
					}
				}
			}
			if higherNeighbors == 4 {
				lowPoints = append(lowPoints, []int{row, col})
			}
		}
	}

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	basins := []int{}
	for _, l := range lowPoints {
		lowPointRow := l[0]
		lowPointCol := l[1]
		visited[lowPointRow][lowPointCol] = true
		basinSize := visitBasin(grid, visited, lowPointRow, lowPointCol, 0)
		basins = append(basins, basinSize)
	}
	sort.Ints(basins)
	result := basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
	println(result)

}

func visitBasin(grid [][]int, visited [][]bool, row, col int, score int) int {
	if grid[row][col] == 9 {
		return score
	}
	score++
	visited[row][col] = true

	for _, d := range directions {
		nextRow, nextCol := row+d[0], col+d[1]
		if nextRow >= 0 && nextRow < len(grid) && nextCol >= 0 && nextCol < len(grid[0]) {
			if !visited[nextRow][nextCol] {
				score = visitBasin(grid, visited, nextRow, nextCol, score)
			}
		}
	}
	return score
}
