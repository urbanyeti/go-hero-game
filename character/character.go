package character

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/urbanyeti/go-hero-game/math"
)

// A character is an agent in the world
type Character struct {
	id        string
	name      string
	desc      string
	hp        int
	stats     Stats
	items     Items
	abilities Abilities
}

type CharacterJSON struct {
	ID        string
	Name      string
	Desc      string
	HP        int
	Stats     map[string]int
	Items     []string
	Abilities []string
}

func NewCharacter(id string, name string, desc string) Character {
	c := Character{id: id, name: name, desc: desc}

	return c
}

// ID returns character's ID
func (c *Character) ID() string {
	return c.id
}

// Name returns character's name
func (c *Character) Name() string {
	return c.name
}

// Desc returns character's description
func (c *Character) Desc() string {
	return c.desc
}

// HP returns the character's current HP
func (c *Character) HP() int {
	return c.hp
}

// AddStat adds to the specified Stat value
func (c *Character) AddStat(statID string, value int) {
	if stat, ok := c.stats[statID]; ok {
		c.stats[statID] = math.MaxOf(stat+value, 1)
		log.WithFields(log.Fields{"character": c.id, "stat": statID, "old": stat, "new": c.stats[statID]}).Info("stat modified")
	} else {
		log.WithFields(log.Fields{"statID": statID}).Warn("cannot add missing stat '%v'", statID)
	}
}

// SetStat sets the specified Stat value
func (c *Character) SetStat(statID string, value int) {
	if _, ok := c.stats[statID]; ok {
		c.stats[statID] = value
	} else {
		log.WithFields(log.Fields{"statID": statID}).Warn("cannot set missing stat '%v'", statID)
	}
}

// Stat retrieves the current stat value
func (c *Character) Stat(statID string) int {
	if stat, ok := c.stats[statID]; ok {
		return stat
	}

	log.WithFields(log.Fields{"statID": statID}).Warn("cannot retrieve missing stat '%v'", statID)
	return 0
}

// Weapon returns equipped item in right arm
func (c *Character) Weapon() (*Item, bool) {
	if item, ok := c.items["arm-r"]; ok {
		return item, true
	}
	return nil, false
}

// SetHP sets the character's HP to the specified value
func (c *Character) SetHP(hp int) {
	c.hp = hp
	log.WithFields(log.Fields{"character": c.id, "hp": hp}).Info("HP set to new value")
}

func (c *Character) TakeDamage(dmg int) {
	c.hp -= dmg
	log.WithFields(log.Fields{"character": c.id, "dmg": dmg}).Info("damage taken")
}

func (c *Character) Heal(health int) {
	c.hp += health
	log.WithFields(log.Fields{"character": c.id, "health": health}).Info("health gained")
}

func (c *Character) Equip(items ...Item) {
	for _, item := range items {

		if item.HasTag("equip") {
			if item.HasTag("weapon") && item.HasTag("arm") {
				c.items["arm-r"] = &item
				log.WithFields(log.Fields{"character": c.id, "item": item.ID()}).Info("equipped weapon")
			} else if item.HasTag("armor") {
				if item.HasTag("torso") {
					c.items["torso"] = &item
					log.WithFields(log.Fields{"character": c.id, "item": item.ID()}).Info("equipped torso")
				} else if item.HasTag("feet") {
					c.items["feet"] = &item
					log.WithFields(log.Fields{"character": c.id, "item": item.ID()}).Info("equipped feet")
				}
			}
		}
	}
}

func (c Character) String() string {
	return fmt.Sprintf("%v - HP: %v | %v | %v | %v", c.name, c.hp, c.stats, c.abilities, c.items)
}
