package main

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/urbanyeti/go-hero-game/math"
)

// An Encounter can be started each turn
type Encounter interface {
	Start(game *Game) bool
}

// A CutsceneEncounter consists of dialogue and stat changes
type CutsceneEncounter struct {
	Description string
}

// Start the CutsceneEncounter
func (encounter CutsceneEncounter) Start(game *Game) bool {
	hero := game.Hero

	log.WithFields(log.Fields{"hero": game.Hero, "encounter": encounter}).Info("cutscene encounter started")
	event := rand.Intn(6)
	switch event {
	case 0:
		hero.AddStat("atk", 2)
	case 1:
		hero.AddStat("def", 2)
	case 2:
		hero.AddStat("hp-max", 2)
		hero.Heal(2)
	case 3:
		hero.AddStat("atk", -2)
	case 4:
		hero.AddStat("def", -2)
	case 5:
		hero.AddStat("hp-max", -2)
		hero.SetHP(math.MinOf(hero.HP(), hero.Stat("hp-max")))
	}
	time.Sleep(messageDelay * time.Millisecond)
	return false
}
