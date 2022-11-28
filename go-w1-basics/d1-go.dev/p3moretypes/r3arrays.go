package main

import "fmt"

func main() {
	//1. 构建一个字符串数组
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])

	//2. 构建一个整形数组
	primes := [6]int{1, 10, 2, 4, 5, 6}
	fmt.Println(primes)

	//3. 构建一个结构体数组
	//这应该是个切片吧
	primes3 := []struct {
		i int
		j bool
	}{
		{1, true},
		{2, false}, //最后这个,不能省略
	}
	fmt.Println(primes3)

	var primes4 = []Vertex{
		{1, 2},
		{2, 3},
	}
	fmt.Println(primes4[0])

}
