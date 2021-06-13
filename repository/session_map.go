package repository

import (
	"errors"
	"sync"
	"userCreation/domain"
)

type sessionRepository struct {
	m sync.Map
}

func NewSyncMapSessionRepository() domain.SessionRepository {
	return &sessionRepository{}
}

func (s *sessionRepository) Store(session domain.Session) error {
	s.m.Store(session.ID, session)
	return nil
}

func (s *sessionRepository) Delete(id string) error {
	s.m.Delete(id)
	return nil
}

func (s *sessionRepository) Load(id string) (domain.Session, error) {
	if value, ok := s.m.Load(id); ok {
		return value.(domain.Session), nil
	}
	return domain.Session{}, errors.New("user not found")
}
