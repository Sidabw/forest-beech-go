package main

import "fmt"

//Go 使用了特殊的 rune 类型来处理 Unicode
//rune类型，代表一个 UTF-8字符。
//rune类型实际是一个int32。

func main() {
	var a = "kkk哈哈"
	fmt.Println(len(a)) //得9，因为得到得byte[]得长度
	for _, bValue := range a {
		fmt.Printf("%v %c  ", bValue, bValue)
		//fmt.Println(bIndex, bValue)
	}
	fmt.Println()

	//修改字符串
	var b = "啊波次的额佛格"
	r := []rune(b)
	r[0] = '哈' //单引号
	fmt.Println(string(r))

}
