package webhook

type Webhook interface {
	//Handler func(http.ResponseWriter, *http.Request)
}

// type Payload interface {
// 	GetRequest() (*http.Request, error)
// }

type Hook struct {
	//handler Webhook
	//webh ook Webhook
}

// func (h *Handler) ServeHTTP (w http.ResponseWriter, r *http.Request) (*http.Request){
// 	return nil
// }

// func NewWebh(addr string) Hook {
// 	return Hook{

// 	}
// }

// func (h *Handler) Handler (addr string, handler http.Handler) error {
// 	return
// }

// func NewCommunicator (webhook Webhook) *Communicator {
// 	return &Communicator{
// 		webhook: webhook,
// 	}
// }

// func (c *Communicator) Payload() error {
// 	r := c.webhook.Handler()
// }

// func NewHandler(webhook Webhook) *Handler {
// 	return &Handler{
// 		webhook: webhook,
// 	}
// }

// func (h Handler) Payload (*http.Request, error) {
// 	return h.webhook.ServeHTTP(), nil
// }
