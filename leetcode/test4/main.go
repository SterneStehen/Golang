package main
import "fmt"


func romanToInt(s string) int {
    strlen := len(s)
	tmp := 0
	res := 0
	
	for i := 0; i < strlen; i++{
		
		if strlen > 1 && i != strlen -1 {
			if s[i] == 'I' && s[i+1] == 'V'{
				tmp = 4
				i++
				res = res + tmp
				continue
			}else if s[i] == 'I' && s[i+1] == 'X'{
				tmp = 9
				i++
				res = res + tmp
				continue
			}else if s[i] == 'X' && s[i+1] == 'L'{
				tmp = 40
				i++
				res = res + tmp
				continue
			}else if s[i] == 'X' && s[i+1] == 'C'{
				tmp = 90
				i++
				res = res + tmp
				continue
			}else if s[i] == 'C' && s[i+1] == 'D'{
				tmp = 400
				i++
				res = res + tmp
				continue
			}else if s[i] == 'C' && s[i+1] == 'M'{
				tmp = 900
				i++
				res = res + tmp
				continue
			}

		}

		if s[i] == 'I'{
			tmp = 1
		}else if s[i] == 'V'{
			tmp = 5
		}else if s[i] == 'X'{
			tmp = 10
		}else if s[i] == 'L'{
			tmp = 50
		}else if s[i] == 'C'{
			tmp = 100
		}else if s[i] == 'D'{
			tmp = 500
		}else if s[i] == 'M'{
			tmp = 1000
		}
		res = res + tmp
	}
	return res
}


func main(){
	res := romanToInt("MCMXCIV")
	fmt.Println("result is ", res)
}