package main

import (
	"fmt"
	"strings"
)

//Go 使用了特殊的 rune 类型来处理 Unicode
//rune类型，代表一个 UTF-8字符。
//rune类型实际是一个int32。

func main() {
	//字符串常规操作
	a1 := "啊哈哈哈"
	fmt.Println("len(a1) : ", len(a1)) //得9，因为得到得byte[]得长度
	//字符串拼接
	a2 := a1 + "卡卡卡"
	a3 := fmt.Sprintf("字符串: %v", a2)
	fmt.Println("a3---->", a3)
	//split
	split := strings.Split(a3, ":")
	fmt.Println("part-0 ", split[0], "  part-1", split[1])
	//contains
	contains := strings.Contains(a3, "卡卡")
	fmt.Println("contains? ", contains)
	//HashPrefix HasSuffix
	fmt.Println("hasPrefix? ", strings.HasPrefix(a3, "字"))
	fmt.Println("hasSuffix? ", strings.HasSuffix(a3, "看"))
	//index, //同len，都是返回的bytes的index
	fmt.Println("first index of '串' in a3：", strings.Index(a3, "串"))
	//join
	fmt.Println("join test: ", strings.Join([]string{"阿阿", "哈哈"}, ","))
	fmt.Println("----------------------------分割线--------------------------")
	/*
	   var a := '中' //字符a的类型是rune，代表一个UTF-8字符。rune实际就是一个int32
	   var b := 'x' //字符b的类型是byte，代表一个ASCII字符。byte实际是一个uint8
	*/
	var a = "kkk哈哈"
	//得9，因为得到得byte[]得长度。string默认底层是byte[]。not rune[]
	fmt.Println(len(a))
	for _, bValue := range a {
		//%c该值对应的unicode码值
		fmt.Printf("%v %c  ", bValue, bValue)
	}
	fmt.Println()

	//修改字符串
	var b = "啊波次的额佛格"
	r := []rune(b)
	r[0] = '哈' //单引号
	fmt.Println("rune test: len:, ", len(r), " value ", string(r))

}
