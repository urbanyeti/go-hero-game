package main

import (
	"github.com/urbanyeti/go-hero-game/game"
)

const loopTurns int = 10
const maxTurns int = 100
const messageDelay = 0
const turnDelay = 200
const loopDelay = 1000

func main() {
	game := game.Game{}
	game.Initialize()

	for i := 0; i < maxTurns; i++ {
		gameOver := game.PlayTurn()
		if gameOver {
			break
		}
	}
}
