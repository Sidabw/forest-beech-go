package main

import "fmt"

//iota是go语言的常量计数器，只能在常量的表达式中使用。
//iota在const关键字出现时将被重置为0
//const中每新增一行常量声明将使iota计数一次
//使用iota能简化定义，在定义枚举时很有用
const (
	n1  = iota //0
	n2         //1
	n3         //2
	n4         //3
	_          //占位
	n5         //5
	n6  = 100  //插队后重新拿回iota
	n7  = iota //7
	n8  = 200  //插队后不要iota了
	n9         //200
	n10        //200
)

//多个iota定义在一行，这个要复杂些...
const (
	a, b = iota + 2, iota + 5
	a1, b1
	a2, b2
)

func main() {
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n7)
	fmt.Println(n9)
	fmt.Println(n10)
	fmt.Println("-----------------------------------------------")

	//意思就是iota每行+1
	//第一行的算数表达式也会'继承下来'
	//按列继承啊...
	fmt.Println(a, b)
	fmt.Println(a1, b1)
	fmt.Println(a2, b2)
}
