package main

import "fmt"


func modifyArray(x [3]int)[3]int {
	x[1]=100
	return x
}

func modifyArray2(x [3][2]int)  {
	x[2][0] = 100
	return
}

func main()  {
	a := [3]int{10,20,30}
	b := modifyArray(a)
	fmt.Println(b)
}