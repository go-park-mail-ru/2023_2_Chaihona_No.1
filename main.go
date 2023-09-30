package main

import (
	"fmt"
	"net/http"
	auth "project/authorization"
	handlers "project/handlers"
	model "project/model"

	reg "project/registration"
	"strings"

	"github.com/gorilla/mux"
)

// заглушка main
func main() {
	user := model.User{
		ID:       32,
		Login:    "12",
		Password: "12",
	}

	registrationHandler := reg.CreateRepoHandler()
	profileHandler := handlers.CreateProfileHandler()
	postHandler := handlers.CreatePostHandler()
	rep := auth.CreateRepoHandler()
	rep.Users.RegisterNewUser(&user)

	go func() {
		resp, err := http.Post("http://127.0.0.1:8080/registration", "application/json", strings.NewReader(`{ "body" : {"login" : "123", "password" : "123", "user_type" : "simple_user"}}`))

		if err != nil {
			fmt.Println("error")
			return
		}

		fmt.Println("Cookies:", resp.Cookies(), "Status:", resp.StatusCode)

		resp, err = http.Post("http://127.0.0.1:8080/registration", "application/json", strings.NewReader(`{ "body" : {"login" : "1234", "password" : "1234", "user_type" : "simple_user"}}`))

		if err != nil {
			fmt.Println("error")
			return
		}

		fmt.Println("Cookies:", resp.Cookies(), "Status:", resp.StatusCode)

		resp, err = http.Post("http://127.0.0.1:8080/registration", "application/json", strings.NewReader(`{ "body" : {"login" : "12345", "password" : "12345", "user_type" : "simple_user"}}`))

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

		resp, err = http.Post("http://127.0.0.1:8080/login", "application/json", strings.NewReader(`{ "body" : {"login" : "1234", "password" : "1234"}}`))

		if err != nil {
			fmt.Println("error")
			return
		}

		fmt.Println("Cookies:", resp.Cookies(), "Status:", resp.StatusCode)

		resp, err = http.Post("http://127.0.0.1:8080/login", "application/json", strings.NewReader(`{ "body" : {"login" : "12345", "password" : "123"}}`))

		if err != nil {
			fmt.Println("error")
			return
		}

		fmt.Println("Cookies:", resp.Cookies(), "Status:", resp.StatusCode)
	}()

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/login", rep.Login).Methods("POST")
	// r.HandleFunc("/api/v1/logout", rep.Logout).Methods("POST")
	r.HandleFunc("/api/v1/registration", registrationHandler.SignUp).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", profileHandler.GetInfo).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/post", postHandler.GetAllUserPosts).Methods("GET")

	http.ListenAndServe(":8080", r)
}
