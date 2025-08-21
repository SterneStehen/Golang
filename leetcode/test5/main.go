package main
import "fmt"

func longestCommonPrefix(strs []string) string {
	
	if len(strs[0])== 0{
			return ""
	}
	count := len(strs[0])
	lenstrs := len(strs)
	
	
    for i := 1; i < lenstrs; i++{
		tmpc := 0
		lensubstr := len(strs[i])
		if lensubstr == 0 {
			return ""
		}
		if lenstrs> 1 && i < lenstrs {
			k :=0
			for k < len(strs[0]) && k < len(strs[i]) && strs[0][k] == strs[i][k]{
				tmpc += 1					
				k++
				
			}
			
		}
		if tmpc < count{
			count = tmpc
		}
		if count == 0{
			return ""
		}
					//fmt.Printf("%d", count)
		
		
	}

	retsetr := string(strs[0][0:count])
	return  retsetr
	
	
}

func main(){
	strs := []string{"fl", "flooo"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
}