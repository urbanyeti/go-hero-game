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
		if contentType == "item" || contentType == "items" || contentType == "i" {
			SaveItems()
		}
	} else {
		SaveItems()
	}

	os.Exit(0)
}

func SaveItems() {
	log.Info("saving items")
	f, err := excelize.OpenFile("./items.xlsx")
	ExitIfError(err)
	items := []*character.ItemJSON{}
	items = append(items, LoadItems(f, "weapons")...)
	items = append(items, LoadItems(f, "armor")...)
	err = Save("../character/json/items.json", items)
	log.Info("items saved")
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
			Tags:  strings.Split(row[3], ","),
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

			item.Stats[header] = val
		}

		items[ri-1] = &item
	}

	return items
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
