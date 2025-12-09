package main

import (
	"fmt"

	"github.com/ahadahamedakash/learning-golang/auth"
	"github.com/ahadahamedakash/learning-golang/user"
)

func main() {
	auth.LoginWithCredentials("ahad", "password")

	sess := auth.GetSession()

	fmt.Println("USER: ", sess)

	user := user.User{

		Email: "demo@mail.com",
		Name:  "Jhon Doe",
	}

	fmt.Println(user.Email, user.Name)
}
