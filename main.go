package main

import "fmt"

const loopTurns int = 10
const maxTurns int = 100
const messageDelay = 0
const turnDelay = 200
const loopDelay = 1000

func main() {
	game := Game{}
	game.LoadData()
	game.Init()

	for i := 0; i < maxTurns; i++ {
		gameOver := game.PlayTurn()
		if gameOver {
			break
		}
	}

	fmt.Println("Finished game!")
	fmt.Println(game.Hero)
}
