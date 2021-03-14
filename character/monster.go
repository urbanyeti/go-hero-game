package character

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

// Monster is an agent in the world
type Monster struct {
	Character
}

func (c CharacterJSON) Load() Monster {
	return Monster{Character{
		id:    c.ID,
		name:  c.Name,
		desc:  c.Desc,
		hp:    c.HP,
		stats: c.Stats,
		items: c.Items,
	}}
}

// Monsters is a map of loaded monsters
type LoadedMonsters map[string]Monster

// LoadMonsters grabs all the monster data from json
func LoadMonsters() LoadedMonsters {
	jsonFile, err := os.Open("./character/json/monsters.json")
	if err != nil {
		log.Error(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var jsonCharacters map[string]*CharacterJSON
	json.Unmarshal(byteValue, &jsonCharacters)
	monsters := make(LoadedMonsters)
	for key, character := range jsonCharacters {
		monsters[key] = Monster(character.Load())
	}

	return monsters
}

// Keys lists all the keys from the Monsters map
func (monsters LoadedMonsters) Keys() []string {
	keys := []string{}
	for key := range monsters {
		keys = append(keys, key)
	}
	return keys
}

// Monster grabs a new instance of the specify monster
func (monsters LoadedMonsters) Monster(monsterID string) Monster {
	if monster, ok := monsters[monsterID]; ok {
		return monster
	}
	log.Warn("Cannot retrieve unknown monster '%v'", monsterID)
	return Monster{}
}
