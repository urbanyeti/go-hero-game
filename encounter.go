package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/urbanyeti/go-hero-game/characters"
)

// An Encounter can be started each turn
type Encounter interface {
	Start(game *Game) error
}

// A CombatEncounter consists of a fight with a monster
type CombatEncounter struct {
	Monster characters.Monster
}

// Start the fight
func (encounter CombatEncounter) Start(game *Game) error {
	hero := game.Hero
	fmt.Printf("  - Combat: A wild %v appears!\n", encounter.Monster.Name)

	playersMove := true
	for i := 1; hero.HP > 0 && encounter.Monster.HP > 0; i++ {
		heroAtk := hero.Stat("atk") + hero.Stat("lvl")
		weaponDmg := 1
		if weapon, ok := hero.Weapon(); ok {
			weaponDmg = rand.Intn(weapon.Stat("dmg-min")+weapon.Stat("dmg-max")) + (weapon.Stat("dmg-min"))
		}
		heroDef := hero.Stat("def")
		monsterAtk := encounter.Monster.Stat("atk")
		monsterDef := encounter.Monster.Stat("def")
		heroDmg := maxOf(heroAtk+weaponDmg-monsterDef, 0)
		monsterDmg := maxOf(monsterAtk-heroDef, 0)

		if playersMove {
			fmt.Printf("    - Round %v: %v deals %v DMG! (%v DMG - %v DEF)\n", i, hero.Name, heroDmg, heroAtk+weaponDmg, monsterDef)
			encounter.Monster.HP -= heroDmg

		} else {
			fmt.Printf("    - Round %v: %v deals %v DMG! (%v DMG - %v DEF)\n", i, encounter.Monster.Name, monsterDmg, monsterAtk, heroDef)
			hero.HP -= monsterDmg
		}
		fmt.Printf("      %v: %v | %v: %v\n", hero.Name, hero.HP, encounter.Monster.Name, encounter.Monster.HP)
		playersMove = !playersMove
		time.Sleep(messageDelay * time.Millisecond)
	}
	fmt.Printf("    %v slays the %v!\n", hero.Name, encounter.Monster.Name)
	hero.GainExp(encounter.Monster.Stat("exp"))

	return nil
}

// A CutsceneEncounter consists of dialogue and stat changes
type CutsceneEncounter struct {
	Description string
}

// Start the CutsceneEncounter
func (encounter CutsceneEncounter) Start(game *Game) error {
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
	return nil
}
