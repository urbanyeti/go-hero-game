package main

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/urbanyeti/go-hero-game/character"
)

// Game contains state and data about the game session
type Game struct {
	Hero       *character.Hero
	Loop       int
	Turn       int
	Encounters []Encounter
	Monsters   character.LoadedMonsters
	Items      character.LoadedItems
}

func (game Game) String() string {
	return fmt.Sprintf("Loop: %v Turn: %v | %v", game.Loop, game.Turn, game.Hero)
}

// LoadData loads resources for the Game
func (game *Game) LoadData() {
	game.Monsters = character.LoadMonsters()
	game.Items = character.LoadItems()
}

// Init sets up the Game
func (game *Game) Init() {
	game.Loop = 1
	game.Turn = 1
	game.Hero = character.NewHero("hero-dan", "Dan", "cool hero")
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
func (game Game) SetDefaultEquipment(hero *character.Hero) {
	hero.Equip(game.Items["item-sword1"])
	hero.Equip(game.Items["item-boots1"])
}

// PlayTurn plays out the next Game turn
func (game *Game) PlayTurn() bool {
	//fmt.Println(game)
	log.WithFields(log.Fields{"game": game}).Info("turn started")
	gameOver := game.Encounters[rand.Intn(len(game.Encounters))].Start(game)
	if gameOver {
		return true
	}
	//fmt.Println()
	time.Sleep(turnDelay * time.Millisecond)
	if game.Turn < loopTurns {
		game.Turn++
	} else {
		log.WithFields(log.Fields{"game": game}).Info("new loop started")
		//fmt.Print("New Loop! Resting...\n\n")
		game.Loop++
		game.Turn = 1
		game.Hero.SetHP(game.Hero.Stat("hp-max"))
		time.Sleep(loopDelay * time.Millisecond)
	}
	return false
}
