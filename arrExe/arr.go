package main

import (
	"fmt"
)

type Product struct{
	title string
	id int
	price float64
}

func NewPoduct(t string, i int, p float64) Product{
	return Product{
		title: t,
		id: i,
		price: p,
	}
}

func main(){
	var hobbies = [3]string{"yoga", "fotball", "ball"}
	fmt.Println(hobbies)
	//2
	fmt.Println(hobbies[0])
	newHobbies := hobbies[1:3]
	fmt.Println(newHobbies)
	//3
	newHobbies01 := hobbies[:2]
	fmt.Println(newHobbies01)
	newHobbies0 := hobbies[:1]
	newHobbies1 := hobbies[1:2]
	fmt.Println(newHobbies0, newHobbies1)
	newHobbies01Plus := append(newHobbies0, "forball")
	fmt.Println(newHobbies01Plus)
	//4
	addHobies := hobbies[:]
	addHobies[1] = "dance"
	addHobies[2] = "music"
	// addHobies := hobbies[:]
	// addHobies = append(addHobies, "fly")
	// fmt.Println(hobbies)
	// fmt.Println(addHobies)
	fmt.Println(hobbies)
	fmt.Println("======================")
	//5
	var gols = []string{"Go", "work"}
	gols[0] = "C++"
	gols[1] = "C"
	gols =  append(gols, "java")
	fmt.Println(gols)
	
	var P Product
	P = NewPoduct("book", 2, 3.3)
	fmt.Println(P)

	var dList = []Product{
		NewPoduct("book", 23424, 3.3),
		NewPoduct("lapTop", 2423, 22.3)}
	dList = append(dList, NewPoduct("cup", 4454, 64.2))
	fmt.Println(dList)
	PP := Product{
		"phone", 
		43, 
		22.6,}
	fmt.Println(PP)
}