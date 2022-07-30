package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	//"strconv"
	"sync"

	"github.com/mmuoDev/go-whatsapp/listen"
	"github.com/mmuoDev/go-whatsapp/menu"
	"github.com/mmuoDev/go-whatsapp/mongo"
	"github.com/mmuoDev/go-whatsapp/navigation"
	"github.com/mmuoDev/go-whatsapp/sender"
	"github.com/mmuoDev/go-whatsapp/sessions"

	//"github.com/mmuoDev/go-whatsapp/sessions"
	"github.com/mmuoDev/go-whatsapp/twilio"
	twilioGo "github.com/twilio/twilio-go"
)

//TODO: Map conversations to request types

const (
	CHATBOT_NUMBER      = "whatsapp:+14155238886"
	USER_PHONE          = "+2348067170799"
	CONVERSATION_TYPE   = "conversation_type"
	REQUEST_TYPE        = "request_type"
	COLLECTION_NAME     = "sessions"
	STATES              = "states"
	CURRENT_STATE_INDEX = "current_state_index"
)

var onlyOnce sync.Once
var menus menu.Menu

//sendText sends a WhatsApp message
func sendText(text string) {
	client := twilioGo.NewRestClientWithParams(twilioGo.ClientParams{
		Username: "ACb5cdd3fae18b869e67abdd16722619b6",
		Password: "5264c8adcad1aa6fc45838c63066cdb6",
	})
	twilioConnector := twilio.NewConnector(CHATBOT_NUMBER, client.Api)
	sender := sender.NewService(twilioConnector)
	//TODO: User's phoneNumber should be taken from the request payload
	err := sender.SendText(text, fmt.Sprintf("whatsapp:%s", USER_PHONE))
	if err != nil {
		log.Fatal(err)
	}
}

// func getMainMenu(menus menu.Menu) string {
// 	m, err := menus.String(menu.PARENT)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return m
// }

func setMenus() menu.Menu {
	//Main menu
	mainItemA := menu.Item{Key: 12345, Title: "1. Mare's Specials 🍰"}
	mainItemB := menu.Item{Key: 678910, Title: "2. Smoothies 🍝"}
	mainMenuItems := []menu.Item{}
	mainMenuItems = append(mainMenuItems, mainItemA, mainItemB)
	//Sub menu 
	subMenuA := menu.Item{Key: 202056, Title: "1. Macaroni Special @ $150 😉"}
	subMenuB := menu.Item{Key: 112267, Title: "2. Village Rice @ $75 🍚"}
	subMenuC := menu.Item{Key: 346711, Title: "3. Fried Plantain and Beans @ $200"}
	subMenuItems := []menu.Item{}
	subMenuItems = append(subMenuItems, subMenuA, subMenuB, subMenuC)
	//specialMenu := map[int]string{202056: "1. Macaroni Special @ $150 😉", 112267: "2. Village Rice @ $75 🍚", 346711: "3. Fried Plantain and Beans @ $200"}
	menus := menu.Menu{}
	if err := menus.Set(true, mainMenuItems, "0", "Welcome to Mare's Foodies Corner. See our menu options below:", "Kindly reply with 1 or 2"); err != nil {
		log.Fatal(err)
	}
	if err := menus.Set(false, subMenuItems, "12345", "Kindly pick from our menu list 🛒", "Pick 1,2,3 or 0 to return to the previous menu"); err != nil {
		log.Fatal(err)
	}
	return menus
}

type MongoConfig struct {
	DBURI  string
	DBName string
}

func retrieveSessionData(sessionMgr *sessions.SessManager) sessions.SessionData {
	s := sessions.Storage{}
	sessionMgr.RetrieveData(USER_PHONE, &s)
	return s.Data
}

func updateSessionData(sessionMgr *sessions.SessManager, data sessions.SessionData) {
	if err := sessionMgr.UpdateSession(USER_PHONE, data); err != nil {
		log.Fatal()
	}
}

func endSession(sessionMgr *sessions.SessManager) {
	if err := sessionMgr.EndSession(USER_PHONE); err != nil {
		log.Fatal()
	}
}

