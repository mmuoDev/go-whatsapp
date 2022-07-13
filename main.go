package main

import (
	"log"
	"net/http"

	dialogflow "github.com/mmuoDev/go-whatsapp/dialogflow"
	"github.com/mmuoDev/go-whatsapp/listen"
	"github.com/mmuoDev/go-whatsapp/twilio"
)

func BotHandler(w http.ResponseWriter, r *http.Request) {
	// client := twilioGo.NewRestClientWithParams(twilioGo.ClientParams{
	// 	Username: "ACb5cdd3fae18b869e67abdd16722619b6",
	// 	Password: "5264c8adcad1aa6fc45838c63066cdb6",
	// })
	// sendTo := "+2348067170799"
	//twilioConnector := twilio.NewConnector(sendTo, client.Api, r)

	//listen
	data := twilio.NewConnector(r)
	serviceMgr := listen.NewListener(r, data)
	//check := serviceMgr.Text("hello")
	check, _ := serviceMgr.Location()
	log.Println("here", check)

	// serviceMgr := messaging.NewService(twilioConnector)
	// if 2 == 3 {
	// 	err := serviceMgr.SendText("hello there")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	//Listen for events

	//Respond to events

}

func main() {

	// http.HandleFunc("/webhook", BotHandler)

	// http.ListenAndServe(":8080", nil)

	conn, err := dialogflow.NewConnector("weatherapp-aodc", "default", "dialogflow-creds.json", "europe-west2")
	if err != nil {
		log.Fatal(err)
	}
	text, err := conn.DetectIntentText("what is the weather in jos", "en")
	if err != nil {
		log.Fatal("error", err)
	}
	log.Fatal(text)
}
