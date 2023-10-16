package registration

import (
	"sync"

	model "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/model"
)

type ProfileStorage struct {
	Profiles map[string]model.Profile
	Mu       sync.RWMutex
	Size     uint32
}

func CreateProfileStorage() *ProfileStorage {
	storage := &ProfileStorage{
		Profiles: make(map[string]model.Profile),
		Mu:       sync.RWMutex{},
		Size:     0,
	}

	return storage
}

func (storage *ProfileStorage) RegisterNewProfile(Profile *model.Profile) error {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()
	storage.Size++
	storage.Profiles[Profile.User.Login] = *Profile
	return nil
}

func (storage *ProfileStorage) DeleteProfile(login string) error {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	if _, ok := storage.Profiles[login]; !ok {
		return ErrNoSuchProfile
	}

	delete(storage.Profiles, login)
	storage.Size--

	return nil
}

func (storage *ProfileStorage) CheckProfile(login string) (*model.Profile, bool) {
	storage.Mu.RLock()
	defer storage.Mu.RUnlock()

	val, ok := storage.Profiles[login]

	if ok {
		copy := val

		return &copy, true
	}

	return nil, false
}

func (storage *ProfileStorage) GetProfile(id uint) (*model.Profile, bool) {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	for _, profile := range storage.Profiles {
		if profile.User.ID == id {
			copy := profile
			return &copy, true
		}
	}

	return nil, false
}

func (storage *ProfileStorage) GetProfiles() ([]model.Profile, error) {
	storage.Mu.RLock()
	defer storage.Mu.RUnlock()

	profiles := make([]model.Profile, storage.Size)

	i := 0
	for _, profile := range storage.Profiles {
		profiles[i] = profile
		i++
	}

	return profiles, nil
}
