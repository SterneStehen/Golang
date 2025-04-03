package main

import "fmt"

func addSome(a, b interface{}) interface{}  {
	aInt, aOk := a.(int)
	bInt, bOk := b.(int)
	
	if aOk && bOk{
		return aInt+bInt
	} 
	return nil
}

func main(){
	a := 3
	b := 4
	fmt.Println(addSome(a, b))
}