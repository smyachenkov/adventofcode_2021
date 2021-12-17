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

	scanner.Scan()
	line := scanner.Text()
	line = strings.Replace(line, "target area: ", "", 1)
	split := strings.Split(line, ", ")

	xPart := strings.Split(strings.Replace(split[0], "x=", "", 1), "..")
	yPart := strings.Split(strings.Replace(split[1], "y=", "", 1), "..")

	xFrom, _ := strconv.Atoi(xPart[0])
	xTo, _ := strconv.Atoi(xPart[1])
	yFrom, _ := strconv.Atoi(yPart[1])
	yTo, _ := strconv.Atoi(yPart[0])

	fmt.Println(xFrom, xTo, yFrom, yTo)

	score := 0
	for i := -500; i < 500; i++ {
		for j := -500; j < 500; j++ {
			if canHitTheArea(i, j, xFrom, xTo, yFrom, yTo) {
				score++
			}
		}
	}

	fmt.Println(score)
}

func canHitTheArea(xVelocity, yVelocity int, xFrom, xTo, yFrom, yTo int) bool {
	xVelocityStart := xVelocity
	yVelocityStart := yVelocity

	maxY := 0

	x := 0
	y := 0

	for true {
		if inArea(x, y, xFrom, xTo, yFrom, yTo) {
			fmt.Printf("Hit the area for %d, %d\n", xVelocityStart, yVelocityStart)
			return true
		}
		if missedArea(x, y, xFrom, xTo, yFrom, yTo) {
			fmt.Printf("Missed the area for %d, %d\n", xVelocityStart, yVelocityStart)
			break
		}
		if xVelocity == 0 && y < yTo {
			fmt.Printf("Dropped before for  %d, %d\n", xVelocityStart, yVelocityStart)
			break
		}
		x += xVelocity
		y += yVelocity

		if xVelocity != 0 {
			if xVelocity > 0 {
				xVelocity--
			} else {
				xVelocity++
			}
		}
		yVelocity--
		maxY = max(maxY, y)

	}
	return false
}

func inArea(x, y, xFrom, xTo, yFrom, yTo int) bool {
	return x >= xFrom && x <= xTo && y <= yFrom && y >= yTo
}

func missedArea(x, y, xFrom, xTo, yFrom, yTo int) bool {
	/*
		 ____
		|____|
					 x > xto
				   y <= yto
		x >= xfrom
		x <= xto
		y < yto
	*/
	return (x >= xFrom && x <= xTo && y < yTo) ||
		(x > xTo && y <= y)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
