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

	values := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if binStringLen == 0 {
			binStringLen = len(line)
		}
		values = append(values, line)
	}

	// most common
	mostCommonFiltered := make([]string, len(values))
	copy(mostCommonFiltered, values)

	idx := 0
	for len(mostCommonFiltered) != 1 {
		currentFrequent := getMostFrequentBit(mostCommonFiltered, idx)
		n := 0
		for _, line := range mostCommonFiltered {
			if line[idx] == currentFrequent {
				mostCommonFiltered[n] = line
				n++
			}
		}
		mostCommonFiltered = mostCommonFiltered[:n]
		idx++
	}

	// least common
	leastCommonFiltered := make([]string, len(values))
	copy(leastCommonFiltered, values)

	idx = 0
	for len(leastCommonFiltered) != 1 {
		currentRare := getLeastFrequentBit(leastCommonFiltered, idx)
		n := 0
		for _, line := range leastCommonFiltered {
			if line[idx] == currentRare {
				leastCommonFiltered[n] = line
				n++
			}
		}
		leastCommonFiltered = leastCommonFiltered[:n]
		idx++
	}

	oxygenGeneratorRating, _ := strconv.ParseInt(mostCommonFiltered[0], 2, 64)
	scrubberRating, _ := strconv.ParseInt(leastCommonFiltered[0], 2, 64)

	println(oxygenGeneratorRating)
	println(scrubberRating)
	println(oxygenGeneratorRating * scrubberRating)
}

func getMostFrequentBit(values []string, idx int) uint8 {
	balance := 0
	for _, v := range values {
		if v[idx] == '0' {
			balance--
		} else {
			balance++
		}
	}
	if balance < 0 {
		return '0'
	} else {
		return '1'
	}
}

func getLeastFrequentBit(values []string, idx int) uint8 {
	mostFrequent := getMostFrequentBit(values, idx)
	if mostFrequent == '1' {
		return '0'
	} else {
		return '1'
	}
}
