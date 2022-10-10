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
	//定义一个复数，有关复数的更多内容，参考，印象笔记，《数学之美-杂录》
	//总结一句话：Java怎么表示一个有理数（无限循环小数）、怎么表示一个无理数（math.Pi）、怎么表示一个虚数？
	//尴尬的就是Java没有复数这个概念...
	var f2 = 0.876 + 0.5i
	fmt.Printf("f1 is of type %T \n", f1)
	fmt.Printf("f2 is of type %T , for value: %v \n", f2, f2)

}
