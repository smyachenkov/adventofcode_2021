package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var binaryInput []int
	for scanner.Scan() {
		line := scanner.Text()
		binaryInput = hexToBin(line)
	}

	// version
	idx := 0
	var version int
	var packetType int

	result := 0
	for idx < len(binaryInput) {
		versionBin := binaryInput[idx : idx+3]
		version = int(parseVersion(versionBin))
		result += version

		packetTypeBin := binaryInput[idx+3 : idx+6]
		packetType = int(parsePacketType(packetTypeBin))

		idx += 6
		if idx >= len(binaryInput) {
			break
		}
		// value
		if packetType == 4 {
			isLastGroup := false
			var packetResult []int
			for !isLastGroup {
				group := binaryInput[idx : idx+5]
				isLastGroup = parseIsLastGroup(group)
				packetResult = append(packetResult, group[1:]...)
				idx += 5
			}
			// operator
		} else {
			/*
			 0, then the next 15 bits are a number that represents the total length
			 1, then the next 11 bits are a number that represents the number of sub-packets contained
			*/
			isTotalLen := binaryInput[idx] == 0
			idx++
			if isTotalLen {
				//subpacketsTotalLength := parseGroup(binaryInput[idx : idx+15])
				idx += 15
			} else {
				//numberOfSubpackets := parseGroup(binaryInput[idx : idx+11])
				idx += 11
			}
		}
	}
	fmt.Printf("Result: %d", result)
}

func parseVersion(s []int) int64 {
	str := ""
	for _, z := range s {
		if z == 0 {
			str += "0"
		} else {
			str += "1"
		}
	}
	n, _ := strconv.ParseInt(str, 2, 64)
	return n
}

func parsePacketType(s []int) int64 {
	str := ""
	for _, z := range s {
		if z == 0 {
			str += "0"
		} else {
			str += "1"
		}
	}
	n, _ := strconv.ParseInt(str, 2, 64)
	return n
}

func parseIsLastGroup(s []int) bool {
	return s[0] == 0
}

var hexChars = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

func hexToBin(hex string) []int {
	var result []int
	for _, c := range hex {
		z := hexChars[string(c)]
		for _, zc := range z {
			var v int
			if string(zc) == "0" {
				v = 0
			} else {
				v = 1
			}
			result = append(result, v)
		}
	}
	return result
}
