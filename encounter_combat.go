package main

import (
	"fmt"
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
	fmt.Printf("  - Combat: A wild %v appears!\n", encounter.Monster.Name())
	log.WithFields(log.Fields{"hero": game.Hero, "monster": encounter.Monster}).Info("combat encounter started")

	playersMove := true
	for i := 1; game.Hero.HP() > 0 && encounter.Monster.HP() > 0; i++ {
		if playersMove {
			totalDmg, message := calculateDamage(game.Hero, &encounter.Monster)
			fmt.Printf("    - Round %v: %v\n", i, message)
			encounter.Monster.TakeDamage(totalDmg)
		} else {
			totalDmg, message := calculateDamage(&encounter.Monster, game.Hero)
			fmt.Printf("    - Round %v: %v\n", i, message)
			game.Hero.TakeDamage(totalDmg)
		}
		fmt.Printf("      %v: %v | %v: %v\n", game.Hero.Name(), game.Hero.HP(), encounter.Monster.Name(), encounter.Monster.HP())
		playersMove = !playersMove
		time.Sleep(messageDelay * time.Millisecond)
	}

	if game.Hero.HP() <= 0 {
		log.WithFields(log.Fields{"hero": game.Hero.ID()}).Info("hero died")
		fmt.Println()
		fmt.Println("*****************")
		fmt.Println("X YOU HAVE DIED X")
		fmt.Println("*****************")
		return true
	}

	fmt.Printf("    %v slays the %v!\n", game.Hero.Name(), encounter.Monster.Name())
	log.WithFields(log.Fields{"hero": game.Hero.ID(), "monster": encounter.Monster.ID()}).Info("monster slayed")
	game.Hero.GainExp(encounter.Monster.Stat("exp"))
	return false
}

func calculateDamage(attacker Fighter, defender Fighter) (int, string) {
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

	return totalDmg, fmt.Sprintf("%v deals %v DMG! (%v DMG - %v DEF)", attacker.Name(), totalDmg, baseDmg+weaponDmg, defenderDef)
}
