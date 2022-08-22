package menu

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMenuIsSetAsExpected(t *testing.T) {
	menus := Menu{}
	//main menu
	mainItemA := Item{Key: 12345, Title: "1. Mare's Specials 🍰"}
	mainItemB := Item{Key: 678910, Title: "2. Smoothies 🍝"}
	mainMenuItems := []Item{}
	mainMenuItems = append(mainMenuItems, mainItemA, mainItemB)
	if err := menus.Set(true, mainMenuItems, "0", "Welcome to Mare's Foodies Corner. See our menu options below:", "Kindly reply with 1 or 2"); err != nil {
		log.Fatal(err)
	}
	//sub menu
	subMenuA := Item{Key: 202056, Title: "1. Macaroni Special @ $150 😉"}
	subMenuB := Item{Key: 112267, Title: "2. Village Rice @ $75 🍚"}
	subMenuC := Item{Key: 346711, Title: "3. Fried Plantain and Beans @ $200"}
	subMenuItems := []Item{}
	subMenuItems = append(subMenuItems, subMenuA, subMenuB, subMenuC)
	header := "Kindly pick from our menu list 🛒"
	footer := "Pick 1,2,3 or 0 to return to the previous menu"
	err := menus.Set(false, subMenuItems, "12345", header, footer)

	t.Run("Menu is set as expected", func(t *testing.T) {
		assert.Equal(t, nil, err)
	})
}
