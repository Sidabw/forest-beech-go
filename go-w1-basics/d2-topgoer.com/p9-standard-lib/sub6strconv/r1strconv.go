package main

import (
	"fmt"
	"strconv"
)

func main() {
	//strconv包实现了基本数据类型与其字符串表示的转换，
	//主要有以下常用函数： Atoi()、Itia()、parse系列、format系列、append系列。

	//part-1, Atoi
	//a说的是array，在c语言中没有字符串，只有array/
	//i自然int
	fmt.Println("part-1------------------------------------")
	i1, _ := strconv.Atoi("1111")
	fmt.Printf("type: %T, value: %#v \n", i1, i1)

	fmt.Println("part-2------------------------------------")
	//part-2, Itoa
	res2 := strconv.Itoa(11111)
	fmt.Printf("type: %T, value: %s \n", res2, res2)

	fmt.Println("part-3------------------------------------")
	//part-3, 字符串转bool、转int、转uint、转float
	res3, _ := strconv.ParseBool("true")
	fmt.Printf("type: %T, value:%v \n", res3, res3)
	//base 表示的是进制（2～36）
	//bitSize 表示整数类型（0、8、16、32、64）
	res4, _ := strconv.ParseInt("-2", 10, 64)
	fmt.Printf("type: %T, value: %v \n", res4, res4)
	res5, _ := strconv.ParseFloat("1.2", 64)
	fmt.Printf("type: %T, value: %v \n", res5, res5)

	fmt.Println("part-4------------------------------------")
	//part-4，
	strRes1 := strconv.FormatBool(true)
	fmt.Printf("type: %T, value: %s", strRes1, strRes1)
	strRes2 := strconv.FormatInt(234333, 10)
	fmt.Printf("type: %T, value: %s \n", strRes2, strRes2)
	//FormatFloat有点麻烦..

}
