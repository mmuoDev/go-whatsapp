/**
Adding navigation allows you move amongst menus.
Whilst users can go from "menu A" to "menu B", want them to be able to move back too.
For sub-menus, we should have a "0. Previous Menu" that allows users return to previous menu when they enter "0".
We don’t need "0. Previous Menu" at the "main menu" (welcome screen)
Resource: https://chatbotslife.com/adding-navigation-menu-to-your-whatsapp-chatbot-34dac0343980
**/
package navigation

import (
	"strings"

	"github.com/mmuoDev/go-whatsapp/menu"
)

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

//PreviousMenu returns the previous menu
func (n *Navigate) PreviousMenu(states string, currentState int32) (string, int, string) {
	ss := strings.Split(states, ",")
	if len(ss) == 0 {
		return menu.PARENT, 0, ""
	}
	if len(ss) > 1 {
		newIndex := currentState - 1
		previousMenu := ss[newIndex]
		newState := ss[:newIndex]
		return previousMenu, int(newIndex), strings.Join(newState, ",")
	}
	return menu.PARENT, 0, states
}

//NextMenu returns a menu
func (n *Navigate) NextMenu(key string) (string, error) {
	return n.Menu.String(key)
}
