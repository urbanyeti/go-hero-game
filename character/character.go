package character

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/urbanyeti/go-hero-game/math"
)

// A character is an agent in the world
type Character struct {
	id    string
	name  string
	desc  string
	hp    int
	stats Stats
	items Items
}

func NewCharacter(id string, name string, desc string) Character {
	c := Character{id: id, name: name, desc: desc}

	return c
}

// AddStat adds to the specified Stat value
func (c *Character) AddStat(statID string, value int) {
	if stat, ok := c.stats[statID]; ok {
		c.stats[statID] = math.MaxOf(stat+value, 1)
		if value > 0 {
			log.Info("%v gains %v %v", c.name, value, statID)
		} else {
			log.Info("%v loses %v %v", c.name, value, statID)
		}
	} else {
		log.Warn("Cannot add to missing stat '%v'", statID)
	}
}

// SetStat sets the specified Stat value
func (c *Character) SetStat(statID string, value int) {
	if _, ok := c.stats[statID]; ok {
		c.stats[statID] = value
	} else {
		log.Warn("Cannot set missing stat '%v'", statID)
	}
}

// Stat retrieves the current stat value
func (c *Character) Stat(statID string) int {
	if stat, ok := c.stats[statID]; ok {
		return stat
	}

	log.Warn("Cannot retrieve missing stat '%v'", statID)
	return 0
}

// Weapon returns equipped item in right arm
func (c *Character) Weapon() (*Item, bool) {
	if item, ok := c.items["arm-r"]; ok {
		return item, true
	}
	return nil, false
}

// Name returns hero's Name
func (c *Character) Name() string {
	return c.name
}

// HP returns the hero's HP
func (c *Character) HP() int {
	return c.hp
}

// SetHP sets the hero's HP
func (c *Character) SetHP(hp int) {
	c.hp = hp
	log.Info("%v's HP set to {hp}", c.name, hp)
}

func (c *Character) TakeDamage(dmg int) {
	c.hp -= dmg
	log.Info("%v took %v damage", c.name, dmg)
}

func (c *Character) Heal(health int) {
	c.hp += health
	log.Info("%v gained %v health", c.name, health)
}

func (c *Character) Equip(items ...Item) {
	for _, item := range items {

		if item.HasTag("equip") {
			if item.HasTag("weapon") && item.HasTag("arm") {
				c.items["arm-r"] = &item
				log.WithFields(log.Fields{"character": c.id, "item": item.ID}).Info("equipped weapon")
			} else if item.HasTag("armor") {
				log.WithFields(log.Fields{"character": c.id, "item": item.ID}).Info("equipped armor")
			}
		}
	}
}

func (c Character) String() string {
	return fmt.Sprintf("%v - HP: %v | %v | %v", c.name, c.hp, c.stats, c.items)
}
