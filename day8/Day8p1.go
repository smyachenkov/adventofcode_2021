package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	result := 0
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "|")
		rightFields := strings.Fields(s[1])
		for _, v := range rightFields {
			if len(v) == 2 || len(v) == 3 || len(v) == 4 || len(v) == 7 {
				result++
			}
		}
	}
	println(result)

}
