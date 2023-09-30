package main

import (
	"fmt"
	"net/http"
	auth "project/authorization"
	model "project/model"
	"strings"
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

	go func() {
		resp, err := http.Post("http://127.0.0.1:8080/registration", "application/json", strings.NewReader(`{ "body" : {"login" : "123", "password" : "123", "user_type" : "simple_user"}}`))

		if err != nil {
			fmt.Println("error")
			return
		}

		fmt.Println("Cookies:", resp.Cookies(), "Status:", resp.StatusCode)

		resp, err = http.Post("http://127.0.0.1:8080/login", "application/json", strings.NewReader(`{ "body" : {"login" : "123", "password" : "123"}}`))
		
		if err != nil {
			fmt.Println("error")
			return
		}

		fmt.Println("Cookies:", resp.Cookies(), "Status:", resp.StatusCode)
	}()

	http.HandleFunc("/", rep.Root)
	http.HandleFunc("/login", rep.Login)
	http.HandleFunc("/registration", rep.Signup)
	http.ListenAndServe(":8080", nil)
}
