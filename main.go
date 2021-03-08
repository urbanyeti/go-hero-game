package main

import (
	"fmt"
	"time"
)

const maxTurns int = 10

func main() {
	game := game{}
	game.Init()

	fmt.Println(game)
	time.Sleep(1 * time.Second)
	for i := 0; i < 15; i++ {
		game.NextTurn()
		time.Sleep(1 * time.Second)
	}
}
