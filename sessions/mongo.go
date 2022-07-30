package sessions

import (
	"github.com/pkg/errors"
)

type SessManager struct {
	connector SessionManager
}

const (
	COLLECTION_NAME = "sessions"
)

type SessionData map[string]interface{}

type Storage struct {
	SessionID string      `bson:"sessionId"`
	Data      SessionData `bson:"data"`
}

func NewSessManager(manager SessionManager) *SessManager {
	return &SessManager{
		connector: manager,
	}
}

func (s *SessManager) SessionExists(sessionId string) (bool, error) {
	return s.connector.SessionExists(sessionId)
}

func (s *SessManager) RetrieveData(sessionId string, result interface{}) {
	s.connector.RetrieveData(sessionId, result)
}

func (s *SessManager) UpdateSession(sessionId string, document map[string]interface{}) error {
	d := Storage{SessionID: sessionId, Data: document}
	return s.connector.UpdateSession(sessionId, d)
}

func (s *SessManager) EndSession(sessionId string) error {
	return s.connector.EndSession(sessionId)
}

func (s *SessManager) StartSession(sessionId string, data map[string]interface{}) error {
	c, err := s.connector.SessionExists(sessionId)
	if err != nil {
		return errors.Wrapf(err, "sessions - unable to validate sessionId=%s", sessionId)
	}
	if c {
		return nil
	}
	d := Storage{SessionID: sessionId, Data: data}
	return s.connector.StartSession(sessionId, d)
}
