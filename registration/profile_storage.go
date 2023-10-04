package registration

import (
	"fmt"
	model "project/model"
	"sync"
)

type ProfileStorage struct {
	Profiles map[string]model.Profile
	Mu       sync.Mutex
	Size     uint32
}

func CreateProfileStorage() *ProfileStorage {
	storage := &ProfileStorage{
		Profiles: make(map[string]model.Profile),
	}

	return storage
}

func (storage *ProfileStorage) RegisterNewProfile(Profile *model.Profile) error {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()
	storage.Size++
	fmt.Println("Registrate profile input")
	fmt.Println(Profile)
	storage.Profiles[Profile.User.Login] = *Profile
	fmt.Println("out")
	fmt.Println(storage.Profiles[Profile.User.Login])
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

func (storage *ProfileStorage) GetProfile(id uint) (*model.Profile, bool) {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	for _, profile := range storage.Profiles {
		if profile.User.ID == id {
			fmt.Println("Get profile by id")
			fmt.Println(storage.Profiles[profile.User.Login])
			copy := profile
			return &copy, true
		}
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
