package handlers

type UserForm struct {
	Body struct {
		ID           uint   `json:"id"`
		Nickname     string `json:"nickname"`
		Login        string `json:"login"`
		UserType     string `json:"user_type"`
		Status       string `json:"status"`
		Avatar       string `json:"avatar"`
		Background   string `json:"background"`
		Description  string `json:"description"`
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
		ID            uint      `json:"id"`
		MinSubLevelId uint      `json:"min_subscription_level_id"`
		Header        string    `json:"header"`
		Body          string    `json:"body,omitempty"`
		Likes         uint      `json:"likes" db:"likes"`
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