func BotHandler(w http.ResponseWriter, r *http.Request) {
	//Conversation types
	//1. welcome screen
	//2. mare's special
	//Request types
	//1. Select category
	//2. Pick from menu
	//3. Request for quantity

	onlyOnce.Do(func() {
		menus = setMenus()
	})

	//navigation
	navigate := navigation.NewNavigator(menus)
	//mongo
	cfg := &MongoConfig{
		DBURI:  "mongodb://localhost:27017",
		DBName: "whatsapp",
	}
	mongoConnector, err := mongo.NewConnector(cfg.DBURI, cfg.DBName)
	if err != nil {
		log.Fatal(err)
	}
	sessionMgr := sessions.NewSessManager(mongoConnector)

	//session
	sessionData := make(map[string]interface{})
	// sess := sessions.Inmemory{
	// 	Memory: make(map[string]sessions.Data),
	// }

	//listening
	data := twilio.NewListener(r)
	listener := listen.NewListener(r, data)

	// Set menus
	// mainMenu := map[int]string{12345: "Mare's Specials 🍰", 678910: "Smoothies 🍝"}
	// specialMenu := map[int]string{202056: "Macaroni Special @ $150 😉", 112267: "Village Rice @ $75 🍚", 346711: "Fried Plantain and Beans @ $200"}
	// menus := menu.Menu{}
	// if err := menus.Set(true, "0", mainMenu); err != nil {
	// 	log.Fatal(err)
	// }
	// if err := menus.Set(false, "12345", specialMenu); err != nil {
	// 	log.Fatal(err)
	// }

	//Set navigation
	//navigate := navigation.Navigate{Menu: menus}
	//Check for other conversations if not 'hello'
	//Maybe user has started a session
	c, err := sessionMgr.SessionExists(USER_PHONE)
	if err != nil {
		log.Fatal(err)
	}

	check := listener.Text("hello")
	if check && !c {
		// m, err := menus.String(menu.PARENT)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		m, err := navigate.NextMenu(menu.PARENT)
		if err != nil {
			log.Fatal(err)
		}
		//Add a header and footer to main menu
		// h := "Welcome to Mare's Foodies Corner. See our menu options below:"
		// f := "Kindly reply with 1 or 2"
		//Start session
		sessionData[CONVERSATION_TYPE] = "1"
		sessionData[REQUEST_TYPE] = "1"
		//sess.Memory[USER_PHONE] = sessionData
		if err := sessionMgr.StartSession(USER_PHONE, sessionData); err != nil {
			log.Fatal(err)
		}
		sendText(m)
		return
	}

	// s := sessions.Storage{}
	// sessionMgr.RetrieveData(USER_PHONE, collectionName, &s)
	// dd := s.Data
	// dd["book"] = "things fall apart"
	// upErr := sessionMgr.UpdateSession(USER_PHONE, collectionName, dd)
	// log.Fatal(c, upErr)
	if c {
		sessData := retrieveSessionData(sessionMgr)
		ct := sessData[CONVERSATION_TYPE]
		rt := sessData[REQUEST_TYPE]
		res := listener.GetText()
		switch ct {
		case "1":
			switch rt {
			case "1":
				//res := listener.GetText()
				if res == "1" {
					//show mare's special menu
					//navigate
					//m, err := navigate.SubMenu("12345")

					// m, err := menus.String("12345")
					// if err != nil {
					// 	log.Fatal(err)
					// }
					m, err := navigate.NextMenu("12345")
					if err != nil {
						log.Fatal(err)
					}
					states := "12345"
					// h := "Kindly pick from our menu list 🛒"
					// f := "Pick 1,2,3 or 0 to return to the previous menu"
					//reply := fmt.Sprintf("%s\n\n%s\n\n%s", h, m, f)
					data := retrieveSessionData(sessionMgr)
					data[CONVERSATION_TYPE] = "2"
					data[REQUEST_TYPE] = "2"
					data[STATES] = states
					data[CURRENT_STATE_INDEX] = 0
					updateSessionData(sessionMgr, data)
					sendText(m)
					return
				}
				if res == "2" {
					text := "Our Smoothies menu will be available soon. Please check back."
					sendText(text)
				}
			}
		case "2":
			switch rt {
			case "2":
				//check if previous menu
				if res == "0" {
					// currentState := []string{}
					// state, ok := sessData[STATES]
					// if ok {
					// 	currentState = state.([]interface{})
					// }

					previousMenu, currentIndex, currentState := navigate.PreviousMenu(sessData[STATES].(string), sessData[CURRENT_STATE_INDEX].(int32))
					m, err := navigate.NextMenu(previousMenu)
					if err != nil {
						log.Fatal(err)
					}
					//TODO: check previous menu and set set session appropriately
					//else session misbehaves
					sessData[CURRENT_STATE_INDEX] = currentIndex
					sessData[STATES] = currentState
					updateSessionData(sessionMgr, sessData)
					sendText(m)
				}else{
					var price int
					var item string
					if res == "1" {
						price = 150
						item = "Village Rice"
					} else if res == "2" {
						price = 75
						item = "Fried Plantain and Beans"
					} else if res == "3" {
						price = 200
						item = "Macaroni Special"
					}
					data := retrieveSessionData(sessionMgr)
					data["item"] = item
					data["price"] = price
					data[REQUEST_TYPE] = "3"
					updateSessionData(sessionMgr, data)
					text := fmt.Sprintf("How many plates of *%s* do you wish to have? 😍", item)
					sendText(text)
				}
				
			case "3":
				if qty, err := strconv.Atoi(res); err == nil {
					data := retrieveSessionData(sessionMgr)
					data["quantity"] = qty
					updateSessionData(sessionMgr, data)
					//show summary
					price := data["price"]
					p := price.(int32)
					item := data["item"]
					h := "See your order summary below 👜"
					f := "Thank you for your order! Someone would reach out. 🤙"
					total := p * int32(qty)
					m := fmt.Sprintf("Item: %s\nPrice:$%d\nQuantity:%d\nTotal:$%d", item, price, qty, total)
					summary := fmt.Sprintf("%s\n\n%s\n\n%s", h, m, f)
					endSession(sessionMgr)
					sendText(summary)
					return
				}
				text := "You have entered an invalid value for quantity"
				sendText(text)
			}
		}
	} else {
		h := "Ooops...I don't understand what you mean. 🙆"
		m := "See our menu options below:"
		f := "Kindly reply with 1 or 2"
		//mStr := getMainMenu(menus)
		mStr, err := navigate.NextMenu(menu.PARENT)
		if err != nil {
			log.Fatal(err)
		}
		reply := fmt.Sprintf("%s\n\n%s\n%s\n%s", h, m, mStr, f)
		//Start session
		sessionData[CONVERSATION_TYPE] = "1"
		sessionData[REQUEST_TYPE] = "1"
		//sess.Memory[USER_PHONE] = sessionData
		if err := sessionMgr.StartSession(USER_PHONE, sessionData); err != nil {
			log.Fatal(err)
		}
		sendText(reply)
	}

}

func main() {
	http.HandleFunc("/webhook", BotHandler)
	http.ListenAndServe(":8080", nil)
}
