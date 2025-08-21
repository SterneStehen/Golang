package main

import "fmt"


func removeDuplicates(nums []int) (int) {
    var res []int
	k := 0
	for i := 0; i < len(nums)-1; i++{
		if nums[i] == nums[i+1]{
			continue
		} else{
			res = append(res, nums[i])
			nums[k] = nums[i]
			k++
			
		}

	}
    res = append(res, nums[len(nums)-1])
	
	nums[k] = nums[len(nums)-1]
	for i := k+1; i < len(nums); i++{
		nums[i] = '_'
	}
	//nums = nums[:k]
	// nums = nil
	// for i, ch :=range res{
	// 	nums = append(nums, res[i])
	// 	fmt.Println("ch = ", ch)
	// }
	
	//fmt.Println(nums) 
	return  len(res)
}

func main(){
	arr := []int{1,1,2}
	res := removeDuplicates(arr)
	fmt.Println( arr, res)
	
}