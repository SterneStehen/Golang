package main

import "fmt"
//import "const rain ts"


type addAny interface{
	~int| ~float64 | ~string
}

func addSomeAny[T addAny](a, b T) T{
	return a + b
}

func addSome(a, b interface{}) interface{}  {
	aInt, aOk := a.(int)
	bInt, bOk := b.(int)
	
	if aOk && bOk{
		return aInt+bInt
	} 

	aFloat, aOk := a.(float64)
	bFloat, bOk := b.(float64)
	if aOk && bOk{
		return aFloat + bFloat
	}

	aStr, aOk := a.(string)
	bStr, bOk := b.(string)

	if aOk && bOk{
		return aStr + " " + bStr
	}

	return nil
}

func main(){
	a := 3
	b := 4
	fmt.Println(addSome(a, b))
	c := "serg"
	d := "mor"
	fmt.Println(addSome(c, d))
	e := 3
	f := 4
	fmt.Println(addSome(e, f))
	k := "serg"
	l := "mor"
	fmt.Println(addSome(k, l))
}