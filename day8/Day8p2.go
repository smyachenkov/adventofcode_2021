package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	result := 0
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "|")
		rightFields := strings.Fields(s[1])
		leftFields := strings.Fields(s[0])
		code := resolveCode(leftFields)
		lineSum := 0
		n := 1000
		for _, v := range rightFields {
			lineSum += code[sortString(v)] * n
			n = n / 10
		}
		result += lineSum
	}
	println(result)
}

func resolveCode(s []string) map[string]int {
	realToBroken := map[string]string{}

	originFrequency := map[int32]int{}
	originFrequency['a'] = 8
	originFrequency['c'] = 8
	originFrequency['g'] = 7
	originFrequency['d'] = 7
	originFrequency['b'] = 6
	originFrequency['e'] = 4
	originFrequency['f'] = 9

	currentFrequency := map[int32]int{}
	for _, word := range s {
		for _, char := range word {
			currentFrequency[char]++
		}
	}

	// unique frequencies
	realToBroken["b"] = string(findKeysWithSameVal(originFrequency, currentFrequency, 6)[1])
	realToBroken["e"] = string(findKeysWithSameVal(originFrequency, currentFrequency, 4)[1])
	realToBroken["f"] = string(findKeysWithSameVal(originFrequency, currentFrequency, 9)[1])
	// bef

	// 7 minus 1 = c to broken
	var digit1 string
	var digit7 string
	for _, v := range s {
		if len(v) == 2 {
			digit1 = v
		}
		if len(v) == 3 {
			digit7 = v
		}
	}
	realToBroken["a"] = string(findUniqueChar(digit1, digit7))
	currentFrequency[findUniqueChar(digit1, digit7)] = 0
	// abef

	// c- one with freq 8 but not a
	realToBroken["c"] = string(findKeyByVal(currentFrequency, 8))
	// abcef

	var digit4 string
	for _, v := range s {
		if len(v) == 4 {
			digit4 = v
		}
	}
	digit4Extra := ""
	digit4Extra += realToBroken["b"]
	digit4Extra += realToBroken["c"]
	digit4Extra += realToBroken["f"]
	realToBroken["d"] = string(findUniqueChar(digit4Extra, digit4))
	currentFrequency[findUniqueChar(digit4Extra, digit4)] = 0
	// abcdef

	realToBroken["g"] = string(findKeyByVal(currentFrequency, 7))
	//abcdefg

	brokenDigitToReal := map[string]int{}

	// 1
	broken0 := realToBroken["a"] + realToBroken["b"] + realToBroken["c"] + realToBroken["e"] + realToBroken["f"] + realToBroken["g"]
	broken1 := realToBroken["c"] + realToBroken["f"]
	broken2 := realToBroken["a"] + realToBroken["c"] + realToBroken["d"] + realToBroken["e"] + realToBroken["g"]
	broken3 := realToBroken["a"] + realToBroken["c"] + realToBroken["d"] + realToBroken["f"] + realToBroken["g"]
	broken4 := realToBroken["b"] + realToBroken["c"] + realToBroken["d"] + realToBroken["f"]
	broken5 := realToBroken["a"] + realToBroken["b"] + realToBroken["d"] + realToBroken["f"] + realToBroken["g"]
	broken6 := realToBroken["a"] + realToBroken["b"] + realToBroken["d"] + realToBroken["e"] + realToBroken["f"] + realToBroken["g"]
	broken7 := realToBroken["a"] + realToBroken["c"] + realToBroken["f"]
	broken8 := realToBroken["a"] + realToBroken["b"] + realToBroken["c"] + realToBroken["d"] + realToBroken["e"] + realToBroken["f"] + realToBroken["g"]
	broken9 := realToBroken["a"] + realToBroken["b"] + realToBroken["c"] + realToBroken["d"] + realToBroken["f"] + realToBroken["g"]

	brokenDigitToReal[sortString(broken0)] = 0
	brokenDigitToReal[sortString(broken1)] = 1
	brokenDigitToReal[sortString(broken2)] = 2
	brokenDigitToReal[sortString(broken3)] = 3
	brokenDigitToReal[sortString(broken4)] = 4
	brokenDigitToReal[sortString(broken5)] = 5
	brokenDigitToReal[sortString(broken6)] = 6
	brokenDigitToReal[sortString(broken7)] = 7
	brokenDigitToReal[sortString(broken8)] = 8
	brokenDigitToReal[sortString(broken9)] = 9
	return brokenDigitToReal
}

// origin to broken
func findKeysWithSameVal(a, b map[int32]int, val int) []int32 {
	for aKey, aVal := range a {
		for bKey, bVal := range b {
			if aVal == val && aVal == bVal {
				return []int32{aKey, bKey}
			}
		}
	}
	return []int32{}
}

func findKeyByVal(m map[int32]int, val int) int32 {
	for k, v := range m {
		if v == val {
			return k
		}
	}
	return -1
}

func findUniqueChar(a, b string) int32 {
	m := map[int32]int{}
	for _, charA := range a {
		m[charA]++
	}
	for _, charB := range b {
		m[charB]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return -1
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
