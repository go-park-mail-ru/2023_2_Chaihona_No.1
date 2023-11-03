package handlers

//swagger:model result
type Result struct {
	Body interface{} `json:"body,omitempty"`
	Err  string      `json:"error,omitempty"`
}