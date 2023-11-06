package sessions

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRegisterNewSession(t *testing.T) {
	session := Session{
		SessionID: "1",
		UserID:    1,
		TTL:       time.Now().Add(10 * time.Hour),
	}

	storage := CreateSessionStorage()
	err := storage.RegisterNewSession(session)

	if err != nil {
		t.Fatalf("Erorr during registration new session!")
	} else {
		if val, ok := storage.Sessions[session.SessionID]; ok {
			if val != session {
				t.Fatalf("Error session mismatch data!")
			}
		} else {
			t.Fatalf("Erorr session was not registered!")
		}
	}
}

func TestRegisterNewSessionNotEmptyStorage(t *testing.T) {
	session1 := Session{
		SessionID: "1",
		UserID:    1,
		TTL:       time.Now().Add(10 * time.Hour),
	}

	session2 := Session{
		SessionID: "1",
		UserID:    1,
		TTL:       time.Now().Add(10 * time.Hour),
	}

	storage := CreateSessionStorage()
	err := storage.RegisterNewSession(session1)
	if err != nil {
		t.Fatalf("Erorr during registration new session!")
	}

	err = storage.RegisterNewSession(session2)

	if err != nil {
		t.Fatalf("Erorr during registration new session!")
	} else {
		if val, ok := storage.Sessions[session2.SessionID]; ok {
			if val != session2 {
				t.Fatalf("Error session mismatch data!")
			}
		} else {
			t.Fatalf("Erorr session was not registered!")
		}
	}
}

func TestCheckSession(t *testing.T) {
	const SESSIONS = 100

	storage := CreateSessionStorage()

	for i := 0; i < SESSIONS; i++ {
		session := Session{
			SessionID: fmt.Sprint(i),
			UserID:    uint32(i),
			TTL:       time.Now().Add(10 * time.Hour),
		}

		storage.RegisterNewSession(session)

		if val, ok := storage.CheckSession(session.SessionID); ok {
			if *val != session {
				t.Fatalf("Error session mismatch data!")
			}
		} else {
			t.Fatalf("Error registered session was not found!")
		}
	}
}

func TestRegisterNewSessionGoroutines(t *testing.T) {
	const SESSIONS = 100

	storage := CreateSessionStorage()

	var m sync.Mutex
	testsPassed := true

	for i := 0; i < SESSIONS; i++ {
		go func(i int, m *sync.Mutex) {
			session := Session{
				SessionID: fmt.Sprint(i),
				UserID:    uint32(i),
				TTL:       time.Now().Add(10 * time.Hour),
			}

			storage.RegisterNewSession(session)

			if val, ok := storage.CheckSession(session.SessionID); ok {
				if val.SessionID != session.SessionID || val.UserID != session.UserID {
					m.Lock()
					testsPassed = false
					m.Unlock()
				}
			} else {
				m.Lock()
				testsPassed = false
				m.Unlock()
			}
		}(i, &m)
	}

	if !testsPassed {
		t.Fatalf("Test failed during register or checking!")
	}
}

func TestDeleteSessionGoroutines(t *testing.T) {
	const SESSIONS = 100

	storage := CreateSessionStorage()

	var m sync.Mutex
	testsPassed := true

	for i := 0; i < SESSIONS; i++ {
		go func(i int, m *sync.Mutex) {
			session := Session{
				SessionID: fmt.Sprint(i),
				UserID:    uint32(i),
				TTL:       time.Now().Add(10 * time.Hour),
			}

			storage.RegisterNewSession(session)

			err := storage.DeleteSession(session.SessionID)
			if err != nil {
				m.Lock()
				testsPassed = false
				m.Unlock()
			}

			if val, ok := storage.CheckSession(session.SessionID); val != nil || ok {
				m.Lock()
				testsPassed = false
				m.Unlock()
			}
		}(i, &m)
	}

	if !testsPassed {
		t.Fatalf("Test failed during deletion!")
	}
}

// func TestGetSessionsSimple(t *testing.T) {
// 	const SESSIONS = 100

// 	storage := CreateSessionStorage()

// 	for i := 0; i < SESSIONS; i++ {
// 		session := Session{
// 			SessionID: fmt.Sprint(i),
// 			UserID:    uint32(i),
// 			TTL:       time.Now().Add(10 * time.Hour),
// 		}

// 		storage.RegisterNewSession(session)
// 	}

// 	sessions, err := storage.GetSessions()
// 	if err != nil {
// 		t.Fatalf("Error during getting all sessions!")
// 	}

// 	for _, session := range sessions {
// 		checkedSession, ok := storage.CheckSession(session.SessionID)

// 		if !ok || session != *checkedSession {
// 			t.Fatalf("Test failed on getting all sessions! UserId: %d", session.UserID)
// 		}
// 	}
// }

// func TestGetSessionsGoroutines(t *testing.T) {
// 	const SESSIONS = 100

// 	storage := CreateSessionStorage()

// 	var m sync.Mutex
// 	testsPassed := true

// 	for i := 0; i < SESSIONS; i++ {
// 		go func(i int, m *sync.Mutex) {
// 			session := Session{
// 				SessionID: fmt.Sprint(i),
// 				UserID:    uint32(i),
// 				TTL:       time.Now().Add(10 * time.Hour),
// 			}

// 			storage.RegisterNewSession(session)

// 			sessions, err := storage.GetSessions()
// 			if err != nil {
// 				m.Lock()
// 				testsPassed = false
// 				m.Unlock()
// 			}

// 			founded := false
// 			for _, val := range sessions {
// 				if val == session {
// 					founded = true
// 				}
// 			}

// 			if !founded {
// 				m.Lock()
// 				testsPassed = false
// 				m.Unlock()
// 			}
// 		}(i, &m)
// 	}

// 	if !testsPassed {
// 		t.Fatalf("Test failed on geting all sessions!")
// 	}
// }
