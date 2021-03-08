package main

import (
	"fmt"
	"math/rand"

	"github.com/urbanyeti/go-hero-game/characters"
)

type game struct {
	hero       *characters.Hero
	loop       int
	turn       int
	encounters []Encounter
}

func (game game) String() string {
	return fmt.Sprintf("Loop: %v Turn: %v | %v", game.loop, game.turn, game.hero)
}

func (game *game) Init() {
	game.loop = 1
	game.turn = 1
	game.hero = &characters.Hero{ID: "hero-dan", Name: "Dan", Description: "a cool hero"}
	game.hero.SetDefaultStats()
	game.hero.SetDefaultEquipment()
	game.encounters = []Encounter{
		CombatEncounter{},
		CombatEncounter{},
		CutsceneEncounter{"A sppooky encounter..."},
		CutsceneEncounter{"A magical gift!"},
		CutsceneEncounter{"A random act of chaos"},
	}

}

func (game *game) NextTurn() {
	if game.turn < maxTurns {
		game.turn++
	} else {
		game.loop++
		game.turn = 1
	}

	event := rand.Intn(5)
	game.encounters[event].Start(game.hero)
	fmt.Println()
	fmt.Println(game)
}
