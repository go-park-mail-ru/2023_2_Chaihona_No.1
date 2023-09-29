package main

import (
	"encoding/json"
	"fmt"
	model "project/model"
)

// заглушка main
func main() {
	var user model.User
	s, _ := json.Marshal(user)
	fmt.Println(string(s))
}
