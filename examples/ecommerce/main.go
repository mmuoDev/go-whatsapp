package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/mmuoDev/go-whatsapp/listen"
	"github.com/mmuoDev/go-whatsapp/menu"
	"github.com/mmuoDev/go-whatsapp/mongo"
	"github.com/mmuoDev/go-whatsapp/navigation"
	"github.com/mmuoDev/go-whatsapp/sender"
	"github.com/mmuoDev/go-whatsapp/sessions"
	"github.com/mmuoDev/go-whatsapp/twilio"
	twilioGo "github.com/twilio/twilio-go"
)

const (
	CHATBOT_NUMBER      = "whatsapp:+14155238886"
	CONVERSATION_TYPE   = "conversation_type"
	REQUEST_TYPE        = "request_type"
	COLLECTION_NAME     = "sessions"
	STATES              = "states"
	CURRENT_STATE_INDEX = "current_state_index"
)

var onlyOnce sync.Once
var menus menu.Menu

//MongoConfig defines configs for mongo
type MongoConfig struct {
	DBURI  string
	DBName string
}

//sendText sends a WhatsApp message using Twilio
func sendText(text, phoneNumber string) {
	client := twilioGo.NewRestClientWithParams(twilioGo.ClientParams{
		Username: "",
		Password: "",
	})
	twilioConnector := twilio.NewConnector(CHATBOT_NUMBER, client.Api)
	sender := sender.NewService(twilioConnector)
	err := sender.SendText(text, phoneNumber)
	if err != nil {
		log.Fatal(err)
	}
}

//setMenus sets the menu for the chatbot
func setMenus() menu.Menu {
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

//retrieveSessionData retrieves data from the session
func retrieveSessionData(sessionMgr *sessions.SessManager, phoneNumber string) sessions.SessionData {
	s := sessions.Storage{}
	sessionMgr.RetrieveData(phoneNumber, &s)
	return s.Data
}

//updateSessionData updates data on a session
func updateSessionData(sessionMgr *sessions.SessManager, data sessions.SessionData, phoneNumber string) {
	if err := sessionMgr.UpdateSession(phoneNumber, data); err != nil {
		log.Fatal()
	}
}

//endSession ends a session
func endSession(sessionMgr *sessions.SessManager, phoneNumber string) {
	if err := sessionMgr.EndSession(phoneNumber); err != nil {
		log.Fatal()
	}
}

//BotHandler defines webhook to be used on Twilio
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

	navigate := navigation.NewNavigator(menus)

	cfg := &MongoConfig{
		DBURI:  "mongodb://localhost:27017",
		DBName: "whatsapp",
	}
	mongoConnector, err := mongo.NewConnector(cfg.DBURI, cfg.DBName)
	if err != nil {
		log.Fatal(err)
	}
	sessionMgr := sessions.NewSessManager(mongoConnector)
	sessionData := make(map[string]interface{})
	
	data := twilio.NewListener(r)
	listener := listen.NewListener(r, data)
	phoneNumber := listener.GetPhoneNumber()

	c, err := sessionMgr.SessionExists(phoneNumber)
	if err != nil {
		log.Fatal(err)
	}

	check := listener.Text("hello")
	if check && !c {
		m, err := navigate.NextMenu(menu.PARENT)
		if err != nil {
			log.Fatal(err)
		}
		sessionData[CONVERSATION_TYPE] = "1"
		sessionData[REQUEST_TYPE] = "1"
		if err := sessionMgr.StartSession(phoneNumber, sessionData); err != nil {
			log.Fatal(err)
		}
		sendText(m, phoneNumber)
		return
	}

	if c {
		sessData := retrieveSessionData(sessionMgr, phoneNumber)
		ct := sessData[CONVERSATION_TYPE]
		rt := sessData[REQUEST_TYPE]
		res := listener.GetText()
		switch ct {
		case "1":
			switch rt {
			case "1":
				if res == "1" {
					m, err := navigate.NextMenu("12345")
					if err != nil {
						log.Fatal(err)
					}
					states := "12345"
					data := retrieveSessionData(sessionMgr, phoneNumber)
					data[CONVERSATION_TYPE] = "2"
					data[REQUEST_TYPE] = "2"
					data[STATES] = states
					data[CURRENT_STATE_INDEX] = 0
					updateSessionData(sessionMgr, data, phoneNumber)
					sendText(m, phoneNumber)
					return
				}
				if res == "2" {
					text := "Our Smoothies menu will be available soon. Please check back."
					sendText(text, phoneNumber)
				}
			}
		case "2":
			switch rt {
			case "2":
				//check if it's previous menu
				if res == "0" {
					previousMenu, currentIndex, currentState := navigate.PreviousMenu(sessData[STATES].(string), sessData[CURRENT_STATE_INDEX].(int32))
					m, err := navigate.NextMenu(previousMenu)
					if err != nil {
						log.Fatal(err)
					}
					//TODO:: check previous menu and set set session appropriately
					//else session misbehaves
					sessData[CURRENT_STATE_INDEX] = currentIndex
					sessData[STATES] = currentState
					updateSessionData(sessionMgr, sessData, phoneNumber)
					sendText(m, phoneNumber)
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
					data := retrieveSessionData(sessionMgr, phoneNumber)
					data["item"] = item
					data["price"] = price
					data[REQUEST_TYPE] = "3"
					updateSessionData(sessionMgr, data, phoneNumber)
					text := fmt.Sprintf("How many plates of *%s* do you wish to have? 😍", item)
					sendText(text, phoneNumber)
				}
			case "3":
				if qty, err := strconv.Atoi(res); err == nil {
					data := retrieveSessionData(sessionMgr, phoneNumber)
					data["quantity"] = qty
					updateSessionData(sessionMgr, data, phoneNumber)
					price := data["price"]
					p := price.(int32)
					item := data["item"]
					h := "See your order summary below 👜"
					f := "Thank you for your order! Someone would reach out. 🤙"
					total := p * int32(qty)
					m := fmt.Sprintf("Item: %s\nPrice:$%d\nQuantity:%d\nTotal:$%d", item, price, qty, total)
					summary := fmt.Sprintf("%s\n\n%s\n\n%s", h, m, f)
					endSession(sessionMgr, phoneNumber)
					sendText(summary, phoneNumber)
					return
				}
				text := "You have entered an invalid value for quantity"
				sendText(text, phoneNumber)
			}
		}
	} else {
		h := "Ooops...I don't understand what you mean. 🙆"
		m := "See our menu options below:"
		f := "Kindly reply with 1 or 2"
		mStr, err := navigate.NextMenu(menu.PARENT)
		if err != nil {
			log.Fatal(err)
		}
		reply := fmt.Sprintf("%s\n\n%s\n%s\n%s", h, m, mStr, f)
		sessionData[CONVERSATION_TYPE] = "1"
		sessionData[REQUEST_TYPE] = "1"
		if err := sessionMgr.StartSession(phoneNumber, sessionData); err != nil {
			log.Fatal(err)
		}
		sendText(reply, phoneNumber)
	}
}

func main() {
	http.HandleFunc("/webhook", BotHandler)
	http.ListenAndServe(":8080", nil)
}
