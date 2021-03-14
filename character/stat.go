package character

import (
	"fmt"
	"strings"
)

// A Stat is a number assoicated with a Hero.
// Stats can be increased, decreased, and checked.
type Stat struct {
	Name        string
	Description string
	Value       int
}

func (stat Stat) String() string {
	return fmt.Sprintf("(%v: %v)", stat.Name, stat.Value)
}

// Stats is a collection of Stat objects
type Stats map[string]int

func (stats Stats) String() string {
	var sb strings.Builder
	for key, stat := range stats {
		sb.WriteString(fmt.Sprintf("(%v: %v)", key, stat))
	}
	return fmt.Sprintf("Stats: {%v}", sb.String())
}
