package user

import (
	"fmt"
	"time"
	"errors"
)


type User struct{
	fName string
	sName string
	bDay string
	curTime time.Time
}

func ImputScan(text string) string{
	var tmp string
	fmt.Print(text)
	fmt.Scanln(&tmp)
	return tmp
}

func (us *User) PrintStruct(){
	fmt.Println(us.fName, us.sName, us.bDay, us.curTime)
}

// func (us *User) ClearUser(){
// 	us.fName = ""
// 	us.sName = ""
// }

func NewUser(firstName, secondName, birthDay string) (*User, error){
	if firstName == "" || secondName == "" || birthDay == "" {
		return nil, errors.New("errmsg. empty data newUser")
	}
	return  &User{
	fName: firstName,
	sName: secondName,
	bDay: birthDay,
	curTime: time.Now(),
	}, nil
}
