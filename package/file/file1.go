package file

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)
const accounName = "balance.txt" 
func GetBalance()(float64, error){
	read, err := os.ReadFile(accounName)
	if err != nil{
		return 0, errors.New("Failed to read balance  file")
	}
	readStr := string(read)
	readFloat, err := strconv.ParseFloat(readStr, 64)
	if err != nil{
		return 0, errors.New("failed to parse string to float")
	}
	fmt.Println("Read balance ", readFloat)
	return readFloat, nil
}

func WriteBalance(balance float64){
	// var data srting
	data := fmt.Sprint(balance)
	os.WriteFile(accounName, []byte(data), 0644)
}
