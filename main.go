package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/brobb923/packingList/packing"
	"github.com/brobb923/packingList/storage"
	"github.com/manifoldco/promptui"
)

func main() {

	namePrompt := promptui.Prompt{
		Label: "Name",
	}
	name, err := namePrompt.Run()
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println(name)

	storage.Load(filename string)

	if len(packing.ListItems()) == 0 {
		err = promptCategories(packing.Categories)
		if err != nil {

			fmt.Println(err)

		}
	}

	packingList := packing.FinalList()
	for {
		itemNames := []string{}

		for _, item := range packingList {
			listEntry := item.Name

			if item.BeenPacked {
				listEntry = listEntry + "[x]"

			} else {
				listEntry = listEntry + "[]"

			}

			itemNames = append(itemNames, listEntry)
		}

		storage.Save(filename string)

		prompt := promptui.Select{
			Label: "Packing List",
			Items: itemNames,
		}
		index, result, err := prompt.Run()
		if err != nil {
			fmt.Println("Prompt failed", err)
			return
		}

		time.Sleep(200 * time.Millisecond)
		fmt.Println("You packed", result)
		time.Sleep(1500 * time.Millisecond)
		packingList[index].BeenPacked = true

	}

}

func promptCategories(categories []packing.Category) error {

	for _, category := range categories {
		fmt.Println(category.ListName)
		err := promptCategory(category)
		if err != nil {

			fmt.Println(err)

		}
	}

	return nil
}

func promptCategory(category packing.Category) error {
	for _, suggestion := range category.Suggestions {
		// use numberPromptHelper to get number of items for this suggestion

		amount, err := numberPromptHelper(suggestion)
		if err != nil {
			fmt.Println(err)
		}

		packing.AddToPackingList(suggestion, amount)
	}
	return nil
}

func numberPromptHelper(label string) (int, error) {
	validate := func(userInput string) error {
		_, err := strconv.Atoi(userInput)
		if err != nil {
			return errors.New("You screwed up!")
		}
		return nil
	}

	numberPrompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}
	numberStr, err := numberPrompt.Run()
	if err != nil {
		return 0, err
	}
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, err
	}

	return number, nil

}
