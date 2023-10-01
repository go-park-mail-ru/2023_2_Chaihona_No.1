package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	handlers "project/handlers"
	model "project/model"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	user := model.User{
		ID:       32,
		Login:    "12",
		Password: "12",
	}

	// registrationHandler := reg.CreateRepoHandler()

	// profileHandler := handlers.CreateProfileHandler()
	// postHandler := handlers.CreatePostHandler()
	rep := handlers.CreateRepoHandler()
	profileHandler := handlers.CreateProfileHandlerViaRepos(&rep.Sessions, &rep.Profiles)
	postStorage := model.CreatePostStorage()
	postStorage.CreateNewPost(model.Post{
		ID:           1,
		AuthorID:     4,
		HasAccess:    true,
		Access:       model.EveryoneAccess,
		CreationDate: "15:08 30.09.2023",
		Header:       "Header",
		Body:         "Body",
		Likes:        10,
	})
	postHandler := handlers.CreatePostHandlerViaRepos(&rep.Sessions, &postStorage, &rep.Profiles)

	rep.Users.RegisterNewUser(&user)

	go func() {
		// resp, err := http.Post("https://127.0.0.1:8081/api/v1/registration", "application/json", strings.NewReader(`{ "body" : {"login" : "123", "password" : "123", "user_type" : "simple_user"}}`))
		caCert, err := ioutil.ReadFile("request_keys/rootCA.crt")
		if err != nil {
			log.Fatal(err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: caCertPool,
				},
				MaxIdleConns:       10,
				IdleConnTimeout:    30 * time.Second,
				DisableCompression: true,
			},
		}
		resp, err := client.Post("https://127.0.0.1:8001/api/v1/registration", "application/json", strings.NewReader(`{ "body" : {"login" : "123", "password" : "123", "user_type" : "simple_user"}}`))

		if err != nil {
			fmt.Println("error")
			return
		}

		fmt.Println("Cookies:", resp.Cookies(), "Status:", resp.StatusCode)

		// resp, err = http.Post("http://127.0.0.1:8081/api/v1/registration", "application/json", strings.NewReader(`{ "body" : {"login" : "1234", "password" : "1234", "user_type" : "simple_user"}}`))

		// if err != nil {
		// 	fmt.Println("error")
		// 	return
		// }

		// fmt.Println("Cookies:", resp.Cookies(), "Status:", resp.StatusCode)

		// resp, err = http.Post("http://127.0.0.1:8081/api/v1/registration", "application/json", strings.NewReader(`{ "body" : {"login" : "12345", "password" : "12345", "user_type" : "simple_user"}}`))

		// if err != nil {
		// 	fmt.Println("error")
		// 	return
		// }

		// fmt.Println("Cookies:", resp.Cookies(), "Status:", resp.StatusCode)

		// resp, err = http.Post("http://127.0.0.1:8081/api/v1/login", "application/json", strings.NewReader(`{ "body" : {"login" : "123", "password" : "123"}}`))

		// if err != nil {
		// 	fmt.Println("error")
		// 	return
		// }

		// fmt.Println("Cookies:", resp.Cookies(), "Status:", resp.StatusCode)

		// resp, err = http.Post("http://127.0.0.1:8081/api/v1/login", "application/json", strings.NewReader(`{ "body" : {"login" : "1234", "password" : "1234"}}`))

		// if err != nil {
		// 	fmt.Println("error")
		// 	return
		// }

		// fmt.Println("Cookies:", resp.Cookies(), "Status:", resp.StatusCode)
		// kuka := resp.Cookies()

		// resp, err = http.Post("http://127.0.0.1:8081/api/v1/login", "application/json", strings.NewReader(`{ "body" : {"login" : "12345", "password" : "123"}}`))

		// if err != nil {
		// 	fmt.Println("error")
		// 	return
		// }

		// fmt.Println("Cookies:", resp.Cookies(), "Status:", resp.StatusCode)

		// // resp, err = http.Post("http://127.0.0.1:8081/api/v1/login", "application/json", strings.NewReader(`{ "body" : {"login" : "12345", "password" : "123"}}`))
		// req, _ := http.NewRequest("GET", "http://127.0.0.1:8081/api/v1/profile/1", nil)
		// req.AddCookie(kuka[0])

		// resp, err = http.DefaultClient.Do(req)
		// fmt.Println("Cookies:", resp.Cookies(), "Status:", resp.StatusCode)

		// // body, _ := ioutil.ReadAll(resp.Body)
		// // fmt.Println(string(body))

		// if err != nil {
		// 	fmt.Println("error")
		// 	return
		// }

		// req, _ = http.NewRequest("GET", "http://127.0.0.1:8081/api/v1/profile/1/post", nil)
		// req.AddCookie(kuka[0])

		// resp, err = http.DefaultClient.Do(req)
		// fmt.Println("Cookies:", resp.Cookies(), "Status:", resp.StatusCode)
		// // body, _ := ioutil.ReadAll(resp.Body)
		// // fmt.Println(string(body))

		// if err != nil {
		// 	fmt.Println("error")
		// 	return
		// }

		// resp, err = http.Post("http://127.0.0.1:8081/api/v1/registration", "application/json", strings.NewReader(`{ "body" : {"login" : "abcd", "password" : "abcd", "user_type" : "creator"}}`))

		// req, _ = http.NewRequest("GET", "http://127.0.0.1:8081/api/v1/profile/4/post", nil)
		// req.AddCookie(kuka[0])
		// resp, err = http.DefaultClient.Do(req)
		// fmt.Println("Cookies:", resp.Cookies(), "Status:", resp.StatusCode)
		// body, _ := ioutil.ReadAll(resp.Body)
		// fmt.Println(string(body))
		// if err != nil {
		// 	fmt.Println("error")
		// 	return
		// }
	}()

	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(handlers.OptionsHandler)
	r.HandleFunc("/api/v1/login", rep.Login).Methods("POST")
	// r.HandleFunc("/api/v1/logout", rep.Logout).Methods("POST")
	r.HandleFunc("/api/v1/registration", rep.Signup).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", profileHandler.GetInfo).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/post", postHandler.GetAllUserPosts).Methods("GET")

	fmt.Println("start")
	// http.ListenAndServeTLS(":8001", "cert.pem", "key.pem", r)
	http.ListenAndServeTLS(":8001", "request_keys/secure.domain.com.crt", "request_keys/secure.domain.com.key", r)
	// http.ListenAndServe(":8081", r)
}
