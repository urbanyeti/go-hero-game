package grpc

import "github.com/urbanyeti/go-hero-game/character"

func PackItem(item *character.Item) *Item {
	r := &Item{}
	r.ID = item.ID()
	r.Name = item.Name()
	r.Desc = item.Desc()
	r.Tags = item.Tags
	r.Stats = make(map[string]int32, len(item.Stats))
	for k, v := range item.Stats {
		r.Stats[k] = int32(v)
	}

	return r
}

func PackItems(items []*character.Item) []*Item {
	l := make([]*Item, len(items))
	for _, v := range items {
		l = append(l, PackItem(v))
	}

	return l
}

func UnpackItem(item *Item) *character.Item {
	r := character.CreateItem(item.ID, item.Name, item.Desc)
	r.Tags = item.Tags
	r.Stats = make(map[string]int, len(item.Stats))
	for k, v := range item.Stats {
		r.Stats[k] = int(v)
	}

	return r
}

func PackMonster(monster *character.Monster) *Monster {
	r := &Monster{
		ID:        monster.ID(),
		Name:      monster.Name(),
		Desc:      monster.Desc(),
		Tags:      monster.Tags,
		HP:        int32(monster.HP()),
		Stats:     make(map[string]int32, len(monster.Stats)),
		Abilities: PackAbilities(monster.Abilities),
		Items:     PackItems(monster.Items),
		Equipment: PackEquipment(monster.Equipment),
	}

	for k, v := range monster.Stats {
		r.Stats[k] = int32(v)
	}

	return r
}

func PackEquipment(equipment character.Equipment) map[string]*Item {
	m := make(map[string]*Item, len(equipment))
	for k, v := range equipment {
		m[k] = PackItem(v)
	}

	return m
}

func PackAbilities(abilities character.Abilities) map[string]*Ability {
	m := make(map[string]*Ability, len(abilities))
	for _, a := range abilities {
		r := &Ability{
			ID:    a.ID(),
			Name:  a.Name(),
			Desc:  a.Desc(),
			Tags:  a.Tags,
			Stats: make(map[string]int32, len(a.Stats)),
		}
		for k, v := range a.Stats {
			r.Stats[k] = int32(v)
		}
		m[r.ID] = r
	}

	return m
}
