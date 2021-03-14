package main

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/urbanyeti/go-hero-game/character"
	"github.com/urbanyeti/go-hero-game/math"
)

// A CombatEncounter consists of a fight with a monster
type CombatEncounter struct {
	Monster character.Monster
}

type Fighter interface {
	ID() string
	Name() string
	HP() int
	Stat(string) int
	Weapon() (*character.Item, bool)
	TakeDamage(int)
}

// Start the fight
func (encounter CombatEncounter) Start(game *Game) bool {
	log.WithFields(log.Fields{"hero": game.Hero, "monster": encounter.Monster}).Info("combat encounter started")

	playersMove := true
	for i := 1; game.Hero.HP() > 0 && encounter.Monster.HP() > 0; i++ {
		if playersMove {
			totalDmg := calculateDamage(game.Hero, &encounter.Monster)
			encounter.Monster.TakeDamage(totalDmg)
		} else {
			totalDmg := calculateDamage(&encounter.Monster, game.Hero)
			game.Hero.TakeDamage(totalDmg)
		}
		playersMove = !playersMove
		time.Sleep(messageDelay * time.Millisecond)
	}

	if game.Hero.HP() <= 0 {
		log.WithFields(log.Fields{"hero": game.Hero.ID()}).Info("hero died")
		return true
	}

	log.WithFields(log.Fields{"hero": game.Hero.ID(), "monster": encounter.Monster.ID()}).Info("monster slayed")
	game.Hero.GainExp(encounter.Monster.Stat("exp"))
	return false
}

func calculateDamage(attacker Fighter, defender Fighter) int {
	baseDmg := attacker.Stat("atk") + defender.Stat("lvl")
	weaponDmg := 1
	if weapon, ok := attacker.Weapon(); ok {
		weaponDmg = rand.Intn(weapon.Stat("dmg-min")+weapon.Stat("dmg-max")) + (weapon.Stat("dmg-min"))
	}
	defenderDef := defender.Stat("def")
	totalDmg := math.MaxOf(baseDmg+weaponDmg-defenderDef, 0)
	log.WithFields(
		log.Fields{
			"attacker":    attacker,
			"defender":    defender,
			"baseDmg":     baseDmg,
			"weaponDmg":   weaponDmg,
			"defenderDef": defenderDef,
			"totalDmg":    totalDmg,
		}).Info("damage calculated")

	return totalDmg
}
