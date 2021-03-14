package character

// Monster is an agent in the world
type Monster struct {
	Character
}

func (c CharacterJSON) LoadMonster() Monster {
	return Monster{Character{
		id:    c.ID,
		name:  c.Name,
		desc:  c.Desc,
		hp:    c.HP,
		stats: c.Stats,
		items: c.Items,
	}}
}
