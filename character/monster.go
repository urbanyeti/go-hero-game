package character

import "github.com/urbanyeti/go-hero-game/graphics"

// Monster is an agent in the world
type Monster struct {
	Character
	Animations *graphics.Animations
}

func (c CharacterJSON) LoadMonster(la map[string]Ability, li map[string]*Item, animations *graphics.Animations) Monster {
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

	tags := make(Tags)
	for _, v := range c.Tags {
		tags[v] = true
	}

	monster := Monster{Character{
		id:        c.ID,
		name:      c.Name,
		desc:      c.Desc,
		Tags:      tags,
		hp:        c.HP,
		Stats:     c.Stats,
		Items:     items,
		Abilities: abilities,
		Equipment: Equipment{},
	}, animations}

	monster.Equip(items...)

	return monster
}

func (m Monster) Clone() Monster {
	n := m
	n.Stats = make(Stats, len(m.Stats))
	for k, v := range m.Stats {
		n.Stats[k] = v
	}
	n.Abilities = make(Abilities, len(m.Abilities))
	for k, v := range m.Abilities {
		n.Abilities[k] = v
	}
	n.Equipment = make(Equipment, len(m.Equipment))
	for k, v := range m.Equipment {
		n.Equipment[k] = v
	}
	anims := graphics.Animations{}
	for k, v := range *m.Animations {
		anims[k] = v
	}
	n.Animations = &anims

	n.Items = make([]*Item, len(m.Items))
	copy(n.Items, m.Items)

	return n
}

func (m *Monster) GetAnimations() *graphics.Animations {
	return m.Animations
}
