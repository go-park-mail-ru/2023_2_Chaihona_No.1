package sessions

type SessionRepository interface {
	RegisterNewSession(session SessionOld) error
	DeleteSession(sessionID string) error
	CheckSession(sessionID string) (*SessionOld, bool)
}
