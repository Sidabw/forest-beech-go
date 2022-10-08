package main

import "fmt"

func main() {
	deferTest1()
	fmt.Println("----------分割线")
	deferTest2()
}

func deferTest1() {
	//dt() 函数会立即执行，但当前defer修饰的函数，则会在其包裹函数return后执行
	defer fmt.Println(dt())

	fmt.Println("hello")
}

func deferTest2() {
	//defer-multi
	fmt.Println("Counting")

	//defer的语句会被压到栈里，所以先进后出，最后打印的顺序是倒着的
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done.")
}

func dt() string {
	fmt.Println("dt method invoked")
	return "world"
}
