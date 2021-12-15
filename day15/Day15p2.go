package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	grid := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			row = append(row, n)
		}
		grid = append(grid, row)
	}

	// expand grid
	originalSize := len(grid)

	// add columns
	for row := 0; row < originalSize; row++ {
		newRow := make([]int, 5*originalSize)
		copy(newRow, grid[row])
		for col := 0; col < originalSize; col++ {
			for i := 1; i < 5; i++ {
				// 0 - 50
				// 1 - 51
				nextIdx := originalSize*i + col
				nextVal := grid[row][col] + i
				if nextVal >= 10 {
					nextVal = nextVal - 9
				}
				newRow[nextIdx] = nextVal
			}
		}
		grid[row] = newRow
	}

	// add rows
	for i := 1; i < 5; i++ {
		for row := 0; row < originalSize; row++ {
			nextRow := make([]int, 5*originalSize)
			copy(nextRow, grid[row])
			grid = append(grid, nextRow)
		}
	}

	// increment values in new rows
	for i := 1; i < 5; i++ {
		for row := originalSize * i; row < originalSize*(i+1); row++ {
			for col := 0; col < len(grid[row]); col++ {
				originalRow := row - originalSize*i
				nextVal := grid[originalRow][col] + i
				if nextVal >= 10 {
					nextVal = nextVal - 9
				}
				grid[row][col] = nextVal
			}
		}
	}

	size := len(grid)
	start := [2]int{0, 0}
	finish := [2]int{size - 1, size - 1}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Node{
		point:    start,
		priority: grid[0][0],
	})

	cost := map[[2]int]int{}
	cost[start] = grid[0][0]

	path := map[[2]int][2]int{}

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Node)
		for _, neighbor := range findNeighbors(current.point, grid) {
			neighborCost := cost[current.point] + grid[neighbor[1]][neighbor[0]]
			// if it's new, or cost less
			if v, ok := cost[neighbor]; !ok || neighborCost < v {
				path[neighbor] = current.point
				cost[neighbor] = neighborCost
				heap.Push(&pq, &Node{
					point:    neighbor,
					priority: neighborCost,
				})
			}
		}
	}

	result := 0
	current := finish
	for current != start {
		result += grid[current[1]][current[0]]
		current = path[current]
	}

	fmt.Printf("Result: %d", result)
}

var directions = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func findNeighbors(current [2]int, grid [][]int) [][2]int {
	var result [][2]int
	for _, d := range directions {
		next := [2]int{current[0] + d[0], current[1] + d[1]}
		if next[0] >= 0 && next[0] < len(grid) && next[1] >= 0 && next[1] < len(grid) {
			result = append(result, next)
		}
	}
	return result
}

/*
	Priority queue implementation
*/

type Node struct {
	point    [2]int
	priority int
	index    int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
