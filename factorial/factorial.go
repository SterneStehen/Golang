package main

import "fmt"

func fact(i int ) int{
	//var k int
	if i == 1{
		return 1
	}
	return i * fact	(i - 1)
}

func random(num ...int) int{
	sum := 0
	for _, val := range num{
		sum += val
	}
	return sum
}

func main(){

	res := fact(5)
	fmt.Println(res)
	numbers := random(1, 4, 55, 8, 23, 19)
	fmt.Println(numbers)
}