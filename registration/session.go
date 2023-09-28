package registration

import (
	"time"
)

type Session struct {
	SessionId string    `json:"session_id"`
	UserId    uint32    `json:"user_id"`
	Ttl       time.Time `json:"ttl"`
}
