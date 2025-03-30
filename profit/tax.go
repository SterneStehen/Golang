package main

import "fmt"
import "os"

func writeBalance(EBT, profit, ratio float64){
	data := fmt.Sprint("EBT is ", EBT, "\nprofit is ", profit, "\nratio is ", ratio)

	os.WriteFile("ratio.txt", []byte(data), 0644)
}

func main(){
	var revenue float64
	var expenses float64
	var tax float64

	revenue, expenses, tax = getUserInput(revenue, expenses, tax)

	if revenue <= 0 || expenses <= 0 || tax <= 0 {
		//fmt.Println("Error. Value is less then 0")
		panic("Error. Value is less then 0")
	}
	EBT := funcEBT(revenue, expenses)
	profit := funcProfit(EBT, tax)
	ratio := funcRatio(EBT, profit)

	writeBalance(EBT, profit, ratio)
	fmt.Println( "EBT is ", EBT, " profit is ", profit , " ratio is ", ratio)

}

func getUserInput(revenue, expenses, tax float64)(float64, float64, float64){
	fmt.Print("Yours revenue is ")
	fmt.Scan(&revenue)
	fmt.Print("Yours expenses is ")
	fmt.Scan(&expenses)
	fmt.Print("Yours tax is ")
	fmt.Scan(&tax)
	return revenue, expenses, tax
}

func funcEBT(revenue, expenses float64) (float64){
	return revenue - expenses
}
func funcProfit(EBT, tax float64)(float64){
	return EBT - (EBT/100 * tax)
}
func funcRatio(EBT, profit float64)(float64){
	res := EBT/profit
	//fmt.Println(EBT, profit, res)
	return res
}