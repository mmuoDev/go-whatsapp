package twilio

import (
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

//Connector wraps connection to Twilio
type Connector struct {
	twilioProvider TwilioProvider
	sendFrom       string
}

type TwilioProvider interface {
	CreateMessage(params *openapi.CreateMessageParams) (*openapi.ApiV2010Message, error)
}

//NewConnection initiates a new connector to twilio
func NewConnector(sendFrom string, provider TwilioProvider) *Connector {
	return &Connector{
		twilioProvider: provider,
		sendFrom:       sendFrom,
	}
}

//sendText sends a message
func (c *Connector) SendText(msg, sendTo string) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo(sendTo)
	params.SetFrom(c.sendFrom)
	params.SetBody(msg)
	_, err := c.twilioProvider.CreateMessage(params)
	if err != nil {
		return err
	}
	return nil
}
