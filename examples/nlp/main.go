package main

import (
	"log"
	"net/http"

	dialogflow "github.com/mmuoDev/go-whatsapp/dialogflow"
	"github.com/mmuoDev/go-whatsapp/listen"
	"github.com/mmuoDev/go-whatsapp/nlp"
	"github.com/mmuoDev/go-whatsapp/sender"
	"github.com/mmuoDev/go-whatsapp/twilio"
	twilioGo "github.com/twilio/twilio-go"
)

const (
	CHATBOT_NUMBER = "whatsapp:+14155238886"
)

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

//BotHandler defines webhook to be used on Twilio
func BotHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := dialogflow.NewConnector("", "default", "xxxxx-creds.json", "europe-west2")
	if err != nil {
		log.Fatal(err)
	}
	nlpService := nlp.NewService(conn)
	data := twilio.NewListener(r)
	listener := listen.NewListener(r, data)
	res := listener.GetText()
	phoneNumber := listener.GetPhoneNumber()

	events, err := nlpService.DetectIntent(res, "en")
	if err != nil {
		log.Fatal(err)
	}
	if events.IntentType == "Default Welcome Intent" {
		sendText("Welcome to my chatbot! 😀", phoneNumber)
	} else {
		sendText("Oops..Something went wrong!", phoneNumber)
	}
}

func main() {

	http.HandleFunc("/webhook", BotHandler)
	http.ListenAndServe(":8080", nil)
}
