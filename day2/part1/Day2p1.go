package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	length := 0
	depth := 0
	for scanner.Scan() {
		line := scanner.Text()
		entry := strings.Split(line, " ")
		command := entry[0]
		value, _ := strconv.Atoi(entry[1])
		switch command {
		case "up":
			depth -= value
		case "down":
			depth += value
		case "forward":
			length += value
		}
	}
	println(length * depth)
}
