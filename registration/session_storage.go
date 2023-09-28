package registration

import (
	"sync"
)

type SessionStorage struct {
	Sessions map[string]Session
	Mu       sync.Mutex
	Size     uint32
}

func (storage *SessionStorage) RegisterNewSession(session Session) error {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	storage.Size++
	storage.Sessions[session.SessionId] = session

	return nil
}

func (storage *SessionStorage) DeleteSession(sessionId string) error {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	if _, ok := storage.Sessions[sessionId]; !ok {
		return ErrNoSuchSession
	}

	delete(storage.Sessions, sessionId)
	storage.Size--

	return nil
}

func (storage *SessionStorage) CheckSession(sessionId string) (*Session, bool) {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	val, ok := storage.Sessions[sessionId]

	if ok {
		copy := val

		return &copy, true
	}

	return nil, false
}

func (storage *SessionStorage) GetSessions() ([]Session, error) {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	sessions := make([]Session, storage.Size)

	for _, session := range storage.Sessions {
		sessions = append(sessions, session)
	}

	return sessions, nil
}
