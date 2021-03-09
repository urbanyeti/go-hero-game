package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/urbanyeti/go-hero-game/characters"
)

// Game contains state and data about the game session
type Game struct {
	Hero       *characters.Hero
	Loop       int
	Turn       int
	Encounters []Encounter
	Monsters   characters.Monsters
	Items      characters.Items
}

func (game Game) String() string {
	return fmt.Sprintf("Loop: %v Turn: %v | %v", game.Loop, game.Turn, game.Hero)
}

// LoadData loads resources for the Game
func (game *Game) LoadData() {
	game.Monsters = characters.LoadMonsters()
	game.Items = characters.LoadItems()
}

// Init sets up the Game
func (game *Game) Init() {
	game.Loop = 1
	game.Turn = 1
	game.Hero = &characters.Hero{ID: "Hero-dan", Name: "Dan", Description: "a cool Hero"}
	game.Hero.SetDefaultStats()
	game.SetDefaultEquipment(game.Hero)

	game.Encounters = []Encounter{
		CutsceneEncounter{"A sppooky encounter..."},
		CutsceneEncounter{"A magical gift!"},
		CutsceneEncounter{"A random act of chaos"},
	}

	for _, monster := range game.Monsters {
		game.Encounters = append(game.Encounters, CombatEncounter{monster})
	}
}

// SetDefaultEquipment initializes the default equipment for the hero
func (game Game) SetDefaultEquipment(hero *characters.Hero) {
	hero.Equipment = make(characters.Equipment)
	hero.Equipment["arm-r"] = game.Item("item-sword1")
	hero.Equipment["boots"] = game.Item("item-boots1")
}

// Item retrieves an item
func (game Game) Item(itemID string) *characters.Item {
	if item, ok := game.Items[itemID]; ok {
		return item
	}

	log.Printf("cannot retrieve unknown item '%v'", itemID)
	return nil
}

// PlayTurn plays out the next Game turn
func (game *Game) PlayTurn() {
	fmt.Println(game)
	game.Encounters[rand.Intn(len(game.Encounters))].Start(game)
	fmt.Println()
	time.Sleep(turnDelay * time.Millisecond)
	if game.Turn < loopTurns {
		game.Turn++
	} else {
		fmt.Print("New Loop! Resting...\n\n")
		game.Loop++
		game.Turn = 1
		game.Hero.HP = game.Hero.Stat("hp-max")
		time.Sleep(loopDelay * time.Millisecond)
	}
}

func maxOf(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

func minOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}
