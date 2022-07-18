package twilio

import (

	"fmt"
	"net/http"
	"strings"

	"github.com/mmuoDev/go-whatsapp/events"
)


type Listener struct {
	Payload *http.Request
}

func NewListener(payload *http.Request) *Listener {
	return &Listener{
		Payload: payload,
	}
}

func (l *Listener) GetText() string {
	defer l.Payload.Body.Close()
	l.Payload.ParseForm()
	data := l.Payload.Form
	return data["Body"][0]
}

func (l *Listener) Text(msg string) bool {
	received := l.GetText()
	return strings.EqualFold(received, msg)
}

func (l *Listener) Location() (bool, events.Location) {
	defer l.Payload.Body.Close()
	l.Payload.ParseForm()
	data := l.Payload.Form
	lats, ok := data["Latitude"]
	if !ok {
		return false, events.Location{}
	}
	lons, ok := data["Longitude"]
	if !ok {
		return false, events.Location{}
	}
	return true, events.Location{
		Longitude: lons[0],
		Latitude:  lats[0],
	}

}

func (l *Listener) Attachments() (bool, events.Attachment) {
	defer l.Payload.Body.Close()
	l.Payload.ParseForm()
	data := l.Payload.Form
	for key, value := range l.Payload.Form {
		fmt.Printf("Key:%s, Value:%s\n", key, value)
	}
	numMedia := data["NumMedia"][0]
	if numMedia != "0" {
		mediaUrl := data["MediaUrl0"][0]
		mediaType := data["MediaContentType0"][0]
		mediaName := data["Body"][0]
		return true, events.Attachment{
			MediaURL:  mediaUrl,
			MediaName: mediaName,
			MediaType: mediaType,
		}
	}

	return false, events.Attachment{}
}

// func (c *Connector) SendText(message interface{}) error {
// 	params := &openapi.CreateMessageParams{}
// 	params.SetTo(c.sendTo)
// 	params.SetFrom("+14155238886")
// 	params.SetBody(fmt.Sprint(message))
// 	_, err := c.twilioProvider.CreateMessage(params)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// Key:MessageSid, Value:[SMf107adff086f69ad677bc216f008c51c]
// Key:From, Value:[whatsapp:+2348067170799]
// Key:NumSegments, Value:[1]
// Key:WaId, Value:[2348067170799]
// Key:AccountSid, Value:[ACb5cdd3fae18b869e67abdd16722619b6]
// Key:To, Value:[whatsapp:+14155238886]
// Key:NumMedia, Value:[0]
// Key:SmsStatus, Value:[received]
// Key:Body, Value:[HI]
// Key:SmsSid, Value:[SMf107adff086f69ad677bc216f008c51c]
// Key:SmsMessageSid, Value:[SMf107adff086f69ad677bc216f008c51c]
// Key:ProfileName, Value:[Uche]
// Key:ApiVersion, Value:[2010-04-01]
// Key:ReferralNumMedia, Value:[0]
