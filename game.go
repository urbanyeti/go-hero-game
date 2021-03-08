package main

import (
	"fmt"
	"math/rand"

	"github.com/urbanyeti/go-hero-game/hero"
)

type game struct {
	hero *hero.Hero
	loop int
	turn int
}

func (game game) String() string {
	return fmt.Sprintf("Loop: %v Turn: %v | %v", game.loop, game.turn, game.hero)
}

func (game *game) Init() {
	game.loop = 1
	game.turn = 1
	game.hero = &hero.Hero{ID: "hero-dan", Name: "Dan", Description: "a cool hero"}
	game.hero.SetDefaultStats()
	game.hero.SetDefaultEquipment()
}

func (game *game) NextTurn() {
	if game.turn < maxTurns {
		game.turn++
	} else {
		game.loop++
		game.turn = 1
	}

	event := rand.Intn(6)
	switch event {
	case 0:
		game.hero.AddStat("atk", 2)
	case 1:
		game.hero.AddStat("def", 2)
	case 2:
		game.hero.HP.Value += 2
	case 3:
		game.hero.AddStat("atk", -2)
	case 4:
		game.hero.AddStat("def", -2)
	case 5:
		game.hero.HP.Value -= 2
	}
	fmt.Println(game)
}
