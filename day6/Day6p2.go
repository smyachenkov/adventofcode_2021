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

	// days to number of fishes
	state := map[int]int64{}
	fields := strings.Split(scanner.Text(), ",")
	for i := range fields {
		days, _ := strconv.Atoi(fields[i])
		if _, ok := state[days]; !ok {
			state[days] = 1
		} else {
			state[days]++
		}
	}

	for day := 0; day < 256; day++ {
		fmt.Println("day %s", day)
		newFish := state[0]
		for i := 1; i <= 8; i++ {
			state[i-1] = state[i]
		}
		state[6] += newFish
		state[8] = newFish
	}

	result := int64(0)
	for _, v := range state {
		result += v
	}
	println(result)
}
