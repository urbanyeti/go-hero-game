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
			"atk":      5,
			"def":      5,
			"lvl":      1,
			"exp":      0,
			"exp-next": 10,
		}
	}
}

// SetDefaultEquipment initializes the default equipment for the hero
func (hero *Hero) SetDefaultEquipment() {
	if len(hero.Equipment) == 0 {
		hero.Equipment = Equipment{
			"item-sword1": &Item{
				ID:          "item-sword1",
				Name:        "Short Sword",
				Description: "a beginner's basic short sword"},
			"item-armor1": &Item{
				ID:          "item-armor1",
				Name:        "Basic Armor",
				Description: "a beginner's basic set of armor"},
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
		fmt.Printf(" (%v/%v)", hero.Stat("exp"), hero.Stat("exp-next"))
	} else {
		hero.AddStat("exp", exp)
		fmt.Printf(" (%v/%v)", hero.Stat("exp"), hero.Stat("exp-next"))
	}
	fmt.Print("\n")
}

// AddStat adds to the specified Stat value
func (hero *Hero) AddStat(statID string, value int) {
	if stat, ok := hero.Stats[statID]; ok {
		hero.Stats[statID] = stat + value
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
