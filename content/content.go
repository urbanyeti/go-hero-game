package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/urbanyeti/go-hero-game/character"
)

var lock sync.Mutex

func main() {
	if len(os.Args) > 1 {
		contentType := os.Args[1]
		if contentType == "ability" || contentType == "abilities" || contentType == "a" {
			SaveAbilities(OpenFile())
		} else if contentType == "item" || contentType == "items" || contentType == "i" {
			SaveItems(OpenFile())
		} else if contentType == "monster" || contentType == "monsters" || contentType == "m" {
			SaveMonsters(OpenFile())
		}
	} else {
		f := OpenFile()
		SaveAbilities(f)
		SaveItems(f)
		SaveMonsters(f)
	}

	os.Exit(0)
}

func OpenFile() *excelize.File {
	f, err := excelize.OpenFile("./content.xlsx")
	ExitIfError(err)
	return f
}

func SaveAbilities(f *excelize.File) {
	log.Info("saving abilities")
	abilities := []*character.AbilityJSON{}
	abilities = append(abilities, LoadAbilities(f, "attacks")...)
	err := Save("../character/json/abilities.json", abilities)
	ExitIfError(err)
	log.Info("abilities saved")
}

func SaveItems(f *excelize.File) {
	log.Info("saving items")
	items := []*character.ItemJSON{}
	items = append(items, LoadItems(f, "weapons")...)
	items = append(items, LoadItems(f, "armor")...)
	err := Save("../character/json/items.json", items)
	ExitIfError(err)
	log.Info("items saved")
}

func SaveMonsters(f *excelize.File) {
	log.Info("saving monsters")
	monsters := []*character.CharacterJSON{}
	monsters = append(monsters, LoadCharacters(f, "monsters")...)
	err := Save("../character/json/monsters.json", monsters)
	ExitIfError(err)
	log.Info("monsters saved")
}

func LoadAbilities(f *excelize.File, sheet string) []*character.AbilityJSON {
	log.WithFields(logrus.Fields{"file": f.Path, "sheet": sheet}).Info("importing sheet of abilities")
	rows, err := f.GetRows(sheet)
	ExitIfError(err)

	abilities := make([]*character.AbilityJSON, len(rows)-1)
	for ri, row := range rows {
		if ri == 0 {
			continue
		}

		ability := character.AbilityJSON{
			ID:    row[0],
			Name:  row[1],
			Desc:  row[2],
			Tags:  SplitString(row[3]),
			Stats: make(character.Stats, 3),
		}

		for i := 4; i < 7; i++ {
			header := rows[0][i]
			if i >= len(row) || row[i] == "" {
				continue
			}
			var val, err = strconv.Atoi(row[i])
			if err != nil {
				log.WithFields(log.Fields{"row": ri, "col": i}).Error("can't convert cell")
				continue
			}

			ability.Stats[header] = (int)(val)
		}

		abilities[ri-1] = &ability
	}

	return abilities
}

func LoadItems(f *excelize.File, sheet string) []*character.ItemJSON {
	log.WithFields(logrus.Fields{"file": f.Path, "sheet": sheet}).Info("importing sheet of items")
	rows, err := f.GetRows(sheet)
	ExitIfError(err)

	items := make([]*character.ItemJSON, len(rows)-1)
	for ri, row := range rows {
		if ri == 0 {
			continue
		}

		item := character.ItemJSON{
			ID:    row[0],
			Name:  row[1],
			Desc:  row[2],
			Tags:  SplitString(row[3]),
			Stats: make(character.Stats, 6),
		}

		for i := 4; i < 10; i++ {
			header := rows[0][i]
			if i >= len(row) || row[i] == "" {
				continue
			}
			var val, err = strconv.Atoi(row[i])
			if err != nil {
				log.WithFields(log.Fields{"row": ri, "col": i}).Error("can't convert cell")
				continue
			}

			item.Stats[header] = (int)(val)
		}

		items[ri-1] = &item
	}

	return items
}

func LoadCharacters(f *excelize.File, sheet string) []*character.CharacterJSON {
	log.WithFields(logrus.Fields{"file": f.Path, "sheet": sheet}).Info("importing sheet of characters")
	rows, err := f.GetRows(sheet)
	ExitIfError(err)

	characters := make([]*character.CharacterJSON, len(rows)-1)
	for ri, row := range rows {
		if ri == 0 {
			continue
		}

		hp, err := strconv.Atoi(row[4])
		if err != nil {
			log.WithFields(log.Fields{"row": ri, "col": 4}).Error("can't convert cell")
			continue
		}
		c := character.CharacterJSON{
			ID:        row[0],
			Name:      row[1],
			Desc:      row[2],
			Tags:      SplitString(row[3]),
			HP:        int(hp),
			Abilities: SplitString(row[5]),
			Items:     SplitString(row[6]),
			Stats:     make(character.Stats, 5),
		}

		for i := 7; i < 12; i++ {
			header := rows[0][i]
			if i >= len(row) || row[i] == "" {
				continue
			}
			var val, err = strconv.Atoi(row[i])
			if err != nil {
				log.WithFields(log.Fields{"row": ri, "col": i}).Error("can't convert cell")
				continue
			}

			c.Stats[header] = int(val)
		}

		characters[ri-1] = &c
	}

	return characters
}

func SplitString(v string) []string {
	if v == "" {
		return nil
	}

	vals := strings.Split(strings.ReplaceAll(v, " ", ""), ",")

	if len(vals) == 0 {
		return nil
	}

	return vals
}

func Save(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	log.WithField("path", path).Info("creating file")
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := MarshalJSON(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	log.WithField("path", path).Info("file created")
	return err
}

func MarshalJSON(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func ExitIfError(err error) {
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}
}
