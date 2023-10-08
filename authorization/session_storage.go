package authorization

import (
	"sync"
)

type SessionStorage struct {
	Sessions map[string]Session
	Mu       sync.RWMutex
	Size     uint32
}

func CreateSessionStorage() *SessionStorage {
	storage := &SessionStorage{
		Sessions: make(map[string]Session),
	}

	return storage
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
	storage.Mu.RLock()
	defer storage.Mu.RUnlock()

	val, ok := storage.Sessions[sessionId]

	if ok {
		copy := val

		return &copy, true
	}

	return nil, false
}

func (storage *SessionStorage) GetSessions() ([]Session, error) {
	storage.Mu.RLock()
	defer storage.Mu.RUnlock()

	sessions := make([]Session, storage.Size)

	i := 0
	for _, session := range storage.Sessions {
		sessions[i] = session
		i++
	}

	return sessions, nil
}
