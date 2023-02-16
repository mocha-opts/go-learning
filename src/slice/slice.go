package main

import "fmt"

//var name []T

func main() {
	// 声明切片类型
	//var a []string              //声明一个字符串切片
  //var b = []int{}             //声明一个整型切片并初始化
	//var c = []bool{false, true} //声明一个布尔切片并初始化
	//var d = []bool{false, true} //声明一个布尔切片并初始化
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]  // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	
	//a[2:]  // 等同于 a[2:len(a)]
	//a[:3]  // 等同于 a[0:3]
	//a[:]   // 等同于 a[0:len(a)]
}
