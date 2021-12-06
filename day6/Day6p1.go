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
	scanner.Scan()

	state := []int{}
	fields := strings.Split(scanner.Text(), ",")
	for i := range fields {
		v, _ := strconv.Atoi(fields[i])
		state = append(state, v)
	}

	for day := 0; day < 80; day++ {
		//fmt.Println(state)
		newFishCnt := 0
		for idx, days := range state {
			if days > 0 {
				state[idx]--
			} else {
				state[idx] = 6
				newFishCnt++
			}
		}
		for i := 0; i < newFishCnt; i++ {
			state = append(state, 8)
		}
	}

	println(len(state))
}
