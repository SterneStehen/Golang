package main
import "fmt"



func twoSum(nums []int, target int) []int {
    // Create a map to store number -> index
    seen := make(map[int]int)

    for i, num := range nums {
        complement := target - num
        if j, found := seen[complement]; found {
            return []int{j, i}
        }
        seen[num] = i
    }

    return nil 
}


func main(){
	arr := []int{2,3,5}
	target := 8

	result := twoSum(arr, target)

	fmt.Println("result is", result)

}