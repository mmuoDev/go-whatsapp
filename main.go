package main

import (
	"log"
	"net/http"

	//dialogflow "github.com/mmuoDev/go-whatsapp/dialogflow"
	//"github.com/mmuoDev/go-whatsapp/listen"
	//"github.com/mmuoDev/go-whatsapp/nlp"
	"github.com/mmuoDev/go-whatsapp/sender"
	"github.com/mmuoDev/go-whatsapp/twilio"
	twilioGo "github.com/twilio/twilio-go"
)

func BotHandler(w http.ResponseWriter, r *http.Request) {
	//send message
	client := twilioGo.NewRestClientWithParams(twilioGo.ClientParams{
		Username: "ACb5cdd3fae18b869e67abdd16722619b6",
		Password: "5264c8adcad1aa6fc45838c63066cdb6",
	})
	twilioConnector := twilio.NewConnector("whatsapp:+14155238886", client.Api)
	sender := sender.NewService(twilioConnector)
	err := sender.SendText("boo, bitch!", "whatsapp:+2348067170799")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal("DONE")

	//listen
	// data := twilio.NewConnector(r)
	// serviceMgr := listen.NewListener(r, data)
	// check, _ := serviceMgr.Location()
	// log.Println("here", check)

	//detect intent 
	// conn, err := dialogflow.NewConnector("weatherapp-aodc", "default", "dialogflow-creds.json", "europe-west2")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// nlpService := nlp.NewService(conn)
	// events, err := nlpService.DetectIntent("what is the weather in jos", "en")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Fatal("res", events)
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

	http.HandleFunc("/webhook", BotHandler)

	http.ListenAndServe(":8080", nil)

	// conn, err := dialogflow.NewConnector("weatherapp-aodc", "default", "dialogflow-creds.json", "europe-west2")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// text, err := conn.DetectIntentText("what is the weather in jos", "en")
	// if err != nil {
	// 	log.Fatal("error", err)
	// }
	// log.Fatal(text)
}
