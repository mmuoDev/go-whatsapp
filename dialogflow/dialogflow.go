package dialogflow

import (
	"context"
	"errors"
	"fmt"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"github.com/mmuoDev/go-whatsapp/events"
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

type Connector struct {
	projectID string
	sessionID string
	client    *dialogflow.SessionsClient
	ctx       context.Context
	region    string
}

func NewConnector(projectID, sessionID, credsPath, region string) (*Connector, error) {
	ctx := context.Background()
	endpoint := ""
	if region != "" {
		endpoint = fmt.Sprintf("%s-dialogflow.googleapis.com:443", region)
	}
	sessionClient, err := dialogflow.NewSessionsClient(ctx, option.WithCredentialsFile(credsPath), option.WithEndpoint(endpoint))
	if err != nil {
		return &Connector{}, err
	}
	//defer sessionClient.Close()
	return &Connector{
		projectID: projectID,
		sessionID: sessionID,
		client:    sessionClient,
		ctx:       ctx,
		region:    region,
	}, nil
}

func (c *Connector) DetectIntentText(text, languageCode string) (events.DetectIntent, error) {
	defer c.client.Close()
	if c.projectID == "" || c.sessionID == "" {
		return events.DetectIntent{}, errors.New(fmt.Sprintf("Received empty project (%s) or session (%s)", c.projectID, c.sessionID))
	}
	sessionPath := fmt.Sprintf("projects/%s/locations/%s/agent/sessions/%s", c.projectID, c.region, c.sessionID)
	textInput := dialogflowpb.TextInput{Text: text, LanguageCode: languageCode}
	queryTextInput := dialogflowpb.QueryInput_Text{Text: &textInput}
	queryInput := dialogflowpb.QueryInput{Input: &queryTextInput}
	request := dialogflowpb.DetectIntentRequest{Session: sessionPath, QueryInput: &queryInput}

	response, err := c.client.DetectIntent(c.ctx, &request) //TODO: Interface?
	if err != nil {
		return events.DetectIntent{}, err
	}
	return events.DetectIntent{
		FulfilmentText:             response.QueryResult.FulfillmentText,
		IntentType:                 response.QueryResult.Intent.DisplayName,
		IntentDectectionConfidence: int(response.QueryResult.IntentDetectionConfidence),
	}, nil
}
