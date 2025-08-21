package main
import "fmt"

func isValid(s string) bool {

	stack :=[]byte{}

	for i := 0; i < len(s); i++{
		if s[i] == '{' || s[i] == '(' || s[i] == '['{
			stack = append(stack, s[i])
		} else { 
			if len(stack) == 0 {
				return false
			}	else if s[i] == '}' && stack[len(stack)-1] != '{' {
					return false
			}	else if s[i] == ']' && stack[len(stack)-1] != '[' {
					return false
			}	else if s[i] == ')' && stack[len(stack)-1] != '(' {
					return false
			}
			stack = stack[:len(stack) -1]
		}
		
	}
	if(len(stack) == 0){
		return true
	} else{
		return false
	}
}

func main(){
	res := isValid("()")
	fmt.Println(res)
}

