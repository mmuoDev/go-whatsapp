package navigation

import (
	"log"
	"testing"

	"github.com/mmuoDev/go-whatsapp/menu"
	"github.com/stretchr/testify/assert"
)

func sampleMenu() menu.Menu {
	menus := menu.Menu{}
	//main menu
	mainItemA := menu.Item{Key: 12345, Title: "1. Mare's Specials 🍰"}
	mainItemB := menu.Item{Key: 678910, Title: "2. Smoothies 🍝"}
	mainMenuItems := []menu.Item{}
	mainMenuItems = append(mainMenuItems, mainItemA, mainItemB)
	if err := menus.Set(true, mainMenuItems, "0", "Welcome to Mare's Foodies Corner. See our menu options below:", "Kindly reply with 1 or 2"); err != nil {
		log.Fatal(err)
	}
	//sub menu
	subMenuA := menu.Item{Key: 202056, Title: "1. Macaroni Special @ $150 😉"}
	subMenuB := menu.Item{Key: 112267, Title: "2. Village Rice @ $75 🍚"}
	subMenuC := menu.Item{Key: 346711, Title: "3. Fried Plantain and Beans @ $200"}
	subMenuItems := []menu.Item{}
	subMenuItems = append(subMenuItems, subMenuA, subMenuB, subMenuC)
	if err := menus.Set(false, subMenuItems, "12345", "Kindly pick from our menu list 🛒", "Pick 1,2,3 or 0 to return to the previous menu"); err != nil {
		log.Fatal(err)
	}
	return menus
}

func TestNavigatorMovesToTheNextMenu(t *testing.T) {
	navigator := NewNavigator(sampleMenu())
	nextMenuKey := "12345"
	expectedMenu := "Kindly pick from our menu list 🛒\n1. Macaroni Special @ $150 😉 \n2. Village Rice @ $75 🍚 \n3. Fried Plantain and Beans @ $200 \n0. Previous Menu\nPick 1,2,3 or 0 to return to the previous menu"
	nextMenu, err := navigator.NextMenu(nextMenuKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Next menu is as expected", func(t *testing.T) {
		assert.Equal(t, expectedMenu, nextMenu)
	})
}

func TestNavigatorMovesToThePreviousMenu(t *testing.T) {
	navigator := NewNavigator(sampleMenu())
	states := "parent,12345"
	currentIndex := 2
	previousMenu, newIndex, newState  := navigator.PreviousMenu(states, int32(currentIndex))

	t.Run("Previous menu key is as expected", func(t *testing.T) {
		assert.Equal(t, "12345", previousMenu)
	})
	t.Run("New menu index is as expected", func(t *testing.T) {
		assert.Equal(t, 1, newIndex)
	})
	t.Run("New menu states is as expected", func(t *testing.T) {
		assert.Equal(t, "parent", newState)
	})
}
