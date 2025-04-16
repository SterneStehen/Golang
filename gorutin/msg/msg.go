package msg


import "fmt"
import "time"


func Great(msg string){
	fmt.Println("Hello, ", msg)
}

func sleepGreat(msg string){
	time.Sleep(time.Second )
	fmt.Println("Hello! ", msg)
}


func GreatMain(){
	go Great("Haw are you?")
	go Great("Nice to meet you")
	go sleepGreat("Haw... are... you....?")
	go Great("I hope are you like ")
	time.Sleep(time.Second*2)
}