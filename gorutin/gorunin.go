package main

import "fmt"
import "time"
import "gorutin/channel"
import "gorutin/msg"

func loopSay(name string){
	for i:=1; i <= 3; i++{
		fmt.Println(name, "say hello", i)
		time.Sleep(time.Second)
	}
}

func say(msg string){
	fmt.Println(msg)
}

func main(){
	go say("run go rutin")
	fmt.Println("hello")
	time.Sleep(time.Second)
	go loopSay("Nik")
	go loopSay("Daryna")
	time.Sleep(time.Second *3)
	channel.Chann()
	msg.GreatMain()
}