package character

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

// An Item is collected and may be used or equipped by the Hero.
type Item struct {
	ID    string
	Name  string
	Desc  string
	Tags  Tags
	Stats Stats
}

func (item Item) String() string {
	return fmt.Sprintf("[%v]", item.Name)
}

// Items is a collection of Item objects
type Items map[string]*Item

// LoadedItems is a map of loaded Item objects
type LoadedItems map[string]Item

// Tags is a map of an item's included tag information
type Tags map[string]bool

// LoadItems grabs all the item data from json
func LoadItems() LoadedItems {
	loaded := make(LoadedItems)
	jsonFile, err := os.Open("./character/json/items.json")
	if err != nil {
		log.Error(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var items []*Item
	json.Unmarshal(byteValue, &items)
	for _, i := range items {
		loaded[i.ID] = *i
	}

	return loaded
}

// Stat retrieves the current stat value
func (item *Item) Stat(statID string) int {
	if stat, ok := item.Stats[statID]; ok {
		return stat
	}

	log.Warn("cannot retrieve unknown stat '%v'", statID)
	return 0
}

// HasTag confirms whether the item has the specified tag
func (item *Item) HasTag(id string) bool {
	if _, ok := item.Tags[id]; ok {
		return true
	}
	return false
}
