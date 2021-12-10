package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	score := map[int32]int{}
	score[')'] = 3
	score[']'] = 57
	score['}'] = 1197
	score['>'] = 25137

	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		x := indexOfFirstIllegalChar(line)
		if x > 0 {
			result += score[x]
		}
	}
	println(result)
}

func indexOfFirstIllegalChar(s string) int32 {
	var stack []int32
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' || v == '<' {
			stack = append(stack, v)
		}
		if v == ')' || v == '}' || v == ']' || v == '>' {
			if len(stack) == 0 {
				return v
			}
		}
		switch v {
		case ')':
			if stack[len(stack)-1] == '(' {
				stack = stack[:(len(stack) - 1)]
			} else {
				return v
			}
		case '}':
			if stack[len(stack)-1] == '{' {
				stack = stack[:(len(stack) - 1)]
			} else {
				return v
			}
		case ']':
			if stack[len(stack)-1] == '[' {
				stack = stack[:(len(stack) - 1)]
			} else {
				return v
			}
		case '>':
			if stack[len(stack)-1] == '<' {
				stack = stack[:(len(stack) - 1)]
			} else {
				return v
			}
		}
	}
	return -1
}
