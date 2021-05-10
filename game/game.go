package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/urbanyeti/go-hero-game/character"
	"github.com/urbanyeti/go-hero-game/graphics"
)

const loopTurns int = 10
const messageDelay = 100
const attackDelay = 500
const turnDelay = 300
const loopDelay = 0
const GRAPHICS_PATH = `graphics/content/`

// Game contains state and data about the game session
type Game struct {
	Hero       *character.Hero
	Loop       int
	Turn       int
	Encounters []Encounter
	Monsters   map[string]character.Monster
	Abilities  map[string]character.Ability
	Items      character.LoadedItems
	Graphics   *graphics.Graphics
	Animations map[string]*graphics.Animations
}

func (g Game) String() string {
	return fmt.Sprintf("Loop: %v Turn: %v | %v", g.Loop, g.Turn, g.Hero)
}

// func (g *Game) GetRandomItem() {
// 	request := grpc.ItemRequest{LoopNumber: int(g.Loop), Level: (int)(g.Hero.Stat("lvl"))}
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	g.Server.GetRandomItem(ctx, &request)
// }

// Initialize sets up the Game
func (g *Game) Initialize() {
	g.LoadContent()
	g.Loop = 1
	g.Turn = 1
	g.Hero = character.NewHero("hero-dan", "Dan", "cool hero", g.Animations["Knight_02"])
	(*g.Graphics.Animations)[g.Hero.ID()] = (*g.Hero.Animations)["01-Idle"]
	g.SetDefaultEquipment(g.Hero)

	g.Encounters = []Encounter{}
	g.Encounters = append(g.Encounters, CutsceneEncounter{[]string{"A sppooky encounter...", "A random act of chaos", "A magical experience"}})
	monsters := []character.Monster{}
	for _, m := range g.Monsters {
		monsters = append(monsters, m)
	}
	g.Encounters = append(g.Encounters, CombatEncounter{monsters})
}

// LoadContent loads resources used by the Game
func (g *Game) LoadContent() {
	g.LoadGraphics()
	g.LoadAbilities()
	g.LoadItems()
	g.LoadMonsters()
}

func (g *Game) LoadGraphics() {
	g.Graphics = graphics.NewGraphics()
	g.Animations = map[string]*graphics.Animations{}

	g.LoadCharacterGraphics("Knight_02", false)
	g.LoadCharacterGraphics("Goblin_02", true)
}

func (g *Game) LoadCharacterGraphics(f string, isEnemy bool) {
	animations := graphics.Animations{}
	animFolders, err := ioutil.ReadDir(GRAPHICS_PATH + f)
	if err != nil {
		log.Fatal()
	}
	for _, animFolder := range animFolders {
		animations[animFolder.Name()] = graphics.NewAnimation(GRAPHICS_PATH+f+"/"+animFolder.Name()+"/", isEnemy)
	}
	g.Animations[f] = &animations
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
		g.Monsters[c.ID] = character.Monster(c.LoadMonster(g.Abilities, g.Items, g.Animations["Goblin_02"]))
	}
}

// SetDefaultEquipment initializes the default equipment for the hero
func (g *Game) SetDefaultEquipment(hero *character.Hero) {
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
