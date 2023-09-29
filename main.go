package main

import (
	"net/http"
	auth "project/authorization"
	model "project/model"
)

// заглушка main
func main() {
	user := model.User{
		ID:       32,
		Login:    "12",
		Password: "12",
	}

	rep := auth.CreateRepoHandler()
	rep.Users.RegisterNewUser(user)


	http.HandleFunc("/", rep.Root)
	http.HandleFunc("/login", rep.Login)
	http.ListenAndServe(":8080", nil)
}
