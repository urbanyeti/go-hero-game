package character

import (
	"fmt"
	"math/rand"

	log "github.com/sirupsen/logrus"
)

// ItemJSON is the DTO for items
type ItemJSON struct {
	ID    string         `json:"id"`
	Name  string         `json:"name"`
	Desc  string         `json:"desc"`
	Stats map[string]int `json:"stats,omitempty"`
	Tags  []string       `json:"tags,omitempty"`
}

// An Item is collected and may be used or equipped by a character
type Item struct {
	id    string
	name  string
	desc  string
	stats Stats
	tags  Tags
}

type LoadedItems map[string]*Item
type Equipment map[string]*Item

// LoadItem generates an Item object from the DTO
func (i ItemJSON) LoadItem() *Item {
	tags := make(Tags)
	for _, v := range i.Tags {
		tags[v] = true
	}

	return &Item{
		id:    i.ID,
		name:  i.Name,
		desc:  i.Desc,
		stats: i.Stats,
		tags:  tags,
	}
}

// ID returns item's ID
func (i *Item) ID() string {
	return i.id
}

// Name returns item's name
func (i *Item) Name() string {
	return i.name
}

// Desc returns item's description
func (i *Item) Desc() string {
	return i.desc
}

// Stat retrieves the current stat value
func (item *Item) Stat(statID string) int {
	if stat, ok := item.stats[statID]; ok {
		return stat
	}

	log.Warn("cannot retrieve unknown stat '%v'", statID)
	return 0
}

// Stat retrieves the current stat value
func (item *Item) StatTry(statID string) int {
	if stat, ok := item.stats[statID]; ok {
		return stat
	}

	return 0
}

// HasTag confirms whether the item has the specified tag
func (item *Item) HasTag(id string) bool {
	if _, ok := item.tags[id]; ok {
		return true
	}
	return false
}

func (li LoadedItems) GetRandomItem() *Item {
	random := rand.Intn((len(li)))
	i := 0
	for _, v := range li {
		if i == random {
			return v.Clone()
		}
		i++
	}
	return &Item{}
}

func (i Item) Clone() *Item {
	n := i
	n.stats = make(Stats, len(i.stats))
	for k, v := range i.stats {
		n.stats[k] = v
	}
	n.tags = make(Tags, len(i.tags))
	for k, v := range i.tags {
		n.tags[k] = v
	}

	return &n
}

func (item Item) String() string {
	return fmt.Sprintf("[%v]", item.name)
}
