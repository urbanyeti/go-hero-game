package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
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
	Monsters   map[string]character.Monster
	Abilities  map[string]character.Ability
	Items      map[string]*character.Item
}

func (g Game) String() string {
	return fmt.Sprintf("Loop: %v Turn: %v | %v", g.Loop, g.Turn, g.Hero)
}

// LoadContent loads resources used by the Game
func (g *Game) LoadContent() {
	g.LoadAbilities()
	g.LoadItems()
	g.LoadMonsters()
}

// Init sets up the Game
func (g *Game) Init() {
	g.LoadContent()
	g.Loop = 1
	g.Turn = 1
	g.Hero = character.NewHero("hero-dan", "Dan", "cool hero")
	g.SetDefaultEquipment(g.Hero)

	g.Encounters = []Encounter{
		CutsceneEncounter{"A sppooky encounter..."},
		CutsceneEncounter{"A magical gift!"},
		CutsceneEncounter{"A random act of chaos"},
	}

	for _, monster := range g.Monsters {
		g.Encounters = append(g.Encounters, CombatEncounter{[]character.Monster{monster}})
	}
}

// LoadAbilities grabs all the abilities from json
func (g *Game) LoadAbilities() {
	g.Abilities = make(map[string]character.Ability)

	jsonFile, err := os.Open("./character/json/abilities.json")
	if err != nil {
		log.Error(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var jsonVals []*character.AbilityJSON
	json.Unmarshal(byteValue, &jsonVals)
	for _, a := range jsonVals {
		g.Abilities[a.ID] = a.LoadAbility()
	}
}

// LoadItems grabs all the items from json
func (g *Game) LoadItems() {
	g.Items = make(map[string]*character.Item)

	jsonFile, err := os.Open("./character/json/items.json")
	if err != nil {
		log.Error(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var jsonVals []*character.ItemJSON
	json.Unmarshal(byteValue, &jsonVals)
	for _, i := range jsonVals {
		g.Items[i.ID] = i.LoadItem()
	}
}

// LoadMonsters grabs all the monster data from json
func (g *Game) LoadMonsters() {
	g.Monsters = make(map[string]character.Monster)

	jsonFile, err := os.Open("./character/json/monsters.json")
	if err != nil {
		log.Error(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var jsonVals []*character.CharacterJSON
	json.Unmarshal(byteValue, &jsonVals)
	for _, c := range jsonVals {
		g.Monsters[c.ID] = character.Monster(c.LoadMonster(g.Abilities, g.Items))
	}
}

// SetDefaultEquipment initializes the default equipment for the hero
func (g Game) SetDefaultEquipment(hero *character.Hero) {
	hero.Equip(g.Items["item-sword1"], g.Items["item-armor1"], g.Items["item-boots1"])
}

// PlayTurn plays out the next Game turn
func (g *Game) PlayTurn() bool {
	log.WithFields(log.Fields{"game": g}).Info("turn started")
	random := rand.Intn(len(g.Encounters))
	gameOver := g.Encounters[random].Start(g)
	if gameOver {
		return true
	}
	time.Sleep(turnDelay * time.Millisecond)
	if g.Turn < loopTurns {
		g.Turn++
	} else {
		log.WithFields(log.Fields{"game": g}).Info("new loop started")
		g.Loop++
		g.Turn = 1
		g.Hero.SetHP(g.Hero.Stat("hp-max"))
		time.Sleep(loopDelay * time.Millisecond)
	}
	return false
}
