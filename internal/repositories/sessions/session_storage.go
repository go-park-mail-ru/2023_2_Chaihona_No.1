package sessions

import (
	"sync"
)

type SessionStorage struct {
	Sessions map[string]SessionOld
	Mu       sync.RWMutex
	Size     uint32
}

func CreateSessionStorage() *SessionStorage {
	storage := &SessionStorage{
		Sessions: make(map[string]SessionOld),
		Mu:       sync.RWMutex{},
		Size:     0,
	}

	return storage
}

func (storage *SessionStorage) RegisterNewSession(session SessionOld) error {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	storage.Size++
	storage.Sessions[session.SessionID] = session

	return nil
}

func (storage *SessionStorage) DeleteSession(sessionID string) error {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	if _, ok := storage.Sessions[sessionID]; !ok {
		return ErrNoSuchSession
	}

	delete(storage.Sessions, sessionID)
	storage.Size--

	return nil
}

func (storage *SessionStorage) CheckSession(sessionId string) (*SessionOld, bool) {
	storage.Mu.RLock()
	defer storage.Mu.RUnlock()

	val, ok := storage.Sessions[sessionId]

	if ok {
		copy := val

		return &copy, true
	}

	return nil, false
}
