package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var depths []int
	for scanner.Scan() {
		intVal, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, intVal)
	}

	result := 0
	prevSum := 0
	currSum := 0
	for i := 0; i < 3; i++ {
		prevSum += depths[i]
	}
	prevSum = currSum

	for i := 3; i < len(depths); i++ {
		currSum = currSum - depths[i-3]
		currSum = currSum + depths[i]
		if currSum > prevSum {
			result++
		}
		prevSum = currSum
	}

	print(result)
}
