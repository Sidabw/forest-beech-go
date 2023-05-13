package main

import (
	"fmt"
	"math/cmplx"
)

//变量声明可以分组作为一个语法块
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	//%T 取到值类型
	//%t 布尔值，
	//前面的 %g 浮点数
	//%v 值的默认格式表示
	//%d 十进制整形 -> r4slices
	//fmt.Printf参考：[fmt.Printf guide](https://blog.csdn.net/weixin_44706011/article/details/106483136)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

/*
Go基本类型：




*/
