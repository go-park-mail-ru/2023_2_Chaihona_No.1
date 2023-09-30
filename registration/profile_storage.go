package registration

import (
	model "project/model"
	"sync"
)

type ProfileStorage struct {
	Profiles map[string]model.Profile
	Mu    sync.Mutex
	Size  uint32
}

func CreateProfileStorage() *ProfileStorage {
	storage := &ProfileStorage{
		Profiles: make(map[string]model.Profile),
	}

	return storage
}

func (storage *ProfileStorage) RegisterNewProfile(Profile model.Profile) error {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	Profile.ID = uint(storage.Size)
	storage.Size++
	storage.Profiles[Profile.User.Login] = Profile

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
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	val, ok := storage.Profiles[login]

	if ok {
		copy := val

		return &copy, true
	}

	return nil, false
}

func (storage *ProfileStorage) GetProfiles() ([]model.Profile, error) {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	Profiles := make([]model.Profile, 0)

	for _, Profile := range storage.Profiles {
		Profiles = append(Profiles, Profile)
	}

	return Profiles, nil
}
