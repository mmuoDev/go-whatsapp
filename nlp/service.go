package nlp

import "github.com/mmuoDev/go-whatsapp/events"

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

func (s *service) DetectIntent(text, languageCode string) (events.DetectIntent, error) {
	return s.intentDetector.DetectIntentText(text, languageCode)
}
