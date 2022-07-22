package navigation

import "github.com/mmuoDev/go-whatsapp/menu"

// Adding navigation allows you move between menus.
// Whilst users can go from "menu A" to "menu B", want them to be able to move back too.
// For sub-menus, we should have a "0. Previous Menu" that allows users return to previous menu when they enter "0".
// We don’t need "0. Previous Menu" at the "main menu" (welcome screen)
// Resource: https://chatbotslife.com/adding-navigation-menu-to-your-whatsapp-chatbot-34dac0343980

var states []string

//Navigator represents methods to be implemented to navigate to main and sub-menus
type Navigator interface {
	MainMenu() (string, error) 
	SubMenu(key string) (string, error)
}

type Navigate struct{
	menu menu.Menu
}

func NewNavigator(m menu.Menu) *Navigate {
	return &Navigate{
		menu: m,
	}
}

//MainMenu navigates to the main menu
func (n *Navigate) MainMenu() (string, error) {
	states = []string{}
	return n.menu.String(menu.PARENT)
}

//SubMenu navigates to a sub-menu
func (n *Navigate) SubMenu(key string) (string, error)  {
	if len(states) > 1 {
		count := len(states)
		previousIndex := count - 2
		previousState := states[previousIndex]
		if previousState != "0" {
			states = states[:len(states)-1] //remove current state
			states = append(states, key)  //set previous state to current 
			return n.menu.String(key)
			
		}else{
			states = []string{}
			stateId := "0"
			states = append(states, stateId)
			return n.MainMenu()
		}
	}
	return n.MainMenu()
}
