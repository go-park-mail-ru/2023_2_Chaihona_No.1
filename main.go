package main

import (
	"encoding/json"
	"fmt"
	model "project/model"
	reg "project/registration"
)

// заглушка main
func main() {
	user := model.User{
		ID:    32,
		Login: "1@gmail.com",
	}

	storage := reg.CreateUserStorage()
	storage.RegisterNewUser(user)
	temp, _ := storage.CheckUser(user.ID)
	s, _ := json.Marshal(temp)
	fmt.Println(string(s))
}
