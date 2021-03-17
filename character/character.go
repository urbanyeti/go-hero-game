package character

import (
	"fmt"
	"math/rand"

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
	equipment Equipment
	abilities Abilities
	items     []*Item
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

// Def calculates the character's defense stat
func (c *Character) Def() int {
	def := c.Stat("def") + c.Stat("lvl")
	for _, e := range c.equipment {
		def += e.StatTry("def")
	}

	return def
}

// Agi calculates the character's speed stat
func (c *Character) Agi() int {
	agi := c.Stat("agi")
	for _, a := range c.equipment {
		agi += a.StatTry("agi")
	}

	return agi
}

// AddStat adds to the specified Stat value
func (c *Character) AddStat(statID string, value int) {
	if stat, ok := c.stats[statID]; ok {
		c.stats[statID] = math.MaxOf(stat+value, 1)
		log.WithFields(log.Fields{"character": c.id, "stat": statID, "old": stat, "new": c.stats[statID]}).Info("stat modified")
	} else {
		log.WithFields(log.Fields{"statID": statID}).Warn("cannot add missing stat")
	}
}

// AddItem adds the specified item to the character
func (c *Character) AddItem(item *Item) {
	c.items = append(c.items, item)
	log.WithFields(log.Fields{"item": item.id, "character": c.id}).Info("added item")
}

// SetStat sets the specified Stat value
func (c *Character) SetStat(statID string, value int) {
	if _, ok := c.stats[statID]; ok {
		c.stats[statID] = value
	} else {
		log.WithFields(log.Fields{"statID": statID}).Warn("cannot set missing stat")
	}
}

// Stat retrieves the current stat value
func (c *Character) Stat(statID string) int {
	if stat, ok := c.stats[statID]; ok {
		return stat
	}

	log.WithFields(log.Fields{"statID": statID}).Warn("cannot retrieve missing stat")
	return 0
}

// Armor returns equipped armor
func (c *Character) Armor() ([]*Item, bool) {
	armor := []*Item{}
	if item, ok := c.equipment["torso"]; ok {
		armor = append(armor, item)
	}
	if item, ok := c.equipment["feet"]; ok {
		armor = append(armor, item)
	}

	return armor, len(armor) > 0
}

// Weapons returns equipped weapons
func (c *Character) Weapons() ([]*Item, bool) {
	weapons := []*Item{}
	if item, ok := c.equipment["arm-r"]; ok {
		weapons = append(weapons, item)
	}
	if item, ok := c.equipment["arm-l"]; ok {
		weapons = append(weapons, item)
	}

	return weapons, len(weapons) > 0
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

type Damager interface {
	ID() string
	Name() string
	Desc() string
	Stat(string) int
}

func (c *Character) SelectDamager() Damager {
	attacks := c.Damagers()
	random := rand.Intn(len(attacks))
	return attacks[random]
}

func (c *Character) Damagers() []Damager {
	attacks := []Damager{}
	if weapons, ok := c.Weapons(); ok {
		for _, v := range weapons {
			attacks = append(attacks, v)
		}
	}

	for _, v := range c.abilities {
		attacks = append(attacks, v)
	}

	return attacks
}

func (c *Character) Equip(items ...*Item) {
	for _, item := range items {

		if item.HasTag("equip") {
			if item.HasTag("weapon") && item.HasTag("arm") {
				c.equipment["arm-r"] = item
				log.WithFields(log.Fields{"character": c.id, "item": item.ID()}).Info("equipped weapon")
			} else if item.HasTag("armor") {
				if item.HasTag("torso") {
					c.equipment["torso"] = item
					log.WithFields(log.Fields{"character": c.id, "item": item.ID()}).Info("equipped torso")
				} else if item.HasTag("feet") {
					c.equipment["feet"] = item
					log.WithFields(log.Fields{"character": c.id, "item": item.ID()}).Info("equipped feet")
				}
			}
		}
	}
}

func (c Character) String() string {
	return fmt.Sprintf("%v - HP: %v | %v | %v | %v", c.name, c.hp, c.stats, c.abilities, c.equipment)
}
