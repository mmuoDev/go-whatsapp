package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mmuoDev/go-whatsapp/messaging"
	"github.com/mmuoDev/go-whatsapp/twilio"
	twilioGo "github.com/twilio/twilio-go"
)

func BotHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	client := twilioGo.NewRestClientWithParams(twilioGo.ClientParams{
		Username: "ACb5cdd3fae18b869e67abdd16722619b6",
		Password: "5264c8adcad1aa6fc45838c63066cdb6",
	})
	sendTo := "+2348067170799"
	twilioConnector := twilio.NewConnector(sendTo, client.Api)
	serviceMgr := messaging.NewService(twilioConnector)
	if 2 == 3 {
		err := serviceMgr.SendText("hello there")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	
	http.HandleFunc("/webhook", BotHandler)
	http.ListenAndServe(":8080", nil)
}
