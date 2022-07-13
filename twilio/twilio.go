package twilio

import (
	//"fmt"

	"fmt"
	"net/http"
	"strings"

	"github.com/mmuoDev/go-whatsapp/events"
)

// type TwilioProvider interface {
// 	CreateMessage(params *openapi.CreateMessageParams) (*openapi.ApiV2010Message, error)
// }

// type Listener interface {
// 	Text(string) bool
// 	Intent(interface{}) (interface{}, error)
// 	Attachments(interface{}) (interface{}, error) //image,video,documents
// 	Location(interface{}) (interface{}, error)
// }

type Connector struct {
	// sendTo         string
	// twilioProvider TwilioProvider
	//listener       Listener
	Payload *http.Request
}

func NewConnector(payload *http.Request) *Connector {
	return &Connector{
		Payload: payload,
	}
}

func (c *Connector) GetText() string {
	defer c.Payload.Body.Close()
	c.Payload.ParseForm()
	data := c.Payload.Form
	return data["Body"][0]
}

func (c *Connector) Text(msg string) bool {
	received := c.GetText()
	return strings.EqualFold(received, msg)
}

func (c *Connector) Location() (bool, events.Location) {
	defer c.Payload.Body.Close()
	c.Payload.ParseForm()
	data := c.Payload.Form
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

func (c *Connector) Attachments() (bool, events.Attachment) {
	defer c.Payload.Body.Close()
	c.Payload.ParseForm()
	data := c.Payload.Form
	for key, value := range c.Payload.Form {
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
