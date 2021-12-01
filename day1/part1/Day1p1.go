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
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			result++
		}
	}

	print(result)
}
