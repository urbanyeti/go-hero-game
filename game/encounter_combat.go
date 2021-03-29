package game

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/urbanyeti/go-hero-game/character"
	"github.com/urbanyeti/go-hero-game/math"
)

// A CombatEncounter consists of a fight with a monster
type CombatEncounter struct {
	Monsters []character.Monster
}

type Fighter interface {
	ID() string
	Name() string
	HP() int
	Def() int
	Agi() int
	Eva() int
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
	Agi      int
	Spd      int
}

// Start the fight
func (encounter CombatEncounter) Start(game *Game) bool {
	monsters := getEncounter(game, encounter)
	log.WithFields(log.Fields{"hero": game.Hero, "monsters": monsters}).Info("combat started")
	keys := make([]string, 0, len(monsters))
	for key := range monsters {
		keys = append(keys, key)
	}

	target, keys := keys[0], keys[1:]
	damager := game.Hero.SelectDamager()
	targetM := monsters[target]
	attacks := map[string]*Attack{game.Hero.ID(): {game.Hero, damager, &targetM, target, 0, damager.Stat("spd")}}

	for k, m := range monsters {
		damager = m.SelectDamager()
		attacks[k] = &Attack{&m, damager, game.Hero, game.Hero.ID(), 0, damager.Stat("spd")}
	}

	for {
		for _, a := range attacks {
			a.Agi += a.Attacker.Agi()
			if a.Agi >= a.Spd {
				// Execute attack
				r := rand.Float64()
				if r < (float64(a.Target.Eva()) / 100) {
					// Target evades attack
					log.WithFields(log.Fields{"attack": a}).Info("attack evaded")
					continue
				}

				a.dealDamage()
				if a.Target.HP() <= 0 {
					if _, isHero := a.Target.(*character.Hero); isHero {
						log.WithFields(log.Fields{"hero": game.Hero}).Info("hero died")
						return true
					}

					log.WithFields(log.Fields{"hero": game.Hero.ID(), "monster": a.Target.ID()}).Info("monster slain")
					game.Hero.GainExp(a.Target.Stat("exp"))
					game.Hero.AddItem(game.Items.GetRandomItem())
					delete(monsters, a.TargetID)
					if len(keys) > 0 {
						a.TargetID, keys = keys[0], keys[1:]
						targetM := monsters[target]
						a.Target = &targetM
						continue
					}

					log.WithFields(log.Fields{"hero": game.Hero.ID()}).Info("combat finished")
					return false
				}
				damager = a.Attacker.SelectDamager()
				a.Agi = 0
				a.Damager = damager
				a.Spd = damager.Stat("spd")
			}
		}
		time.Sleep(messageDelay * time.Millisecond)
	}
}

func getEncounter(game *Game, encounter CombatEncounter) map[string]character.Monster {
	monsters := map[string]character.Monster{}

	// Get random assortment of monsters that equal hero's level
	cr := game.Hero.Stat("lvl")
	for cr > 0 {
		m := encounter.getRandomMonster(cr).Clone()
		m.AddStat("lvl", (game.Loop-1)*1)
		cr -= m.Stat("lvl")
		monsters[fmt.Sprint(m.ID(), len(monsters))] = m
	}

	return monsters
}

// Get a random level-appropriate monster
func (e CombatEncounter) getRandomMonster(maxLvl int) *character.Monster {

	ml := []character.Monster{}
	for _, c := range e.Monsters {
		if c.Stat("lvl") <= maxLvl {
			ml = append(ml, c)
		}
	}
	return &ml[rand.Intn(len(ml))]
}

func (a Attack) dealDamage() {
	baseDmg := a.Attacker.Stat("atk") + a.Attacker.Stat("lvl")
	rollDmg := rand.Intn(a.Damager.Stat("dmg-min")+a.Damager.Stat("dmg-max")) + (a.Damager.Stat("dmg-min"))
	defenderDef := a.Target.Def()
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
