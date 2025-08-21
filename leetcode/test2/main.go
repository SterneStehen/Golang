package main

import "fmt"
import "strconv"


func isPalindrome(x int) bool {
    str := strconv.Itoa(x)
	
	lenstr := len(str) - 1
	fmt.Println("string ", str, "len ", lenstr)
	for i := 0 ; i <= lenstr ; i++{
		if(str[i] != str[lenstr-i]){
			return false
		}
		fmt.Println("digit ", str[i],  " =  lenstr", str[lenstr-i] )
	}
	return true
}

func main(){
	x := 121
	res := isPalindrome(x)
	if res == true{
		fmt.Println("True")
	} else{
		fmt.Println("False")
	}
}