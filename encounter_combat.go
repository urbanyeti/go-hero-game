package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/urbanyeti/go-hero-game/character"
)

// A CombatEncounter consists of a fight with a monster
type CombatEncounter struct {
	Monster character.Monster
}

type fighter interface {
	Name() string
	HP() int
	Stat(string) int
	Weapon() (*character.Item, bool)
	TakeDamage(int)
}

// Start the fight
func (encounter CombatEncounter) Start(game *Game) bool {
	hero := game.Hero
	fmt.Printf("  - Combat: A wild %v appears!\n", encounter.Monster.Name())

	playersMove := true
	for i := 1; hero.HP() > 0 && encounter.Monster.HP() > 0; i++ {
		if playersMove {
			totalDmg, message := calculateDamage(game.Hero, &encounter.Monster)
			fmt.Printf("    - Round %v: %v\n", i, message)
			encounter.Monster.TakeDamage(totalDmg)
		} else {
			totalDmg, message := calculateDamage(&encounter.Monster, game.Hero)
			fmt.Printf("    - Round %v: %v\n", i, message)
			hero.TakeDamage(totalDmg)
		}
		fmt.Printf("      %v: %v | %v: %v\n", hero.Name(), hero.HP(), encounter.Monster.Name(), encounter.Monster.HP())
		playersMove = !playersMove
		time.Sleep(messageDelay * time.Millisecond)
	}

	if hero.HP() <= 0 {
		fmt.Println()
		fmt.Println("*****************")
		fmt.Println("X YOU HAVE DIED X")
		fmt.Println("*****************")
		return true
	}

	fmt.Printf("    %v slays the %v!\n", hero.Name(), encounter.Monster.Name())
	hero.GainExp(encounter.Monster.Stat("exp"))
	return false
}

func calculateDamage(atk fighter, def fighter) (int, string) {
	baseDmg := atk.Stat("atk") + def.Stat("lvl")
	weaponDmg := 1
	if weapon, ok := atk.Weapon(); ok {
		weaponDmg = rand.Intn(weapon.Stat("dmg-min")+weapon.Stat("dmg-max")) + (weapon.Stat("dmg-min"))
	}
	defenderDef := def.Stat("def")
	totalDmg := maxOf(baseDmg+weaponDmg-defenderDef, 0)

	return totalDmg, fmt.Sprintf("%v deals %v DMG! (%v DMG - %v DEF)", atk.Name(), totalDmg, baseDmg+weaponDmg, defenderDef)
}
