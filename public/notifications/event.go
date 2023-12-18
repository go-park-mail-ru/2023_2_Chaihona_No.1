package notifications

const (
	PostEventId = iota
)

type Event struct {
	EventType int
	Body      map[string]any
}

// type Event interface {
// 	GetEventType() int // в табличке бд smallint
// 	GetMarshalled() ([]byte, error)
// }

// type PostEvent struct {
// 	PostId int `json:"id" db:"id"`
// }

// func (ev PostEvent) GetEventType() int {
// 	return PostEventId
// }

// func (ev PostEvent) GetMarshalled() ([]byte, error) {
// 	return json.Marshal(ev)
// }
