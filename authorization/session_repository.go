package authorization

type SessionRepository interface {
	RegisterNewSession(session Session) error
	DeleteSession(sessionId string) error
	CheckSession(sessionId string) (*Session, bool)
	GetSessions() ([]Session, error)
}
