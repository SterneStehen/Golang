package main

import "fmt"

func main(){
	accounBallance := 1000.0
	var depositAmount float64
	var depositWithdtaw float64

	fmt.Println("Welcome to Go Bank!")
	for {
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("1.Check balance")
		fmt.Println("2.Deposit money")
		fmt.Println("3.Withdtaw money")
		fmt.Println("4.Exit")

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
				fmt.Println("\nbalance updated. new amout: ", accounBallance)
				} else if choice == 3{
					fmt.Print("Your Withdtaw: ")
					fmt.Scan(&depositWithdtaw)
					if depositWithdtaw > accounBallance{
						fmt.Println("\nError. deposit Withdtaw mare than balans/")
						continue
					}
					accounBallance = accounBallance - depositWithdtaw
					fmt.Println("\nbalance updated. new amout: ", accounBallance)
					} else if choice == 4{
						fmt.Println("Good bay!")
			return	
		} else {
			fmt.Println("Error. Try again")
		}
	}
	
}