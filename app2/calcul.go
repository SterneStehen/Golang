package main

import "fmt"
import "math"


const inflRate = 2.5
func main(){

	var investAmount float64
	var expectReturnRate float64
	var years float64
	var futurValue float64
	var realFuturValue float64

	//inflationRate := 4.0

	outputText(" invest Amount : ")
	fmt.Scan(&investAmount)

	outputText(" expect Return Rate : ")
	fmt.Scan(&expectReturnRate)

	outputText(" years : ")
	fmt.Scan(&years)

	futurValue, realFuturValue = calculateValue(investAmount, expectReturnRate, years)
	fmt.Printf(" futurValue is %.1f\n realFuturValue is %.1f\n", futurValue, realFuturValue)
}

func outputText(text string){
	fmt.Print(text)
}

func calculateValue(investAmount, expectReturnRate, years float64)(fv, rfv float64){
	fv = investAmount * math.Pow(1+expectReturnRate/100, years)
	rfv = fv / math.Pow(1+ inflRate/100, years)

	return fv, rfv
}