package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/urbanyeti/go-hero-game/characters"
)

// An Encounter can be started each turn
type Encounter interface {
	Start(hero *characters.Hero) error
}

// A CombatEncounter consists of a fight with a monster
type CombatEncounter struct {
}

// Start the fight
func (encounter CombatEncounter) Start(hero *characters.Hero) error {
	monster := characters.LoadGoblin()
	fmt.Printf("  - Combat: A wild %v appears!\n", monster.Name)

	playersMove := true
	for i := 1; hero.HP.Value > 0 && monster.HP.Value > 0; i++ {
		heroAtk := hero.Stat("atk")
		heroDef := hero.Stat("def")
		monsterAtk := monster.Stat("atk")
		monsterDef := monster.Stat("def")
		heroDamage := maxOf(heroAtk-monsterDef, 0)
		monsterDamage := maxOf(monsterAtk-heroDef, 0)

		if playersMove {
			fmt.Printf("    - Round %v: %v deals %v DMG! (%v ATK - %v DEF)\n", i, hero.Name, heroDamage, heroAtk, monsterDef)
			monster.HP.Value -= heroDamage

		} else {
			fmt.Printf("    - Round %v: %v deals %v DMG! (%v ATK - %v DEF)\n", i, monster.Name, monsterDamage, monsterAtk, heroDef)
			hero.HP.Value -= monsterDamage
		}
		fmt.Printf("      %v: %v | %v: %v\n", hero.Name, hero.HP, monster.Name, monster.HP)
		playersMove = !playersMove
		time.Sleep(1 * time.Second)
	}

	return nil
}

// A CutsceneEncounter consists of dialogue and stat changes
type CutsceneEncounter struct {
	Description string
}

// Start the CutsceneEncounter
func (encounter CutsceneEncounter) Start(hero *characters.Hero) error {
	fmt.Printf("  - Cutscene: %v\n    - ", encounter.Description)
	event := rand.Intn(6)
	switch event {
	case 0:
		hero.AddStat("atk", 2)
	case 1:
		hero.AddStat("def", 2)
	case 2:
		fmt.Printf("%v gains %v HP", hero.Name, 2)
		hero.HP.Value += 2
	case 3:
		hero.AddStat("atk", -2)
	case 4:
		hero.AddStat("def", -2)
	case 5:
		fmt.Printf("%v loses %v HP", hero.Name, 2)
		hero.HP.Value -= 2
	}
	fmt.Println()
	return nil
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
