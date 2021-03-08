package main

import (
	"fmt"
	"time"

	"github.com/urbanyeti/go-hero-game/hero"
)

type game struct {
	hero *hero.Hero
	loop int
	turn int
}

func (game game) String() string {
	return fmt.Sprintf("Loop: %v Turn: %v | %v", game.loop, game.turn, game.hero)
}

func main() {
	hero := hero.Hero{
		ID:          "hero-dan",
		Name:        "Dan",
		Description: "a cool hero",
		HP: hero.Stat{
			ID:          "stat-hp",
			Name:        "HP",
			Description: "health points",
			Value:       100},
		Stats: hero.Stats{
			"stat-atk": hero.Stat{
				ID:          "stat-atk",
				Name:        "ATK",
				Description: "attack stat",
				Value:       5},
			"stat-def": hero.Stat{
				ID:          "stat-def",
				Name:        "DEF",
				Description: "defense stat",
				Value:       5},
		},
		Equipment: hero.Equipment{
			"item-sword1": hero.Item{
				ID:          "item-sword1",
				Name:        "Short Sword",
				Description: "a beginner's basic short sword"},
			"item-armor1": hero.Item{
				ID:          "item-armor1",
				Name:        "Basic Armor",
				Description: "a beginner's basic set of armor"},
		}}

	game := game{hero: &hero, loop: 1, turn: 1}
	fmt.Println(game)
	time.Sleep((1000))
}
