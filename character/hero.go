package character

// The Hero is the main player character.
type Hero struct {
	Character
}

func NewHero(id string, name string, desc string) *Hero {
	h := Hero{NewCharacter(id, name, desc)}
	h.setDefaultStats()
	h.items = make(Items)

	return &h
}

// SetDefaultStats initializes the default stats for the hero
func (hero *Hero) setDefaultStats() {
	hero.hp = 100
	hero.stats = Stats{
		"hp-max":   100,
		"atk":      5,
		"def":      5,
		"eva":      0,
		"spd":      5,
		"lvl":      1,
		"exp":      0,
		"exp-next": 10,
	}
}

// GainExp increases the exp to Hero.
// The hero may level up
func (hero *Hero) GainExp(exp int) {
	totalExp := hero.Stat("exp") + exp
	//fmt.Print("    ")
	if totalExp >= hero.Stat("exp-next") {
		hero.AddStat("lvl", totalExp/hero.Stat("exp-next"))
		hero.SetStat("exp", totalExp%hero.Stat("exp-next"))
		//fmt.Printf(" | Level %v (%v/%v)", hero.Stat("lvl"), hero.Stat("exp"), hero.Stat("exp-next"))
	} else {
		hero.AddStat("exp", exp)
		//fmt.Printf(" | Level %v (%v/%v)", hero.Stat("lvl"), hero.Stat("exp"), hero.Stat("exp-next"))
	}
	//fmt.Print("\n")
}
