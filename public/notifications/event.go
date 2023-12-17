package notifications

import "encoding/json"

const (
	PostEventId = iota
)

type Event interface {
	GetEventType() int // в табличке бд smallint
	GetMarshalled() ([]byte, error)
}

type PostEvent struct {
	PostId int `json:"id" db:"id"`
}

func (ev PostEvent) GetEventType() int {
	return PostEventId
}

func (ev PostEvent) GetMarshalled() ([]byte, error) {
	return json.Marshal(ev)
}
