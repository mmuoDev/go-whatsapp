package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mmuoDev/go-whatsapp/listen"
	"github.com/mmuoDev/go-whatsapp/menu"
	"github.com/mmuoDev/go-whatsapp/sender"
	"github.com/mmuoDev/go-whatsapp/sessions"
	"github.com/mmuoDev/go-whatsapp/twilio"
	twilioGo "github.com/twilio/twilio-go"
)

//TODO: Map conversations to request types

const (
	CHATBOT_NUMBER    = "whatsapp:+14155238886"
	USER_PHONE        = "+2348067170799"
	CONVERSATION_TYPE = "conversation_type"
	REQUEST_TYPE      = "request_type"
)

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

func getMainMenu(menus menu.Menu) string {
	m, err := menus.String(menu.PARENT)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func BotHandler(w http.ResponseWriter, r *http.Request) {
	//Conversation types
	//1. welcome screen
	//2. mare's special
	//Request types
	//1. Select category
	//2. Pick from menu
	//3. Request for quantity

	//session
	sessionData := make(map[string]interface{})
	sess := sessions.Inmemory{
		Memory: make(map[string]sessions.Data),
	}

	//listening
	data := twilio.NewListener(r)
	listener := listen.NewListener(r, data)

	// Set menus
	mainMenu := map[int]string{12345: "Mare's Specials 🍰", 678910: "Smoothies 🍝"}
	specialMenu := map[int]string{202056: "Macaroni Special @ $150 😉", 112267: "Village Rice @ $75 🍚", 346711: "Fried Plantain and Beans @ $200"}
	menus := menu.Menu{}
	if err := menus.Set(true, "0", mainMenu); err != nil {
		log.Fatal(err)
	}
	if err := menus.Set(false, "12345", specialMenu); err != nil {
		log.Fatal(err)
	}

	//Set navigation
	//navigate := navigation.Navigate{Menu: menus}

	check := listener.Text("hello")
	if check {
		m, err := menus.String(menu.PARENT)
		if err != nil {
			log.Fatal(err)
		}
		//Add a header and footer to main menu
		h := "Welcome to Mare's Foodies Corner. See our menu options below:"
		f := "Kindly reply with 1 or 2"
		reply := fmt.Sprintf("%s\n\n%s\n%s", h, m, f)

		//Start session
		sessionData[CONVERSATION_TYPE] = "1"
		sessionData[REQUEST_TYPE] = "1"
		sess.Memory[USER_PHONE] = sessionData
		sendText(reply)
	}

	//Check for other conversations if not 'hello'
	//Maybe user has started a session
	if sess.SessionExists(USER_PHONE) {
		ct := sess.RetrieveData(USER_PHONE, CONVERSATION_TYPE)
		rt := sess.RetrieveData(USER_PHONE, REQUEST_TYPE)
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

					m, err := menus.String("12345")
					if err != nil {
						log.Fatal(err)
					}
					h := "Kindly pick from our menu list"
					f := "Pick 1,2,3 or 0 to return to the previous menu"
					reply := fmt.Sprintf("%s\n\n%s\n%s", h, m, f)

					sess.UpdateSession(USER_PHONE, CONVERSATION_TYPE, "2")
					sess.UpdateSession(USER_PHONE, REQUEST_TYPE, "2")
					sendText(reply)
				}
				if res == "2" {
					text := "Our Smoothies menu will be available soon. Please check back."
					sendText(text)
				}
			}
		case "2":
			switch rt {
			case "2":
				var price int
				var item string
				if res == "1" {
					price = 150
					item = "Macaroni Special"
				} else if res == "2" {
					price = 75
					item = "Village Rice"
				} else if res == "3" {
					price = 200
					item = "Fried Plantain and Beans"
				}
				sess.UpdateSession(USER_PHONE, "item", item)
				sess.UpdateSession(USER_PHONE, "price", price)
				sess.UpdateSession(USER_PHONE, REQUEST_TYPE, "3")
				text := "How many plates do you wish to have? 😍"
				sendText(text)
			case "3":
				if qty, err := strconv.Atoi(res); err == nil {
					sess.UpdateSession(USER_PHONE, "quantity", qty)
					//show summary
					price := sess.RetrieveData(USER_PHONE, "price")
					p := price.(int)
					item := sess.RetrieveData(USER_PHONE, "item")
					h := "See your order summary below"
					f := "Thank you for your order! Someone would reach out. 😋"
					total := p * qty
					m := fmt.Sprintf("Item: %s\nPrice:%s\nQuantity:%d\n\nTotal:%d", item, price, qty, total)
					summary := fmt.Sprintf("%s\n%s\n\n%s", h, m, f)
					sendText(summary)
					sess.EndSession(USER_PHONE)
				}
				text := "You have entered an invalid value for quantity"
				sendText(text)
			}
		}
	}
	h := "Ooops...I don't understand what you mean. 🙆"
	m := "See our menu options below:"
	f := "Kindly reply with 1 or 2"
	mStr := getMainMenu(menus)
	reply := fmt.Sprintf("%s\n\n%s\n%s\n%s", h, m, mStr, f)
	sendText(reply)
}

func main() {
	http.HandleFunc("/webhook", BotHandler)
	http.ListenAndServe(":5723", nil)
}
