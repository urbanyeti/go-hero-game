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
	HP          int
	Stats       Stats
	Equipment   Equipment
}

func (monster Monster) String() string {
	return fmt.Sprintf("%v - %v | %v | %v", monster.Name, monster.HP, monster.Stats, monster.Equipment)
}

// LoadGoblin creates a default goblin
func LoadGoblin() Monster {
	jsonFile, err := os.Open("./characters/monsters/monsters.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var monsters map[string]Monster
	json.Unmarshal(byteValue, &monsters)
	return monsters["monster-goblin"]
}

// Stat retrieves the current stat value
func (monster *Monster) Stat(statID string) int {
	if stat, ok := monster.Stats[statID]; ok {
		return stat
	}

	log.Printf("cannot retrieve to unknown stat '%v'", statID)
	return 0
}
