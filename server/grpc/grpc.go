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

func UnpackItem(item *Item) *character.Item {
	r := character.CreateItem(item.ID, item.Name, item.Desc)
	r.Tags = item.Tags
	r.Stats = make(map[string]int, len(item.Stats))
	for k, v := range item.Stats {
		r.Stats[k] = int(v)
	}

	return r
}
