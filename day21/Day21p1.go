package main

import "fmt"

type state struct {
	player   int
	score    int
	position int
}

var nextDice = 1
var diceRollsCnt = 0

func main() {
	player1StartPos := 10
	player2StartPos := 9

	p1State := state{
		player:   1,
		score:    0,
		position: player1StartPos,
	}
	p2State := state{
		player:   2,
		score:    0,
		position: player2StartPos,
	}
	for p1State.score < 1000 && p2State.score < 1000 {
		roll(&p1State)
		if p1State.score < 1000 {
			roll(&p2State)
		}
	}

	fmt.Printf("P1 score: %d\n", p1State.score)
	fmt.Printf("P2 score: %d\n", p2State.score)

	result := 0
	if p1State.score >= 1000 {
		result = p2State.score
	} else {
		result = p1State.score
	}
	result *= diceRollsCnt
	fmt.Printf("Result: %d\n", result)

}

func roll(p *state) {
	toAdd := 0
	for i := 0; i < 3; i++ {
		toAdd += nextDice
		nextDice++
	}
	p.position = (p.position + toAdd) % 10
	if p.position == 0 {
		p.score += 10
	} else {
		p.score += p.position
	}
	diceRollsCnt += 3
	fmt.Printf("player %d at pos %d with score %d\n", p.player, p.position, p.score)
}
