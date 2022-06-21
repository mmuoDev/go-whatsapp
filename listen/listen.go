package listen

import (
	"net/http"

	"github.com/mmuoDev/go-whatsapp/events"
)

type Listener interface {
	Text(string) bool
	// Intent(interface{}) (interface{}, error)
	Attachments() (bool, events.Attachment) //image,video,documents,contacts
	Location() (bool, events.Location)
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

func (l *Listen) Text(msg string) bool {
	return l.listen.Text(msg)
}

// func (l *Listen) Intent(interface{}) (interface{}, error) {
// 	return nil, nil
// }

func (l *Listen) Attachments() (bool, events.Attachment) {
	return l.listen.Attachments()
}

func (l *Listen) Location() (bool, events.Location) {
	return l.listen.Location()
}
