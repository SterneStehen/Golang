package main

import (
	"fmt"
)

type Product struct{
	title string
	id int
	prise float64
}



func main(){
	prodName := [4]string{"book", "pencil"}
	prise := [4]float64{22.3, 22.5, 44.5, 63.2}
	fmt.Println(prodName)
	fmt.Println(prise)
	prise[2] = 12.5
	fmt.Println(prise)
	newPrise := prise[0:2]
	newPrise[1] = 99.9
	fmt.Println(newPrise)
	fmt.Println(prise)
	fmt.Println("price len = ", len(prise), "prise cap =", cap(prise))
	fmt.Println("newPrise len = ", len(newPrise), "newPrise cap =", cap(newPrise))
}
