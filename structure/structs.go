package main

import (
	"fmt"
	"structure/userF"
)

func main(){
	firstName := user.ImputScan("Enter First Name: ")
	secondName := user.ImputScan("Enter Second Name: ")
	birthDay := user.ImputScan("Enter Brith Day: ")

	var admin *user.Admin
	admin = user.NewAdmin("7353718@gmail.com", "pass123")
	var appUser *user.User
	appUser, err := user.NewUser(firstName, secondName, birthDay)
	if err != nil{
		fmt.Println(err);
		return
	}
	appUser.PrintStruct()
	admin.PrintStruct()
	// appUser.ClearUser()
	// appUser.PrintStruct()
}