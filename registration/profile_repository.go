package registration

import (
	model "project/model"
)

type ProfileRepository interface {
	RegisterNewProfile(user *model.Profile) error
	DeleteProfile(login string) error
	CheckProfile(login string) (*model.Profile, bool)
	GetProfiles() ([]model.Profile, error)
}

