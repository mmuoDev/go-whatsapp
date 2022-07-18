package listen

import (
	"net/http"

	"github.com/mmuoDev/go-whatsapp/events"
)

//Listener represents methods needed to be implemented to detect different types of messages
type Listener interface {
	Text(string) bool
	Attachments() (bool, events.Attachment) //image,video,documents,contacts
	Location() (bool, events.Location)
	GetText() string 

}

type Listen struct {
	payload *http.Request
	listen  Listener
}

func NewListener(r *http.Request, listener Listener) *Listen {
	return &Listen{
		payload: r,
		listen:  listener,
	}
}

func (l *Listen) GetText() string {
	return l.listen.GetText()
}

func (l *Listen) Text(msg string) bool {
	return l.listen.Text(msg)
}

func (l *Listen) Attachments() (bool, events.Attachment) {
	return l.listen.Attachments()
}

func (l *Listen) Location() (bool, events.Location) {
	return l.listen.Location()
}
