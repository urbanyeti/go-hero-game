package character

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// AbilityJSON is the DTO for abilities
type AbilityJSON struct {
	ID    string
	Name  string
	Desc  string
	Stats map[string]int
	Tags  []string
}

// An Ability is a skill or effect that can be used by a character
type Ability struct {
	id    string
	name  string
	desc  string
	stats Stats
	tags  Tags
}

type Abilities map[string]*Ability

// LoadAbility generates an Ability object from the DTO
func (a AbilityJSON) LoadAbility() Ability {
	tags := make(Tags)
	for _, v := range a.Tags {
		tags[v] = true
	}

	return Ability{
		id:    a.ID,
		name:  a.Name,
		desc:  a.Desc,
		stats: a.Stats,
		tags:  tags,
	}
}

// ID returns ability's ID
func (a *Ability) ID() string {
	return a.id
}

// Name returns ability's name
func (a *Ability) Name() string {
	return a.name
}

// Desc returns ability's description
func (a *Ability) Desc() string {
	return a.desc
}

// Stat retrieves the current stat value
func (ability *Ability) Stat(statID string) int {
	if stat, ok := ability.stats[statID]; ok {
		return stat
	}

	log.Warn("cannot retrieve unknown stat '%v'", statID)
	return 0
}

// HasTag confirms whether the ability has the specified tag
func (ability *Ability) HasTag(id string) bool {
	if _, ok := ability.tags[id]; ok {
		return true
	}
	return false
}

func (a Ability) String() string {
	return fmt.Sprintf("[%v]", a.name)
}
