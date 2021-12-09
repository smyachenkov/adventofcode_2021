package main

import (
	"bufio"
	"fmt"
	"os"
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
	result := 0
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
				fmt.Println(grid[row][col])
				result += 1 + grid[row][col]
			}
		}
	}
	println(result)

}
