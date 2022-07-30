package navigation

import (
	"strings"

	"github.com/mmuoDev/go-whatsapp/menu"
)

// Adding navigation allows you move between menus.
// Whilst users can go from "menu A" to "menu B", want them to be able to move back too.
// For sub-menus, we should have a "0. Previous Menu" that allows users return to previous menu when they enter "0".
// We don’t need "0. Previous Menu" at the "main menu" (welcome screen)
// Resource: https://chatbotslife.com/adding-navigation-menu-to-your-whatsapp-chatbot-34dac0343980

//Navigator represents methods to be implemented to navigate to menus
type Navigator interface {
	NextMenu(key string) (string, error)
	PreviousMenu(states string, currentState int32) (string, int, string)
}

type Navigate struct {
	Menu menu.Menu
}

func NewNavigator(m menu.Menu) *Navigate {
	return &Navigate{
		Menu: m,
	}
}

//MainMenu navigates to the main menu
// func (n *Navigate) MainMenu() (string, error) {
// 	//states = []string{}
// 	//n.session.StartSession()
// 	return n.Menu.String(menu.PARENT)
// }

func (n *Navigate) PreviousMenu(states string, currentState int32) (string, int, string) {
	ss := strings.Split(states, ",")
	if len(ss) == 0 {
		return menu.PARENT, 0, ""
	}
	if len(ss) > 1 {
		newIndex := currentState - 1
		previousMenu := ss[newIndex]
		newState := ss[:newIndex+1]
		return previousMenu, int(newIndex), strings.Join(newState, ",")
	}
	return menu.PARENT, 0, states
}

//NextMenu navigates to the next menu
func (n *Navigate) NextMenu(key string) (string, error) {
	// if len(states) > 1 {
	// 	count := len(states)
	// 	previousIndex := count - 2
	// 	previousState := states[previousIndex]
	// 	if previousState != "0" {
	// 		states = states[:len(states)-1] //remove current state
	// 		states = append(states, key)    //set previous state to current
	// 		return n.Menu.String(key)

	// 	} else {
	// 		states = []string{}
	// 		stateId := "0"
	// 		states = append(states, stateId)
	// 		return n.MainMenu()
	// 	}
	// }
	// return n.MainMenu()
	return n.Menu.String(key)
}
