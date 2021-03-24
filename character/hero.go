package character

// The Hero is the main player character.
type Hero struct {
	Character
}

func NewHero(id string, name string, desc string) *Hero {
	h := Hero{NewCharacter(id, name, desc)}
	h.setDefaultStats()
	h.Items = []*Item{}
	h.Equipment = Equipment{}

	return &h
}

// SetDefaultStats initializes the default stats for the hero
func (hero *Hero) setDefaultStats() {
	hero.hp = 100
	hero.Stats = Stats{
		"hp-max":   100,
		"atk":      5,
		"def":      5,
		"eva":      0,
		"agi":      5,
		"lvl":      1,
		"exp":      0,
		"exp-next": 10,
	}
}

// GainExp increases the exp to Hero.
// The hero may level up
func (hero *Hero) GainExp(exp int) {
	totalExp := hero.Stat("exp") + exp
	if totalExp >= hero.Stat("exp-next") {
		hero.AddStat("lvl", totalExp/hero.Stat("exp-next"))
		hero.AddStat("exp-next", hero.Stat("lvl")*3)
		hero.SetStat("exp", totalExp%hero.Stat("exp-next"))
	} else {
		hero.AddStat("exp", exp)
	}
}
