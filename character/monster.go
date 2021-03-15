package character

// Monster is an agent in the world
type Monster struct {
	Character
}

func (c CharacterJSON) LoadMonster(la map[string]Ability, li map[string]Item) Monster {

	abilities := make(Abilities)
	for _, a := range c.Abilities {
		value := la[a]
		abilities[a] = &value
	}

	items := make(Items)
	for _, i := range c.Items {
		value := li[i]
		items[i] = &value
	}

	return Monster{Character{
		id:        c.ID,
		name:      c.Name,
		desc:      c.Desc,
		hp:        c.HP,
		stats:     c.Stats,
		items:     items,
		abilities: abilities,
	}}
}
