package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	binStringLen := 0
	count := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		if binStringLen == 0 {
			binStringLen = len(line)
		}
		for i := 0; i < binStringLen; i++ {
			checkBitAt(line, i, count)
		}
	}

	gammaBinaryString := ""
	for i := 0; i < binStringLen; i++ {
		gammaBinaryString += string(getMostFrequentBit(count, i))
	}

	epsilonBinaryString := ""
	for i := 0; i < binStringLen; i++ {
		epsilonBinaryString += string(getLeastFrequentBit(count, i))
	}

	gamma, _ := strconv.ParseInt(gammaBinaryString, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonBinaryString, 2, 64)

	println(gamma * epsilon)
}

func checkBitAt(s string, idx int, count map[int]int) {
	if s[idx] == '1' {
		count[idx]++
	} else {
		count[idx]--
	}
}

func getMostFrequentBit(count map[int]int, idx int) int32 {
	if count[idx] >= 0 {
		return '1'
	} else {
		return '0'
	}
}

func getLeastFrequentBit(count map[int]int, idx int) int32 {
	if count[idx] >= 0 {
		return '0'
	} else {
		return '1'
	}
}
