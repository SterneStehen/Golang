package main

import "fmt"

func main(){
	websites := map[int]string{
		1: "Google",
		2: "Amazon",
	}
	fmt.Println(websites)
	websites[3] = "Schwarz"
	fmt.Println(websites)
	delete(websites, 1)
	fmt.Println(websites)
	strMap := map[string]string{
		"PC": "dell",
		"Phone": "xiaomi",
	}
	fmt.Println(strMap)
}