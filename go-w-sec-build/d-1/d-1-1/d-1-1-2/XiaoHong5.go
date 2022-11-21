package d_1_1_2

import "fmt"

func HaFromD5() {
	//同一个包内的可以直接调用
	HaFromD2()
	fmt.Println("HaFromD5 invoked!")
}
