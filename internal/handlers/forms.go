package handlers

import (
	"mime/multipart"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

type StatusForm struct {
	Body struct {
		Status string `json:"status"`
	} `json:"body"`
}

func (f StatusForm) IsValide() bool {
	return true
}

func (f StatusForm) IsEmpty() bool {
	return false
}

type DescriptionForm struct {
	Body struct {
		Description string `json:"description"`
	} `json:"body"`
}

func (f DescriptionForm) IsValide() bool {
	return true
}

func (f DescriptionForm) IsEmpty() bool {
	return false
}

type UserForm struct {
	Body struct {
		User struct {
			ID          uint   `json:"id"`
			Nickname    string `json:"nickname"`
			Login       string `json:"login"`
			OldPassword string `json:"old_password"`
			NewPassword string `json:"new_password"`
			Status      string `json:"status"`
			Avatar      string `json:"avatar"`
			Background  string `json:"background"`
			Description string `json:"description"`
			IsAuthor    bool   `json:"is_author"`
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
		Id            string `json:"id"`
		MinSubLevelId uint   `json:"min_subscription_level_id"`
		Header        string `json:"header"`
		Body          string `json:"body,omitempty"`
		Tags          []model.Tag    `json:"tags,omitempty"`
		Attaches []model.Attach `json:"attaches,omitempty"`
		Pinned   struct {
			Files   []model.Attach `json:"files,omitempty"`
			Deleted []string       `json:"deleted"`
		} `json:"pinned,omitempty"`
	} `json:"body"`
}

func (f PostForm) IsValide() bool {
	return true
}

func (f PostForm) IsEmpty() bool {
	return false
}

type PaymentForm struct {
	Body struct {
		DonaterId uint   `json:"donater_id"`
		CreatorId uint   `json:"creator_id"`
		Currency  string `json:"currency,omitempty"`
		Value     string `json:"value,omitempty"`
	} `json:"body"`
}

func (f PaymentForm) IsValide() bool {
	return true
}

func (f PaymentForm) IsEmpty() bool {
	return false
}

type FileForm struct {
	Form multipart.Form
}

func (f FileForm) IsValide() bool {
	return true
}

func (f FileForm) IsEmpty() bool {
	return false
}

type FollowForm struct {
	Body struct {
		SubscriptionLevelId int `json:"id"`
		SubscriptionId      int `json:"subscription_id"`
	} `json:"body"`
}

func (f FollowForm) IsValide() bool {
	return true
}

func (f FollowForm) IsEmpty() bool {
	return false
}

type RatingForm struct {
	Body struct {
		Rating int `json:"rating"`
	} `json:"body"`
}

func (f RatingForm) IsValide() bool {
	return true
}

func (f RatingForm) IsEmpty() bool {
	return false
}

type CommentForm struct {
	Body struct {
		PostId int    `json:"post_id"`
		Text   string `json:"text"`
	} `json:"body"`
}

func (f CommentForm) IsValide() bool {
	return true
}

func (f CommentForm) IsEmpty() bool {
	return false
}

type DeviceIdForm struct {
	Body struct {
		DeviceId string `json:"device_id"`
	} `json:"body"`
}

func (d DeviceIdForm) IsEmpty() bool {
	return false
}

func (d DeviceIdForm) IsValide() bool {
	return true
}