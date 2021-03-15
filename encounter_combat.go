package main

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/urbanyeti/go-hero-game/character"
	"github.com/urbanyeti/go-hero-game/math"
)

// A CombatEncounter consists of a fight with a monster
type CombatEncounter struct {
	Monsters map[string]character.Monster
}

type Fighter interface {
	ID() string
	Name() string
	HP() int
	Stat(string) int
	Weapons() ([]*character.Item, bool)
	TakeDamage(int)
	SelectDamager() character.Damager
}

type Attack struct {
	Attacker Fighter
	Damager  character.Damager
	Target   Fighter
	TargetID string
	Spd      int
	ExecSpd  int
}

// Start the fight
func (encounter CombatEncounter) Start(game *Game) bool {
	log.WithFields(log.Fields{"hero": game.Hero, "monsters": encounter.Monsters}).Info("combat started")
	keys := make([]string, 0, len(encounter.Monsters))
	for key := range encounter.Monsters {
		keys = append(keys, key)
	}

	target, keys := keys[0], keys[1:]
	damager := game.Hero.SelectDamager()
	targetM := encounter.Monsters[target]
	attacks := map[string]*Attack{game.Hero.ID(): {game.Hero, damager, &targetM, target, 0, damager.Stat("spd")}}

	for k, m := range encounter.Monsters {
		damager = m.SelectDamager()
		attacks[k] = &Attack{&m, damager, game.Hero, game.Hero.ID(), 0, damager.Stat("spd")}
	}

	for {
		for _, a := range attacks {
			a.Spd += a.Attacker.Stat("spd")
			if a.Spd >= a.ExecSpd {
				// Execute attack
				a.dealDamage()
				if a.Target.HP() <= 0 {
					if _, isHero := a.Target.(*character.Hero); isHero {
						log.WithFields(log.Fields{"hero": game.Hero.ID()}).Info("hero died")
						return true
					}

					log.WithFields(log.Fields{"hero": game.Hero.ID(), "monster": a.Target.ID()}).Info("monster slain")
					game.Hero.GainExp(a.Target.Stat("exp"))
					delete(encounter.Monsters, a.TargetID)
					if len(keys) > 0 {
						a.TargetID, keys = keys[0], keys[1:]
						targetM := encounter.Monsters[target]
						a.Target = &targetM
					}

					log.WithFields(log.Fields{"hero": game.Hero.ID()}).Info("combat finished")
					return false
				}
				damager = a.Attacker.SelectDamager()
				a.Spd = 0
				a.Damager = damager
				a.ExecSpd = damager.Stat("spd")
			}
		}
		time.Sleep(messageDelay * time.Millisecond)
	}
}

func (a Attack) dealDamage() {
	baseDmg := a.Attacker.Stat("atk") + a.Target.Stat("lvl")
	rollDmg := rand.Intn(a.Damager.Stat("dmg-min")+a.Damager.Stat("dmg-max")) + (a.Damager.Stat("dmg-min"))
	defenderDef := a.Target.Stat("def")
	totalDmg := math.MaxOf(baseDmg+rollDmg-defenderDef, 0)

	a.Target.TakeDamage(totalDmg)
	log.WithFields(
		log.Fields{
			"attack":      a,
			"baseDmg":     baseDmg,
			"rollDmg":     rollDmg,
			"defenderDef": defenderDef,
			"totalDmg":    totalDmg,
		}).Info("damage dealt")
}

func removeAttack(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
