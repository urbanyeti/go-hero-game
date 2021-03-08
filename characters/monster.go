package characters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Monster for combat encounters
type Monster struct {
	ID          string
	Name        string
	Description string
	HP          *Stat
	Stats       Stats
	Equipment   Equipment
}

func (monster Monster) String() string {
	return fmt.Sprintf("%v - %v | %v | %v", monster.Name, monster.HP, monster.Stats, monster.Equipment)
}

// LoadGoblin creates a default goblin
func LoadGoblin() Monster {
	jsonFile, err := os.Open("./characters/monsters/monster-goblin.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var goblin Monster
	json.Unmarshal(byteValue, &goblin)
	return goblin
}

// Stat retrieves the current stat value
func (monster *Monster) Stat(statID string) int {
	if stat, ok := monster.Stats["stat-"+statID]; ok {
		return stat.Value
	}

	log.Printf("cannot retrieve to unknown stat '%v'", statID)
	return 0
}
