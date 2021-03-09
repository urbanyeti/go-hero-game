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

// Monsters is a map of loaded monsters
type Monsters map[string]Monster

// LoadMonsters grabs all the monster data from json
func LoadMonsters() Monsters {
	jsonFile, err := os.Open("./characters/json/monsters.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var monsters Monsters
	json.Unmarshal(byteValue, &monsters)
	return monsters
}

// Keys lists all the keys from the Monsters map
func (monsters Monsters) Keys() []string {
	keys := []string{}
	for key := range monsters {
		keys = append(keys, key)
	}
	return keys
}

// Monster grabs a pointer to the specify monster
func (monsters Monsters) Monster(monsterID string) Monster {
	if monster, ok := monsters[monsterID]; ok {
		return monster
	}
	log.Printf("cannot retrieve unknown monster '%v'", monsterID)
	return Monster{}
}

// Weapon returns equipped item in right arm
func (monster *Monster) Weapon() (*Item, bool) {
	if item, ok := monster.Equipment["arm-r"]; ok {
		return item, true
	}
	return nil, false
}

// GetName returns the Name
func (monster *Monster) GetName() string {
	return monster.Name
}

// Stat retrieves the current stat value
func (monster *Monster) Stat(statID string) int {
	if stat, ok := monster.Stats[statID]; ok {
		return stat
	}

	log.Printf("cannot retrieve to unknown stat '%v'", statID)
	return 0
}
