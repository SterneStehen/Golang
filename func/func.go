package main

import "fmt"

type tFunc  func(int) int

func doubleNum(input *[]int) []int{
	DNum := []int{}
	for _, val := range *input{
		DNum = append(DNum, val*2)
	}
	return DNum
}

func doub(i int) int{
	return i*2
}
func tri(i int) int{
	return i*3
}

func transform(input *[]int, funcRes tFunc) []int{
	DNum := []int{}
	for _, val := range *input{
		DNum = append(DNum, funcRes(val))
	}
	return DNum
}

func main(){
	number := []int{20, 30, 40}
	for i, value := range number{
		fmt.Println("i = ", i, "val = ", value)
	}

	mapN := map[int]string{
		44: "Nick",
		12: "serg",
		32: "Sem",
	}

	for i, val := range mapN{
		fmt.Println("i = ", i, "val = ", val)
	}

	digit := []int{34, 55, 2}
	Multidig := doubleNum(&digit)
	fmt.Println(Multidig)

	fmt.Println("======================================")
	d2 := transform(&digit, doub)
	fmt.Println(d2)
	d3 := transform(&digit, tri)
	fmt.Println(d3)
}