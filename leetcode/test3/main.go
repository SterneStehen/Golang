package main
import "fmt"

func romanToInt(s string) int {
    // I := 1
	// V := 5
	// X := 10
	// L := 50
	// C := 100
	// D := 500
	// M := 1000
	res := 0
	tmp := 0
	lenstr := len(s)

	for i := 0; i < lenstr; i++{
		
		if i == 0{
			switch{
			case s[i] == 'I' && s[i+1] != 'V' && s[i+1] != 'X':
					tmp = 1
			// case s[i] == 'V' :
			// 	tmp = 4
			case s[i] == 'V':
					tmp = 5
			case s[i] == 'X' && s[i+1] == 'L':
					tmp = 40
					i++
			case s[i] == 'X' && s[i+1] == 'C':
					tmp = 90
					i++
			case s[i] == 'X':
					tmp = 10
			case s[i] == 'L':
					tmp = 50
			case s[i] == 'C' && s[i+1] == 'D':
					tmp = 400
					i++
			case s[i] == 'C' && s[i+1] == 'M':
					tmp = 900
					i++
			case s[i] == 'C':
					tmp = 100
			case s[i] == 'D':
					tmp = 500
			case s[i] == 'M':
					tmp = 1000
			}
		}else if  i == lenstr-1{
			switch{
			case s[i] == 'I':
					tmp = 1
			case s[i] == 'V' && s[i-1] == 'I':
				tmp = 4
				i++
			case s[i] == 'V' && s[i-1] != 'I':
					tmp = 5
			case s[i] == 'X' && s[i-1] == 'I':
					tmp = 9
				i++
			case s[i] == 'X' && s[i-1] != 'I':
					tmp = 10
			case s[i] == 'L' && s[i-1] != 'X':
					tmp = 50
			case s[i] == 'C':
					tmp = 100
			case s[i] == 'D':
					tmp = 500
			case s[i] == 'M':
					tmp = 1000
			}

		}else {
			switch{
			case s[i] == 'I' && s[i+1] != 'V' && s[i+1] != 'X':
					tmp = 1
			case s[i] == 'V' && s[i-1] == 'I':
				tmp = 4
			case s[i] == 'V' && s[i-1] != 'I':
					tmp = 5
			case s[i] == 'X' && s[i-1] == 'I':
					tmp = 9
			case s[i] == 'X' && s[i-1] != 'I' && s[1+1] != 'L':
					tmp = 10
			case s[i] == 'L' && s[i-1] == 'X':
					tmp = 40
			case s[i] == 'L' && s[i-1] != 'x':
					tmp = 50
			case s[i] == 'C' && s[i-1] == 'X':
					tmp = 90
			case s[i] == 'C':
					tmp = 100
			case s[i] == 'D' && s[i-1] == 'X':
					tmp = 400
			case s[i] == 'D' && s[i-1] == 'C':
					tmp = 500
			case s[i] == 'M' && s[i-1] == 'X':
					tmp = 900
			case s[i] == 'M' && s[i-1] == 'C':
					tmp = 1000
			}
		}
		fmt.Println("tmp is ", tmp)
		
		res = res + tmp
		
	}
	return res


}


func main(){
	res := romanToInt("MCMXCIV")
	fmt.Println("result is ", res)
}