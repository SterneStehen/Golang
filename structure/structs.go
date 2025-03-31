package main

import (
	"fmt"
	"structure/userF"
)

func main(){
	firstName := user.ImputScan("Enter First Name: ")
	secondName := user.ImputScan("Enter Second Name: ")
	birthDay := user.ImputScan("Enter Brith Day: ")

	var appUser *user.User
	appUser, err := user.NewUser(firstName, secondName, birthDay)
	if err != nil{
		fmt.Println(err);
		return
	}
	appUser.PrintStruct()
	// appUser.ClearUser()
	// appUser.PrintStruct()
}