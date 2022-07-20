package navigation

import (
	"errors"
	"fmt"
)

//Add comments
//https://chatbotslife.com/adding-navigation-menu-to-your-whatsapp-chatbot-34dac0343980

//MenuItems represents a type for items on a menu.
type menuItems map[int]string

type menuCollection map[string]menuItems

//menus represents a collection of all menus for a chatbot
var menus menuCollection

//Menufier defines methods to set, retrieve and show a menu
//'parent' specifies if it is the main menu. Note that there can only be one main menu.
type Menufier interface {
	Set(parent bool, key int, items menuItems) error
	Get(key string) (menuItems, error)
	String(key string) (string, error)
}

const (
	PARENT = "parent"
)

type Menu struct {
}

func (m *Menu) Set(parent bool, key int, items menuItems) error {
	if len(menus) == 0 {
		menus = make(menuCollection)
	}
	if len(items) == 0 {
		return errors.New("No menu items found")
	}
	if parent == false && key == 0 {
		return errors.New("You must either set a main menu or provide menu key")
	}
	if parent == true {
		menus[PARENT] = items
		return nil
	}
	menus[fmt.Sprint(key)] = items
	return nil
}

func (m *Menu) Get(ID string) (menuItems, error) {
	if items, ok := menus[ID]; ok {
		return items, nil
	}
	return menuItems{}, fmt.Errorf("No menu items found for specified key=%s", ID)
}

func (m *Menu) String(key string) (string, error) {
	menuStr := ""
	items, err := m.Get(key)
	if err != nil {
		return menuStr, err
	}
	i := 0
	for _, v := range items {
		i++
		menuStr += fmt.Sprintf("%d. %s \n", i, v)
	}
	if key != PARENT {
		menuStr += "0. Previous Menu"
	}
	return menuStr, nil
}

func RetrieveMap() menuCollection {
	return menus
}
