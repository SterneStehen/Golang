package main


import (
	"fmt"
	"math"
)

func main() {
	const inflationRat = 4.5
	var investAmount float64 = 1000
	var expectedReturnRate   = 5.5
	var years float64

	fmt.Print("invesment Amaunt: ")
	fmt.Scan(&investAmount)
	fmt.Print("expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := float64(investAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	inflatValue := investAmount * math.Pow(1+inflationRat/100, years)
	result := futureValue - inflatValue
	fmt.Println("Yours benifit is : ", result)	
	fmt.Printf("Form: Yours benifit is : %.2f\n", result)
}

// func main(){
// 	investAmount, expectedReturnRate, years := 1000.0, 5.5, 10.0
// 	futureValue :=  investAmount * math.Pow(1 + expectedReturnRate/100, years)
// 	fmt.Println(futureValue) 
// }