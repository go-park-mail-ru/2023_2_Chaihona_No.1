package registration

import (
	model "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/model"
)

type ProfileRepository interface {
	RegisterNewProfile(user *model.Profile) *ErrorRegistration
	DeleteProfile(login string) error
	CheckProfile(login string) (*model.Profile, bool)
	GetProfiles() ([]model.Profile, error)
	GetProfile(id uint) (*model.Profile, bool)
}
