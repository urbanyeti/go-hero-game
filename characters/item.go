package characters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// An Item is collected and may be used or equipped by the Hero.
type Item struct {
	ID          string
	Name        string
	Description string
	Stats       Stats
}

func (item Item) String() string {
	return fmt.Sprintf("[%v]", item.Name)
}

// Items is a collection of Item objects
type Items map[string]*Item

// LoadItems grabs all the item data from json
func LoadItems() Items {
	jsonFile, err := os.Open("./characters/json/items.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var items Items
	json.Unmarshal(byteValue, &items)
	return items
}

// Stat retrieves the current stat value
func (item *Item) Stat(statID string) int {
	if stat, ok := item.Stats[statID]; ok {
		return stat
	}

	log.Printf("cannot retrieve unknown stat '%v'", statID)
	return 0
}

// Equipment is a collection of equipped Item objects
type Equipment map[string]*Item

// func (equipment Equipment) String() string {
// 	var sb strings.Builder
// 	for _, item := range equipment {
// 		sb.WriteString(item.String())
// 	}
// 	return fmt.Sprintf("Equipment: {%v}", sb.String())
// }
