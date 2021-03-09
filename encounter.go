package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/urbanyeti/go-hero-game/characters"
)

// An Encounter can be started each turn
type Encounter interface {
	Start(game *Game) bool
}

// A CombatEncounter consists of a fight with a monster
type CombatEncounter struct {
	Monster characters.Monster
}

// Start the fight
func (encounter CombatEncounter) Start(game *Game) bool {
	hero := game.Hero
	fmt.Printf("  - Combat: A wild %v appears!\n", encounter.Monster.Name)

	playersMove := true
	for i := 1; hero.HP > 0 && encounter.Monster.HP > 0; i++ {
		if playersMove {
			totalDmg, message := calculateDamage(game.Hero, &encounter.Monster)
			fmt.Printf("    - Round %v: %v\n", i, message)
			encounter.Monster.HP -= totalDmg
		} else {
			totalDmg, message := calculateDamage(&encounter.Monster, game.Hero)
			fmt.Printf("    - Round %v: %v\n", i, message)
			hero.HP -= totalDmg
		}
		fmt.Printf("      %v: %v | %v: %v\n", hero.Name, hero.HP, encounter.Monster.Name, encounter.Monster.HP)
		playersMove = !playersMove
		time.Sleep(messageDelay * time.Millisecond)
	}

	if hero.HP <= 0 {
		fmt.Println()
		fmt.Println("*****************")
		fmt.Println("X YOU HAVE DIED X")
		fmt.Println("*****************")
		return true
	}

	fmt.Printf("    %v slays the %v!\n", hero.Name, encounter.Monster.Name)
	hero.GainExp(encounter.Monster.Stat("exp"))
	return false
}

type combatant interface {
	GetName() string
	Stat(string) int
	Weapon() (*characters.Item, bool)
}

func calculateDamage(attacker combatant, defender combatant) (int, string) {
	baseDmg := attacker.Stat("atk") + attacker.Stat("lvl")
	weaponDmg := 1
	if weapon, ok := attacker.Weapon(); ok {
		weaponDmg = rand.Intn(weapon.Stat("dmg-min")+weapon.Stat("dmg-max")) + (weapon.Stat("dmg-min"))
	}
	defenderDef := defender.Stat("def")
	totalDmg := maxOf(baseDmg+weaponDmg-defenderDef, 0)

	return totalDmg, fmt.Sprintf("%v deals %v DMG! (%v DMG - %v DEF)", attacker.GetName(), totalDmg, baseDmg+weaponDmg, defenderDef)
}

// A CutsceneEncounter consists of dialogue and stat changes
type CutsceneEncounter struct {
	Description string
}

// Start the CutsceneEncounter
func (encounter CutsceneEncounter) Start(game *Game) bool {
	hero := game.Hero
	fmt.Printf("  - Cutscene: %v\n    - ", encounter.Description)
	event := rand.Intn(6)
	switch event {
	case 0:
		hero.AddStat("atk", 2)
	case 1:
		hero.AddStat("def", 2)
	case 2:
		hero.AddStat("hp-max", 2)
		hero.HP += 2
	case 3:
		hero.AddStat("atk", -2)
	case 4:
		hero.AddStat("def", -2)
	case 5:
		hero.AddStat("hp-max", -2)
		hero.HP = minOf(hero.HP, hero.Stat("hp-max"))
	}
	fmt.Println()
	time.Sleep(messageDelay * time.Millisecond)
	return false
}
