package character

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// An Ability is collected and may be used or equipped by the Hero.
type Ability struct {
	ID    string
	Name  string
	Desc  string
	Tags  Tags
	Stats Stats
}

func (a Ability) String() string {
	return fmt.Sprintf("[%v]", a.Name)
}

// Abilities is a collection of Item objects
type Abilities map[string]*Ability

// LoadedAbilities is a map of loaded Item objects
type LoadedAbilities map[string]Ability

// Stat retrieves the current stat value
func (ability *Ability) Stat(statID string) int {
	if stat, ok := ability.Stats[statID]; ok {
		return stat
	}

	log.Warn("cannot retrieve unknown stat '%v'", statID)
	return 0
}

// HasTag confirms whether the ability has the specified tag
func (ability *Ability) HasTag(id string) bool {
	if _, ok := ability.Tags[id]; ok {
		return true
	}
	return false
}
