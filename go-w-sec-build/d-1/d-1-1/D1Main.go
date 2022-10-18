package main

import (
	"go-w-sec-build/d-1/d-1-1/D-1-1-1"
)

func main() {
	//D1Main.go -> d-1-1.XiaoHong2 -> d-1-1.XiaoHong1
	//func main 必须在main包下
	//外部可调用的，即是可导出的
	//可导出的，函数名/包名首字母都必须大写
	D_1_1_1.HaFromD2()
}
