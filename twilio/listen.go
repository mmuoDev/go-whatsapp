package twilio

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mmuoDev/go-whatsapp/whatsapp"
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

func (l *Listener) Location() (bool, whatsapp.Location) {
	defer l.Payload.Body.Close()
	l.Payload.ParseForm()
	data := l.Payload.Form
	lats, ok := data["Latitude"]
	if !ok {
		return false, whatsapp.Location{}
	}
	lons, ok := data["Longitude"]
	if !ok {
		return false, whatsapp.Location{}
	}
	return true, whatsapp.Location{
		Longitude: lons[0],
		Latitude:  lats[0],
	}

}

func (l *Listener) Attachments() (bool, whatsapp.Attachment) {
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
		return true, whatsapp.Attachment{
			MediaURL:  mediaUrl,
			MediaName: mediaName,
			MediaType: mediaType,
		}
	}

	return false, whatsapp.Attachment{}
}

