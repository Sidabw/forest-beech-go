package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	   1.实际在GO中，数组是不常用的，相比较于slices(切片)
	   2.数组大小不可变，但是slices可以。
	   3.slices有个头，有个尾，构成一个区间，并指向原始数组
	*/

	primes := [6]int{2, 3, 5, 7, 11, 13}
	//数组切割, 角标从0开始，包头不保尾
	//切割后不产生新的数据，即改变slice数组的内容，原始primes数组也会变
	var slice []int = primes[1:4]
	fmt.Println("primes[1: 4]: ", slice)

	//slice默认值，头的默认值是0，尾巴的默认值为arr.len
	//此时会从2角标位置一致切到最后，包含最后
	fmt.Println("省略尾巴：", primes[2:])
	fmt.Println("省略头部：", primes[:6])

	//Slice Literals 切片字面量（切片定义）
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("直接定义一个切片: ", q)

	q1 := q[1:]
	fmt.Println("在切片上‘切片’：", q1)

	q2 := q1[2:]
	fmt.Println("在切片上‘切片’-2:", q2)

	fmt.Println("len-cap test------------------------------------>")
	/*
	   1. 一个slice有两个属性，长度len 和 容量capacity
	   2. len是slice内元素的个数，capacity是该slice的下层数组的长度
	*/
	lc := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("原始数组: ", lc)
	lc = lc[:4]
	printSlices(lc)

	//这种切法，会直接把slice的下层数组的前两个元素给丢掉
	lc = lc[2:]
	printSlices(lc)

	fmt.Println("Nil slices test--------------------------------------->")
	var s []int
	//slice的0值就是nil
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil")
	}
	//index out of range
	//fmt.Println(s[0])

	fmt.Println("making slices test--------------------------------------->")
	a := make([]int, 5)
	fmt.Print("a: ")
	printSlices(a)

	fmt.Println("二维数组--------------------------------------->")
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Println("内嵌append方法测试--------------------------------------->")
	var ap []int
	printSlices(ap)
	ap = append(ap, 1)
	printSlices(ap)
	ap = append(ap, 2, 4, 5)
	printSlices(ap)

	fmt.Println("使用range关键字遍历slice or map--------------------------------------->")
	var pow = []int{1, 3, 5, 7}
	for i, v := range pow {
		//range的返回值有两个，分别是角标和对应元素的copy
		fmt.Printf("index：%d, corresponding value: %d \n", i, v)
	}
	//for _, value := range pow  简写，只要value
	//for i, _ := range pow 简写，只要index
	//for i := range pow    简写，只要index

}

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
}
