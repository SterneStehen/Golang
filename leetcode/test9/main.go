package main

import "fmt"



func removeElement(nums []int, val int) int {
   
	var i int
	l := len(nums)
	for i = 0; i < l; i++{
		if nums[i] == val{
			nums = append(nums[:i], nums[i+1:]...)
			l = len(nums)
			i--
			}
		}
	
		fmt.Println(nums)
		return  len(nums)
	}




func main(){
	arr := []int{2,3,3,4,6}
	res := removeElement(arr, 4)
	fmt.Println(res)
}