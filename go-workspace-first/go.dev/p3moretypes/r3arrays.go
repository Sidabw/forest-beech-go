package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])

	primes := [6]int{1, 10, 2, 4, 5, 6}
	fmt.Println(primes)

	//构建结构体数组
	primes3 := []struct {
		i int
		j bool
	}{
		{1, true},
		{2, false}, //最后这个,不能省略
	}
	fmt.Println(primes3)

}
