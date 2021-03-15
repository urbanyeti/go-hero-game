package character

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// ItemJSON is the DTO for items
type ItemJSON struct {
	ID    string
	Name  string
	Desc  string
	Stats map[string]int
	Tags  []string
}

// An Item is collected and may be used or equipped by a character
type Item struct {
	id    string
	name  string
	desc  string
	stats Stats
	tags  Tags
}

type Items map[string]*Item

// LoadItem generates an Item object from the DTO
func (i ItemJSON) LoadItem() Item {
	tags := make(Tags)
	for _, v := range i.Tags {
		tags[v] = true
	}

	return Item{
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

// HasTag confirms whether the item has the specified tag
func (item *Item) HasTag(id string) bool {
	if _, ok := item.tags[id]; ok {
		return true
	}
	return false
}

func (item Item) String() string {
	return fmt.Sprintf("[%v]", item.name)
}
