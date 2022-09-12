package main

import "fmt"

/*
类型推导：
即不用写变量类型了，其实前面同时声明多个变量并赋初始值的例子(r4variables.go)，已经是在用类型推导了
*/
func main() {
	v1 := 1
	var v2 = 2
	fmt.Printf("v1 is of type %T \n", v1)
	fmt.Printf("v2 is of type %T \n", v2)

	f1 := 3.142
	//guess 0.5i 精度？
	var f2 = 0.876 + 0.5i
	fmt.Printf("f1 is of type %T \n", f1)
	fmt.Printf("f2 is of type %T \n", f2)

}
