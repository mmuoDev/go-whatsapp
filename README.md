# **Go-WhatsApp**

Build a WhatsApp chatbot using Go. This package makes it easy to use NLP (Natural Language Processing), navigation, menus and sessions on a WhatsApp chatbot. 

You can easily plugin your WhatsApp Business API either directly from Facebook or third-party platforms like Twilio,Clickatel,etc. 

https://user-images.githubusercontent.com/16281845/181497645-f984def0-1be1-4d98-acb1-902b8b8e2a6e.MP4

This package provides an implementation for Twilio. The code snippets uses Twilio's WhatsApp.

## **Listening** 
You can listen for different types of messages.  
### Listen for Text
```go
twilio := twilio.NewListener(r)
listener := listen.NewListener(r, twilio)
text := listener.GetText()
```
### Listen for Location
```go
twilio := twilio.NewListener(r)
listener := listen.NewListener(r, twilio)
f, location := listener.Location()
```
### Listen for Attachments
```go
twilio := twilio.NewListener(r)
listener := listen.NewListener(r, twilio)
f, attachments := listener.Attachments()
```

## **Sending**

### Send Text
```go
client := twilioGo.NewRestClientWithParams(twilioGo.ClientParams{
		Username: "",
		Password: "",
})
twilioConnector := twilio.NewConnector("", client.Api)
sender := sender.NewService(twilioConnector)
if err := sender.SendText(text, phoneNumber); err != nil {
	log.Fatal(err)
}
```

## **Natural Language Processing (NLP)**
Natural language processing (NLP) is the ability of a computer program to understand human language as it is spoken and written -- referred to as natural language. 
This package provides an implementation for [DialogFlow](https://cloud.google.com/dialogflow/docs)

### Detect Intect of a WhatsApp message
```go
conn, err := dialogflow.NewConnector("", "default", "xxxxx-creds.json", "europe-west2")
if err != nil {
	log.Fatal(err)
}
nlpService := nlp.NewService(conn)
twilio := twilio.NewListener(r)
listener := listen.NewListener(r, twilio)
text := listener.GetText()
events, err := nlpService.DetectIntent(text, "en")
if err != nil {
	log.Fatal(err)
}
```

## **Sessions**
### Add data to sessions using `Mongo` as the driver
```go
cfg := &MongoConfig{
	DBURI:  "mongodb://localhost:27017",
	DBName: "test",
}
mongoConnector, err := mongo.NewConnector(cfg.DBURI, cfg.DBName)
if err != nil {
	log.Fatal(err)
}
sessionMgr := sessions.NewSessManager(mongoConnector)
sessionData := make(map[string]interface{})
sessionData["message"] = "hello world"
if err := sessionMgr.StartSession("", sessionData); err != nil {
	log.Fatal(err)
}
```

## **Menus**
### Add a menu
```go
menus := menu.Menu{}
itemA := menu.Item{Key: 12345, Title: "1. Mare's Specials 🍰"}
itemB := menu.Item{Key: 678910, Title: "2. Smoothies 🍝"}
items := []menu.Item{}
menuA = append(items, itemA, itemB)
header := "Welcome to Mare's Foodies Corner. See our menu options below:"
footer := "Kindly reply with 1 or 2"
if err := menus.Set(true, menuA, "0", header, footer); err != nil {
	log.Fatal(err)
}
```

## **Navigation**
### Navigate to a menu
```go
navigate := navigation.NewNavigator(menus)
m, err := navigate.NextMenu("menu key here")
if err != nil {
	log.Fatal(err)
}
```

## **Examples**
`/examples` provides examples on how to use this package. 

