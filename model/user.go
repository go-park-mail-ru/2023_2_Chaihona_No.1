package model

const CreatorStatus = "creator"
const SimpleUserStatus = "simple_user"

type User struct {
	ID         uint   `json:"id"`
	Login      string `json:"login,string"`
	Password   string `json:"-"`
	Avatar     string `json:"avatar,string"`
	Background string `json:"background,string"`
	UserType   string `json:"user_type,string"`
	Status     string `json:"status,string"`
}
