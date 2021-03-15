package main

const loopTurns int = 10
const maxTurns int = 100
const messageDelay = 100
const turnDelay = 200
const loopDelay = 1000

func main() {
	game := Game{}
	game.Init()

	for i := 0; i < maxTurns; i++ {
		gameOver := game.PlayTurn()
		if gameOver {
			break
		}
	}
}
