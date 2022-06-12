package messaging

type MessageProvider interface {
	SendText (message interface{}) error
	//SendMedia (media interface{}) error
} 

type service struct {
	messageProvider MessageProvider
}

func NewService(m MessageProvider) *service {
	return &service{
		messageProvider: m,
	}
}

func (s *service) SendText (message interface{}) error {
	return nil 
}

// func (s *service) SendMedia (media interface{}) error {
// 	return nil 
// }