package main

import "fmt"

func typeSwitchTest(a interface{}) {
	//这个相比Java省了挺多代码的...
	switch b := a.(type) {
	case int:
		//不用强转了
		fmt.Println("int coming in", b*2)
	case string:
		fmt.Println("string coming in", len(b))
	default:
		fmt.Println("default in ...")
	}
}

func main() {
	typeSwitchTest(1)
	typeSwitchTest("aaaaaa")
	typeSwitchTest(true)
}
