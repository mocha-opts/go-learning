package main

import (
	"fmt"
)

func main()  {
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	
	n := 10;
	m := 200
	fmt.Println("Hello World!",n,m)
}
func foo()(int,string){
	return 10, "hh"
}
//匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。 
//函数外的每个语句都必须以关键字开始（var、const、func等）
//:=不能使用在函数外。
//_多用于占位，表示忽略值。