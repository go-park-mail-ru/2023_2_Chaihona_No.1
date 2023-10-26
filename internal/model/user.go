package model

const (
	CreatorStatus    = "creator"
	SimpleUserStatus = "simple_user"
)

type User struct {
	ID           uint   `json:"id" db:"id"`
	Nickname     string `json:"nickname" db:"nickname"`
	Login        string `json:"login" db:"email"`
	Password     string `json:"-" db:"password"`
	UserType     string `json:"user_type" db:""`
	Status       string `json:"status" db:"status"`
	Avatar       string `json:"avatar" db:"avatar_path"`
	Background   string `json:"background" db:"background_path"`
	Description  string `json:"description" db:"description"`
	CreationDate string `json:"-" db:"creation_date"`
	LastUpdate   string `json:"-" db:"last_update"`
	Is_author    bool   `db:"is_author"`
	Subscribers  uint   `json:"-" db:"subscribers"`
}
