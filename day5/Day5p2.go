package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	maxX := 0
	maxY := 0

	lines := [][]int{}
	for scanner.Scan() {
		s := scanner.Text()
		line := readInputLine(s)
		maxX = max(maxX, max(line[0], line[2]))
		maxY = max(maxY, max(line[1], line[3]))

		x1 := line[0]
		x2 := line[2]
		y1 := line[1]
		y2 := line[3]

		// use only horizontal or diagonal
		if isHorizontal(x1, x2, y1, y2) || isVertical(x1, x2, y1, y2) || isDiagonal(x1, x2, y1, y2) {
			lines = append(lines, line)
		}
	}

	grid := make([][]int, maxY+1)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, maxX+1)
	}

	// fill the grid
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		x1 := line[0]
		x2 := line[2]
		y1 := line[1]
		y2 := line[3]

		if isVertical(x1, x2, y1, y2) {
			// from min y to max y
			// x is the same
			for y := min(y1, y2); y <= max(y1, y2); y++ {
				grid[y][x1]++
			}
		} else if isHorizontal(x1, x2, y1, y2) {
			// from min x to max x
			// y is the same
			for x := min(x1, x2); x <= max(x1, x2); x++ {
				grid[y1][x]++
			}
		} else {
			/*
					1		down - 42 64 -> y grows, x grows
				     1
				      1

				      1     up - 13-31   -> y decreases, x grows
				     1
				   1

				find most left element
			*/

			var leftX int
			var leftY int
			//var rightX int
			var rightY int
			if x1 < x2 {
				leftX = x1
				leftY = y1
				//rightX = x2
				rightY = y2
			} else {
				leftX = x2
				leftY = y2
				//rightX = x1
				rightY = y1
			}

			var isGrowing = leftY > rightY
			diff := int(math.Abs(float64(x1) - float64(x2)))

			for n := 0; n <= diff; n++ {
				if isGrowing {
					grid[leftY-n][leftX+n]++
				} else {
					grid[leftY+n][leftX+n]++
				}
			}
		}
	}

	// count result
	result := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] > 1 {
				result++
			}
		}
	}
	println(result)
}

/*
	returns [x1,y1,x2,y2]
*/
func readInputLine(s string) []int {
	result := make([]int, 4)
	fields := strings.Fields(s)
	start := strings.Split(fields[0], ",")
	dest := strings.Split(fields[2], ",")
	result[0], _ = strconv.Atoi(start[0])
	result[1], _ = strconv.Atoi(start[1])
	result[2], _ = strconv.Atoi(dest[0])
	result[3], _ = strconv.Atoi(dest[1])
	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func isHorizontal(x1, x2, y1, y2 int) bool {
	return y1 == y2
}

func isVertical(x1, x2, y1, y2 int) bool {
	return x1 == x2
}

func isDiagonal(x1, x2, y1, y2 int) bool {
	maxX := max(x1, x2)
	minX := min(x1, x2)
	maxY := max(y1, y2)
	minY := min(y1, y2)
	return (maxX - minX) == (maxY - minY)
}
