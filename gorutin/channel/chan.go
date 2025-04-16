package channel

import "fmt"

func Chann(){
	numbers := make(chan int)
	go func(){
	for i:=0; i < 5; i++{
		numbers <-i
	}
	close(numbers)
	}()

	for val := range numbers{
		fmt.Println("recive numbers: ", val)
	}
}