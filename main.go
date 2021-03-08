package main

import (
	"fmt"

	"github.com/urbanyeti/go-hero-game/hero"
)

func main() {
	message := hero.Hello("James")
	fmt.Printf(message)
}
