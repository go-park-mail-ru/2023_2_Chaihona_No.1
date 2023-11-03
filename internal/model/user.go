package model

const (
	CreatorStatus    = "creator"
	SimpleUserStatus = "simple_user"
)

type User struct {
	ID          uint   `json:"id"`
	Nickname    string `json:"nickname"`
	Login       string `json:"login"`
	Password    string `json:"-"`
	Avatar      string `json:"avatar"`
	Background  string `json:"background"`
	UserType    string `json:"user_type"`
	Status      string `json:"status"`
	Description string `json:"description"`
}
