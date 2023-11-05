package handlers

type UserForm struct {
	Body struct {
		User struct {
			ID           uint   `json:"id"`
			Nickname     string `json:"nickname"`
			Login        string `json:"login"`
			OldPassword string `json:"old_password"`
			NewPassword string `json:"new_password"`
			Status       string `json:"status"`
			Avatar       string `json:"avatar"`
			Background   string `json:"background"`
			Description  string `json:"description"`
			IsAuthor bool `json:"is_author"`
		} `json:"user"`
	} `json:"body"`
}

func (f UserForm) IsValide() bool {
	return true
}

func (f UserForm) IsEmpty() bool {
	return false
}

type PostForm struct {
	Body struct {
		MinSubLevelId uint      `json:"min_subscription_level_id"`
		Header        string    `json:"header"`
		Body          string    `json:"body,omitempty"`
	} `json:"body"`
}

func (f PostForm) IsValide() bool {
	return true
}

func (f PostForm) IsEmpty() bool {
	return false
}

type PaymentForm struct {
	DonaterId         uint   `json:"donater_id"`
	CreatorId         uint   `json:"creator_id"`
	Currency          string `json:"currency,omitempty"`
	Value             string `json:"value,omitempty"`
}

func (f PaymentForm) IsValide() bool {
	return true
}

func (f PaymentForm) IsEmpty() bool {
	return false
}