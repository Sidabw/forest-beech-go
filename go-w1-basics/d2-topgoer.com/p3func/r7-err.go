package main

import "fmt"

func testPanic() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("catch a panic. panic info:: ", e)
		}
	}()

	panic("throw a panic")

	fmt.Println("Am i invoked 1?")
}

func main() {
	//https://blog.csdn.net/qq_39458487/article/details/124798684
	//https://www.topgoer.com/%E5%87%BD%E6%95%B0/%E5%BC%82%E5%B8%B8%E5%A4%84%E7%90%86.html

	//Go 语言没有异常系统，其使用 panic 触发宕机类似于其他语言的抛出异常， recover 的宕机恢复机制就对应其它语言中的 try/catch 机制。

	testPanic()
	fmt.Println("Am i invoked 2?")
}
