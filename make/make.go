package main

import "fmt"

func  main()  {
	//name := []string{}
	name := make([]string, 2, 5)
	name = append(name, "Nick")
	name = append(name, "Sem")
	name[0 ] = "Serg"
	fmt.Println(name)
	ware := make(map[string]float64, 3)
	ware["Laptop"] = 55.3
	fmt.Println(ware)

	student := make(map[string][]int)
	
	student["sergii"] = append(student["sergii"], 4)
	student["sergii"] = append(student["sergii"], 2)
	student["sergii"] = append(student["sergii"], 3)
	// student["max"][0] = 4
	// student["denis"][0] = 5
	
	fmt.Println(student)
}