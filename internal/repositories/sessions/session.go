package sessions

import (
	"time"
)

type SessionOld struct {
	SessionID string    `json:"session_id"`
	UserID    uint32    `json:"user_id"`
	TTL       time.Time `json:"ttl"`
}
