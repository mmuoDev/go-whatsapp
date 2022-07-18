package nlp

import "github.com/mmuoDev/go-whatsapp/events"

//IntentDetector represents method needed to implement to detect text from an nlp e.g. dialogflow
type IntentDetector interface {
	DetectIntentText(text, languageCode string) (events.DetectIntent, error)
}

type service struct {
	intentDetector IntentDetector
}

func NewService(intent IntentDetector) *service {
	return &service{
		intentDetector: intent,
	}
}

//DetectIntent detects intect from a text
func (s *service) DetectIntent(text, languageCode string) (events.DetectIntent, error) {
	return s.intentDetector.DetectIntentText(text, languageCode)
}
