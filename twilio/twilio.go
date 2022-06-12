package twilio

import (
	"fmt"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioProvider interface {
	CreateMessage(params *openapi.CreateMessageParams) (*openapi.ApiV2010Message, error)
}

type Connector struct {
	sendTo         string
	twilioProvider TwilioProvider
}

func NewConnector(sendTo string, tp TwilioProvider) *Connector {
	return &Connector{
		sendTo:         sendTo,
		twilioProvider: tp,
	}
}

func (c *Connector) SendText(message interface{}) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo(c.sendTo)
	params.SetFrom("+14155238886")
	params.SetBody(fmt.Sprint(message))
	_, err := c.twilioProvider.CreateMessage(params)
	if err != nil {
		return err
	}
	return nil
}
