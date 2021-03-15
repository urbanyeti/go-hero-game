package character

// Monster is an agent in the world
type Monster struct {
	Character
}

func (c CharacterJSON) LoadMonster(la map[string]Ability, li map[string]*Item) Monster {
	// Add values of loaded abilities to monster
	abilities := make(Abilities)
	for _, a := range c.Abilities {
		value := la[a]
		abilities[a] = &value
	}

	// Add values of loaded items to monster
	items := []*Item{}
	for _, i := range c.Items {
		value := li[i]
		items = append(items, value)
	}

	monster := Monster{Character{
		id:        c.ID,
		name:      c.Name,
		desc:      c.Desc,
		hp:        c.HP,
		stats:     c.Stats,
		items:     items,
		abilities: abilities,
		equipment: Equipment{},
	}}

	monster.Equip(items...)

	return monster
}

func (m Monster) Clone() Monster {
	n := m

	n.stats = make(Stats, len(m.stats))
	for k, v := range m.stats {
		n.stats[k] = v
	}

	n.abilities = make(Abilities, len(m.abilities))
	for k, v := range m.abilities {
		n.abilities[k] = v
	}

	n.equipment = make(Equipment, len(m.equipment))
	for k, v := range m.equipment {
		n.equipment[k] = v
	}

	n.items = make([]*Item, len(m.items))
	copy(n.items, m.items)

	return n
}
