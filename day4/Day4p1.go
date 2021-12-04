package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const boardSize = 5

type GameState struct {
	boards       [][][]int
	boardNumbers map[int]map[int][]int // number to board id to [row col]
	boardResult  [][]int               // matrix id to [] with boardSize * 2 len.
}

/*        * [01] -> inc (0 + row) inc (3 + col)
		[101] 2
		[010] 2
		[111] 3
         222

		[223222]
		boardResult = {0, {223222}}
*/

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	// numbers
	scanner.Scan()
	numbers := parseNumbersLine(scanner.Text())

	// boards
	boards := parseBoards(*scanner)

	// init numbers map
	numbersMap := map[int]map[int][]int{}
	for i, board := range boards {
		for row := 0; row < boardSize; row++ {
			for col := 0; col < boardSize; col++ {
				n := board[row][col]
				if _, ok := numbersMap[n]; !ok {
					numbersMap[n] = map[int][]int{}
				}
				numbersMap[n][i] = []int{row, col}
			}
		}
	}

	boardResult := make([][]int, len(boards))
	for i := 0; i < len(boards); i++ {
		boardResult[i] = make([]int, 2*boardSize)
	}

	gameState := GameState{
		boards:       boards,
		boardNumbers: numbersMap,
		boardResult:  boardResult,
	}

	// process input
	calledNumbers := []int{}
	winner := -1
	for _, n := range numbers {
		calledNumbers = append(calledNumbers, n)
		affectedBoards := gameState.boardNumbers[n]
		for boardId, rowcol := range affectedBoards {
			row := rowcol[0]
			col := rowcol[1]
			boardRes := gameState.boardResult[boardId]
			boardRes[row]++
			boardRes[boardSize+col]++
			if boardRes[row] == boardSize || boardRes[boardSize+col] == boardSize {
				winner = boardId
				break
			}
		}
		if winner != -1 {
			break
		}
	}
	println("Winner board id: " + strconv.Itoa(winner))

	score := calculateScore(gameState.boards[winner], calledNumbers)
	println(score)
}

func parseNumbersLine(s string) []int {
	numbersStrings := strings.Split(s, ",")
	result := make([]int, len(numbersStrings))
	for i := range numbersStrings {
		n, _ := strconv.Atoi(numbersStrings[i])
		result[i] = n
	}
	return result
}

func parseBoards(s bufio.Scanner) [][][]int {
	result := [][][]int{}
	for s.Scan() {
		board := [][]int{}
		for i := 0; i < boardSize; i++ {
			s.Scan()
			line := parseBoardLine(s.Text())
			board = append(board, line)
		}
		result = append(result, board)

	}
	return result
}

func parseBoardLine(s string) []int {
	numbersStrings := strings.Fields(s)
	result := make([]int, len(numbersStrings))
	for i := range numbersStrings {
		n, _ := strconv.Atoi(numbersStrings[i])
		result[i] = n
	}
	return result
}

func calculateScore(board [][]int, calledNumbers []int) int {
	winningNumber := calledNumbers[len(calledNumbers)-1]
	unmarkedNumbers := map[int]bool{}

	// at start all numbers are unmarked
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			n := board[row][col]
			unmarkedNumbers[n] = true
		}
	}

	// exclude all called numbers
	for _, v := range calledNumbers {
		unmarkedNumbers[v] = false
	}

	unmkarkedSum := 0
	for k, v := range unmarkedNumbers {
		if v {
			unmkarkedSum += k
		}
	}
	return unmkarkedSum * winningNumber
}
