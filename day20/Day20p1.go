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
	algoLine := scanner.Text()

	scanner.Scan()

	var grid [][]bool
	for scanner.Scan() {

		charLine := strings.Split(scanner.Text(), "")
		boolLine := []bool{}
		for _, c := range charLine {
			if c == "#" {
				boolLine = append(boolLine, true)
			} else {
				boolLine = append(boolLine, false)
			}
		}

		grid = append(grid, boolLine)
	}

	// first
	for i := 0; i < 2; i++ {
		grid = expandMatrix(grid)
		grid = enhance(grid, algoLine)
	}
	fmt.Printf("Result part 1: %d\n", countResult(grid))
	for i := 0; i < 48; i++ {
		grid = expandMatrix(grid)
		grid = enhance(grid, algoLine)
	}
	fmt.Printf("Result part 2: %d\n", countResult(grid))

}

func countResult(grid [][]bool) int {
	result := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] {
				result++
			}
		}
	}
	return result
}

func printImage(image [][]bool) {
	for row := 0; row < len(image); row++ {
		for col := 0; col < len(image[0]); col++ {
			if image[row][col] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func enhance(original [][]bool, algo string) [][]bool {
	result := [][]bool{}
	for startRow := 0; startRow < len(original)-2; startRow++ {
		resultRow := []bool{}
		for startCol := 0; startCol < len(original[0])-2; startCol++ {
			code := []string{}
			for row := 0; row < 3; row++ {
				for col := 0; col < 3; col++ {
					if original[startRow+row][startCol+col] {
						code = append(code, "1")
					} else {
						code = append(code, "0")
					}
				}
			}
			binNumber := strings.Join(code, "")
			byteToInt, _ := strconv.ParseInt(binNumber, 2, 64)
			mapped := string(algo[byteToInt])
			//fmt.Printf("row %d col %d bin %s int %s mapped %s\n", startRow+1, startCol+1, binNumber, strconv.FormatInt(byteToInt, 10), mapped)
			if mapped == "#" {
				resultRow = append(resultRow, true)
			} else {
				resultRow = append(resultRow, false)
			}
		}
		result = append(result, resultRow)
	}
	//fmt.Printf("Total new light: %d\n", cnt)
	return result
}

func expandMatrix(original [][]bool) [][]bool {
	result := make([][]bool, len(original))
	for i := range original {
		result[i] = make([]bool, len(original[i]))
		copy(result[i], original[i])
	}

	// add 2x rows and cols
	for row := 0; row < len(result); row++ {
		newLine := []bool{false, false}
		newLine = append(newLine, result[row]...)
		newLine = append(newLine, []bool{false, false}...)
		result[row] = newLine
	}
	lineSize := len(result[0])
	for row := 0; row < 2; row++ {
		result = append(result, makeEmptyLine(lineSize))
	}
	for row := 0; row < 2; row++ {
		result = append([][]bool{makeEmptyLine(lineSize)}, result...)
	}

	return result
}

func makeEmptyLine(size int) []bool {
	line := []bool{}
	for i := 0; i < size; i++ {
		line = append(line, false)
	}
	return line
}
