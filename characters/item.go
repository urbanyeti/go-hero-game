package characters

import (
	"fmt"
	"strings"
)

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
type Equipment map[string]*Item

func (equipment Equipment) String() string {
	var sb strings.Builder
	for _, item := range equipment {
		sb.WriteString(item.String())
	}
	return fmt.Sprintf("Equipment: {%v}", sb.String())
}
