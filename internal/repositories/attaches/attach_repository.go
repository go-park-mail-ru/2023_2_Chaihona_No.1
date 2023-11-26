package attaches

import (
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

type AttachRepository interface {
	PinAttach(attach model.Attach) (int, error)
	GetPostAttaches(postID int) ([]model.Attach, error)
	// DeleteAttach(id uint) error
}