package sessions

import (
	"github.com/pkg/errors"
)

type SessManager struct {
	connector SessionManager
}

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

func (s *SessManager) SessionExists(sessionId, collectionName string) (bool, error) {
	return s.connector.SessionExists(sessionId, collectionName)
}

func (s *SessManager) RetrieveData(sessionId, collectionName string, result interface{}) {
	s.connector.RetrieveData(sessionId, collectionName, result)
}

func (s *SessManager) UpdateSession(sessionId, collectionName string, document map[string]interface{}) error {
	d := Storage{SessionID: sessionId, Data: document}
	return s.connector.UpdateSession(sessionId, collectionName, d)
}

func (s *SessManager) EndSession(sessionId, collectionName string) error {
	return s.connector.EndSession(sessionId, collectionName)
}

func (s *SessManager) StartSession(sessionId, collectionName string, data map[string]interface{}) error {
	c, err := s.connector.SessionExists(sessionId, collectionName)
	if err != nil {
		return errors.Wrapf(err, "sessions - unable to validate sessionId=%s", sessionId)
	}
	if c {
		return nil
	}
	d := Storage{SessionID: sessionId, Data: data}
	return s.connector.StartSession(sessionId, collectionName, d)
}
