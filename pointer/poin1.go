package main

import "fmt"

func changeValue(x *int){
	*x = 100
}

func main(){
	var x int = 10
	var p *int = &x
	var n *int

	fmt.Println("x = ", x)
	fmt.Println("p = ", p)
	fmt.Println("*p = ", *p)
	changeValue(p)
	fmt.Println("x = ", x)
	if(n == nil){
		fmt.Println("nil pointer")
	}
}