package menu

//The package handles menu specifications for a chatbot.
//You may want to use menus so your users can easily navigate your chatbot. Menus are usually better compared
//to making your chatbot open ended. A menu allows you present your users with set of options they can easily pick from.
//Here is an example from World Health Organization's WhatsApp Chatbot
/******
WELCOME TO HEALTH ALERT
Get information and guidance about COVID-19 and related health topics.
Reply with a number, emoji or WORD IN BOLD to explore one of these topics:
1 🔢 LATEST news & cases
2 💉 VACCINE updates
3 🌞 HEALTH topics
4 🦠️ COVID-19 info and advice
5 🔗 CONNECT and share
***/

//TODO:: Using a map distorts the order of the menu items. 

import (
	"errors"
	"fmt"
)

const (
	PARENT = "parent" //default key for "main menu". Main menu is more like your welcome screen options.
	//You can change this value provided it does not conflicts with other menu keys.
	PREVIOUS_MENU = "Previous Menu"
)

//MenuItems represents a type for items on a menu.
//For each item, you should set a unique key and  name e.g. ["924999429": "See Our FAQs"]
//Unique Keys are important in order to navigate to a submenu for the particular item.
type menuItems map[int]string

//menus represents a type for collection of all menus for a chatbot
type menuCollection map[string]menu

type Item struct {
	Key int 
	Title string
}
type menu struct {
	items []Item
	header string
	footer string
}

var menus menuCollection

//Menufier defines methods to set, retrieve and show a menu.
//'parent' specifies if it is the main menu (welcome screen). Note that there can only be one main menu.
type Menufier interface {
	Set(parent bool, items []Item, key, header, footer string) error
	Get(key string) (menuItems, error)
	String(key string) (string, error)
	GetMenuCollection() menuCollection
}

type Menu struct{}

//GetMenuCollection returns all menus for a chatbot
func (m *Menu) GetMenuCollection() menuCollection {
	return menus
}

//Set adds a menu to the chatbot
func (m *Menu) Set(parent bool, items []Item, key, header, footer string) error {
	if len(menus) == 0 {
		menus = make(menuCollection)
	}
	if len(items) == 0 {
		return errors.New("No menu items found")
	}
	if parent == false && key == "" {
		return errors.New("You must either set a main menu or provide menu key")
	}
	if len(menus) > 0 {
		if m, ok := menus[key]; ok {
			return fmt.Errorf("This key=%d already has a submenu=%v assigned to it", key, m)
		}
	}
	if parent == true {
		menus[PARENT] = menu{items: items, header: header, footer: footer}
		return nil
	}
	menus[fmt.Sprint(key)] = menu{items: items, header: header, footer: footer}
	return nil
}

//Get retrieves a menu
func (m *Menu) Get(ID string) (menu, error) {
	if menu, ok := menus[ID]; ok {
		return menu, nil
	}
	return menu{}, fmt.Errorf("No menu found for specified key=%s", ID)
}

//String returns the string representation for a menu
func (m *Menu) String(key string) (string, error) {
	menu, err := m.Get(key)
	itemStr := ""
	if err != nil {
		return itemStr, err
	}
	items := menu.items
	header := menu.header
	footer := menu.footer
	//i := 0
	for _, v := range items {
		//i++
		itemStr += fmt.Sprintf("%s \n", v.Title)
	}
	if fmt.Sprint(key) != PARENT {
		itemStr += fmt.Sprintf("0. %s", PREVIOUS_MENU) //See navigation package.
	}
	return fmt.Sprintf("%s\n%s\n%s", header, itemStr, footer), nil
}
