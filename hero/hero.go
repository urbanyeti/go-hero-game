package hero

import (
	"fmt"
	"strings"
)

// The Hero is the main player character.
type Hero struct {
	ID          string
	Name        string
	Description string
	HP          Stat
	Stats       Stats
	Equipment   Equipment
}

func (hero Hero) String() string {
	return fmt.Sprintf("%v - %v | %v | %v", hero.Name, hero.HP, hero.Stats, hero.Equipment)
}

// An Item is collected and may be used or equipped by the Hero.
type Item struct {
	ID          string
	Name        string
	Description string
}

func (item Item) String() string {
	return fmt.Sprintf("[%v]", item.Name)
}

// Equipment is a collection of Item objects
type Equipment map[string]Item

func (equipment Equipment) String() string {
	var sb strings.Builder
	for _, item := range equipment {
		sb.WriteString(item.String())
	}
	return fmt.Sprintf("Equipment: {%v}", sb.String())
}

// A Stat is a number assoicated with a Hero.
// Stats can be increased, decreased, and checked.
type Stat struct {
	ID          string
	Name        string
	Description string
	Value       int
}

func (stat Stat) String() string {
	return fmt.Sprintf("(%v: %v)", stat.Name, stat.Value)
}

// Stats is a collection of Stat objects
type Stats map[string]Stat

func (stats Stats) String() string {
	var sb strings.Builder
	for _, stat := range stats {
		sb.WriteString(stat.String())
	}
	return fmt.Sprintf("Stats: {%v}", sb.String())
}

// Hello returns a hello message for the provided name.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
