package main

import (
	"fmt"
	"package/bank"
	"package/file"
	"github.com/Pallinder/go-randomdata"

)

func main(){
	var depositAmount float64
	var depositWithdtaw float64
	accounBallance, err := file.GetBalance()
	if err != nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		//panic("Can`t continue, sorry.")
	}
	fmt.Println("\nWelcome to Go Bank!")
	fmt.Println("Reach us 24/7 tel.", randomdata.PhoneNumber())
	for {
		bank.PresentOption()

		var choice int
		fmt.Print("Your choice:  ")
		fmt.Scan(&choice)
		fmt.Println("Your choice is:" , choice)

		if depositAmount < 0{
			fmt.Println("invalid amount. Must be greater than 0 ")
		}

		
		if choice == 1{
			fmt.Println("\nYours balans is ", accounBallance)
			} else if choice == 2{
				fmt.Print("Your deposit: ")
				fmt.Scan(&depositAmount)
				accounBallance = accounBallance + depositAmount
				file.WriteBalance(accounBallance)
				fmt.Println("\nbalance updated. new amout: ", accounBallance)
				} else if choice == 3{
					fmt.Print("Your Withdtaw: ")
					fmt.Scan(&depositWithdtaw)
					if depositWithdtaw > accounBallance{
						fmt.Println("\nError. deposit Withdtaw mare than balans. try again")
						continue
					}
					accounBallance = accounBallance - depositWithdtaw
					file.WriteBalance(accounBallance)
					fmt.Println("\nbalance updated. new amout: ", accounBallance)
					} else if choice == 4{
						//_, err := getBalance()
						fmt.Println("Good bay!")
			return	
		} else {
			fmt.Println("Error. Try again")
		}
	}

}