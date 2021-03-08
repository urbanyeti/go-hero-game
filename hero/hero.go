package hero

import (
	"fmt"
)

// Hello returns a hello message for the provided name.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
