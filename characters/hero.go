package characters

import (
	"fmt"
	"log"
)

// The Hero is the main player character.
type Hero struct {
	ID          string
	Name        string
	Description string
	HP          int
	Stats       Stats
	Equipment   Equipment
}

func (hero Hero) String() string {
	return fmt.Sprintf("%v - HP: %v | %v | %v", hero.Name, hero.HP, hero.Stats, hero.Equipment)
}

// SetDefaultStats initializes the default stats for the hero
func (hero *Hero) SetDefaultStats() {
	if hero.HP == 0 {
		hero.HP = 100
	}

	if len(hero.Stats) == 0 {
		hero.Stats = Stats{
			"hp-max":   100,
			"atk":      5,
			"def":      5,
			"eva":      0,
			"lvl":      1,
			"exp":      0,
			"exp-next": 10,
		}
	}
}

// GainExp increases the exp to Hero.
// The hero may level up
func (hero *Hero) GainExp(exp int) {
	totalExp := hero.Stat("exp") + exp
	fmt.Print("    ")
	if totalExp >= hero.Stat("exp-next") {
		hero.AddStat("lvl", totalExp/hero.Stat("exp-next"))
		hero.SetStat("exp", totalExp%hero.Stat("exp-next"))
		fmt.Printf(" | Level %v (%v/%v)", hero.Stat("lvl"), hero.Stat("exp"), hero.Stat("exp-next"))
	} else {
		hero.AddStat("exp", exp)
		fmt.Printf(" | Level %v (%v/%v)", hero.Stat("lvl"), hero.Stat("exp"), hero.Stat("exp-next"))
	}
	fmt.Print("\n")
}

// AddStat adds to the specified Stat value
func (hero *Hero) AddStat(statID string, value int) {
	if stat, ok := hero.Stats[statID]; ok {
		hero.Stats[statID] = maxOf(stat+value, 1)
		if value > 0 {
			fmt.Printf("%v gains %v %v", hero.Name, value, statID)
		} else {
			fmt.Printf("%v loses %v %v", hero.Name, value, statID)
		}
	} else {
		log.Printf("cannot add to unknown stat '%v'", statID)
	}
}

// SetStat sets the specified Stat value
func (hero *Hero) SetStat(statID string, value int) {
	if _, ok := hero.Stats[statID]; ok {
		hero.Stats[statID] = value
	} else {
		log.Printf("cannot set unknown stat '%v'", statID)
	}
}

// Stat retrieves the current stat value
func (hero *Hero) Stat(statID string) int {
	if stat, ok := hero.Stats[statID]; ok {
		return stat
	}

	log.Printf("cannot retrieve unknown stat '%v'", statID)
	return 0
}

// Weapon returns equipped item in right arm
func (hero *Hero) Weapon() (*Item, bool) {
	if item, ok := hero.Equipment["arm-r"]; ok {
		return item, true
	}
	return nil, false
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

func minOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}
