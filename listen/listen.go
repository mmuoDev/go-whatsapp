package listen

import (
	"net/http"

	"github.com/mmuoDev/go-whatsapp/whatsapp"
)

type Listener interface {
	Text(string) bool
	Attachments() (bool, whatsapp.Attachment) //image,video,documents,contacts
	Location() (bool, whatsapp.Location)
	GetText() string
	GetPhoneNumber() string
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

func (l *Listen) GetPhoneNumber() string {
	return l.listen.GetPhoneNumber()
}

func (l *Listen) Text(msg string) bool {
	return l.listen.Text(msg)
}

func (l *Listen) Attachments() (bool, whatsapp.Attachment) {
	return l.listen.Attachments()
}

func (l *Listen) Location() (bool, whatsapp.Location) {
	return l.listen.Location()
}
