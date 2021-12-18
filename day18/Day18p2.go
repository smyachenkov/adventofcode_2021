package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type node struct {
	isVal  bool
	val    int
	left   *node
	right  *node
	parent *node
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	nodes := make([]*node, 0, 1)
	for scanner.Scan() {
		n := &node{}
		line := scanner.Text()
		for _, s := range line {
			switch s {
			case ',':
				n = n.parent.right
			case '[':
				n.left = &node{parent: n}
				n.right = &node{parent: n}
				n = n.left
			case ']':
				n = n.parent
			default:
				num, _ := strconv.Atoi(string(s))
				n.val = num
				n.isVal = true
			}
		}
		nodes = append(nodes, n)
	}

	result := 0
	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(nodes); j++ {
			if i == j {
				continue
			}
			sum := getMagnitude(add(copyNode(nodes[i]), copyNode(nodes[j])))
			result = max(result, sum)
		}
	}
	fmt.Println(result)

}

func add(a, b *node) *node {
	result := &node{
		left:  a,
		right: b,
		val:   0,
		isVal: false,
	}
	result.left.parent = result
	result.right.parent = result
	changes := 0
	for {
		changes = reduceExplode(result, result, 0)
		if changes != 0 {
			continue
		}
		changes = reduceSplit(result, result)
		if changes == 0 {
			break
		}
	}
	return result
}

func reduceExplode(a, head *node, level int) int {
	if level == 5 {
		explode(a.parent)
		return 1
	}
	if a.parent == nil && level != 0 {
		return 0
	}
	if a.left == nil {
		return 0
	}
	changes := reduceExplode(a.left, head, level+1)
	if a.right == nil || changes != 0 {
		return changes
	} else {
		return reduceExplode(a.right, head, level+1)
	}
}

func reduceSplit(a, head *node) int {
	if a.left == nil && a.right == nil {
		if a.val > 9 {
			a.left = &node{
				parent: a,
				val:    int(math.Floor(float64(a.val) / 2)),
				isVal:  true,
			}
			a.right = &node{
				parent: a,
				val:    int(math.Ceil(float64(a.val) / 2)),
				isVal:  true,
			}
			a.val = 0
			return 1
		}
		return 0
	}
	if a.left == nil {
		return 0
	}
	changes := reduceSplit(a.left, head)
	if a.right == nil || changes != 0 {
		return changes
	} else {
		return reduceSplit(a.right, head)
	}
}

func explode(n *node) {
	left := findFirstLeftVal(n)
	if left != nil {
		left.val += n.left.val
	}
	right := findFirstRightVal(n)
	if right != nil {
		right.val += n.right.val
	}
	n.left = nil
	n.right = nil
	n.val = 0
}

func findFirstLeftVal(n *node) *node {
	prev := n
	for number := n.parent; number != nil; {
		if number.right == prev {
			prev = number
			number = number.left
			continue
		}
		if number.left == prev {
			if n.parent == nil {
				break
			}
			prev = number
			number = number.parent
			continue
		}
		if number.left == nil && number.right == nil {
			return number
		}
		number = number.right
	}
	return nil
}

func findFirstRightVal(n *node) *node {
	prev := n
	for number := n.parent; number != nil; {
		if number.left == prev {
			prev = number
			number = number.right
			continue
		}
		if number.right == prev {
			if n.parent == nil {
				break
			}
			prev = number
			number = number.parent
			continue
		}
		if number.left == nil && number.right == nil {
			return number
		}
		number = number.left
	}
	return nil
}

func getMagnitude(n *node) int {
	result := 0
	if n.left != nil {
		result += getMagnitude(n.left) * 3
	}
	if n.right != nil {
		result += getMagnitude(n.right) * 2
	}
	if n.left == nil && n.right == nil {
		return n.val
	}
	return result
}

func copyNode(n *node) *node {
	result := &node{}
	if n.left != nil {
		result.left = copyNode(n.left)
		result.left.parent = result
	}
	if n.right != nil {
		result.right = copyNode(n.right)
		result.right.parent = result
	}
	result.val = n.val
	result.isVal = n.isVal
	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
