package main

import "fmt"

//var用于声明一个变量列表,类型在最后面，不赋值的话走默认值，即 0 false ""
//var可用于包级别、函数级别
var c, python, java bool

var i2, i3 = 1, 2

func main() {
	var i int
	fmt.Println(i, c, python, java)

	var j1, j2, j3 = true, false, "no!"
	fmt.Println(i2, i3, j1, j2, j3)

	//简洁赋值语句 := 可以替代var，但只能在函数内使用
	k := 1
	k1, k2, k3 := true, true, "yes!"
	fmt.Println(k, k1, k2, k3)
}
