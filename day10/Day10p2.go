package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var repairScores []int
	for scanner.Scan() {
		line := scanner.Text()
		score, err := repairScore(line)
		if err == nil {
			repairScores = append(repairScores, score)
		}
	}
	sort.Slice(repairScores, func(a, b int) bool {
		return repairScores[a] < repairScores[b]
	})
	println(repairScores[len(repairScores)/2])
}

func repairScore(s string) (int, error) {
	var stack []int32
	for i, v := range s {
		if v == '(' || v == '{' || v == '[' || v == '<' {
			stack = append(stack, v)
		}
		if v == ')' || v == '}' || v == ']' || v == '>' {
			if len(stack) == 0 {
				return -1, errors.New(fmt.Sprintf("Invalid char at %d", i))
			}
		}
		switch v {
		case ')':
			if stack[len(stack)-1] == '(' {
				stack = stack[:(len(stack) - 1)]
			} else {
				return -1, errors.New(fmt.Sprintf("Invalid char at %d", i))
			}
		case '}':
			if stack[len(stack)-1] == '{' {
				stack = stack[:(len(stack) - 1)]
			} else {
				return -1, errors.New(fmt.Sprintf("Invalid char at %d", i))
			}
		case ']':
			if stack[len(stack)-1] == '[' {
				stack = stack[:(len(stack) - 1)]
			} else {
				return -1, errors.New(fmt.Sprintf("Invalid char at %d", i))
			}
		case '>':
			if stack[len(stack)-1] == '<' {
				stack = stack[:(len(stack) - 1)]
			} else {
				return -1, errors.New(fmt.Sprintf("Invalid char at %d", i))
			}
		}
	}

	score := 0
	for i := len(stack) - 1; i >= 0; i-- {
		score *= 5
		switch stack[i] {
		case '(':
			print(")")
			score += 1
		case '[':
			print("]")
			score += 2
		case '{':
			print("}")
			score += 3
		case '<':
			print(">")
			score += 4
		}
	}
	print(" ")
	print(score)
	println()
	return score, nil
}
