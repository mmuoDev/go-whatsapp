package sender

//Sender represents method needed to implement in order to send whatsapp message
type Sender interface {
	SendText (msg, sendTo string) error
} 

type service struct {
	sender Sender
}

func NewService(s Sender) *service {
	return &service{
		sender: s,
	}
}

//SendText sends a whatsapp message
func (s *service) SendText (msg, sendTo string) error {
	return s.sender.SendText(msg, sendTo)
}
