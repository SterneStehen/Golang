package main

import (
	"fmt"
	"os"
	"strconv"
)

const accounName = "balance.txt"

func getBalance()(float64){
	read, _ := os.ReadFile(accounName)
	readStr := string(read)
	readFloat, _ := strconv.ParseFloat(readStr, 64)
	fmt.Println("read file result is ", readFloat)
	return readFloat
}

func writeBalance(balance float64){
	// var data srting
	data := fmt.Sprint(balance)
	os.WriteFile(accounName, []byte(data), 0644)
}

func main(){
	var depositAmount float64
	var depositWithdtaw float64
	accounBallance := getBalance()

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
				writeBalance(accounBallance)
				fmt.Println("\nbalance updated. new amout: ", accounBallance)
				} else if choice == 3{
					fmt.Print("Your Withdtaw: ")
					fmt.Scan(&depositWithdtaw)
					if depositWithdtaw > accounBallance{
						fmt.Println("\nError. deposit Withdtaw mare than balans. try again")
						continue
					}
					accounBallance = accounBallance - depositWithdtaw
					writeBalance(accounBallance)
					fmt.Println("\nbalance updated. new amout: ", accounBallance)
					} else if choice == 4{
						getBalance()
						fmt.Println("Good bay!")
			return	
		} else {
			fmt.Println("Error. Try again")
		}
	}

}