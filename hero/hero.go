package hero

import (
	"fmt"
	"log"
)

// The Hero is the main player character.
type Hero struct {
	ID          string
	Name        string
	Description string
	HP          *Stat
	Stats       Stats
	Equipment   Equipment
}

func (hero Hero) String() string {
	return fmt.Sprintf("%v - %v | %v | %v", hero.Name, hero.HP, hero.Stats, hero.Equipment)
}

// SetDefaultStats initializes the default stats for the hero
func (hero *Hero) SetDefaultStats() {
	if hero.HP == nil {
		hero.HP = &Stat{
			ID:          "stat-hp",
			Name:        "HP",
			Description: "health points",
			Value:       100}
	}

	if len(hero.Stats) == 0 {
		hero.Stats = Stats{
			"stat-atk": &Stat{
				ID:          "stat-atk",
				Name:        "ATK",
				Description: "attack stat",
				Value:       5},
			"stat-def": &Stat{
				ID:          "stat-def",
				Name:        "DEF",
				Description: "defense stat",
				Value:       5},
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

// AddStat adds to the specified Stat value
func (hero *Hero) AddStat(statID string, value int) {
	if stat, ok := hero.Stats["stat-"+statID]; ok {
		stat.Value += value
	} else {
		log.Printf("cannot add to unknown stat '%v'", statID)
	}
}
