package storage

import (
	"encoding/json"
	"io/ioutil"

	"github.com/brobb923/packingList/packing"
)

func Load(filename string) error {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var savedList []*packing.Item
	err = json.Unmarshal(fileContents, &savedList)
	if err != nil {
		return err
	}

	packing.SetLists(savedList)

	return nil
}

func Save(filename string) error {
	itemList := packing.ListItems()

	itemListBytes, err := json.MarshalIndent(itemList, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, itemListBytes, 0775)
	if err != nil {
		return err
	}

	return nil
}
