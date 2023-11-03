package sessions

type SessionRepository interface {
	RegisterNewSession(session Session) error
	DeleteSession(sessionID string) error
	CheckSession(sessionID string) (*Session, bool)
}
