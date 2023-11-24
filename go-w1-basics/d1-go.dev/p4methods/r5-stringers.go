package main

import "fmt"

type IpAddr [4]byte

// 这个method的父类是Stringer
// 有点多态的味道了哈...
// d1-go.dev/tour/methods/19， r6-err，
// func (e *MyError) Error() string 一样的
// 实际是type error interface { Error() string } 的实现
func (dd IpAddr) String() string {
	var a string = ""
	for i := range dd {
		//byte 转 string..
		a2 := fmt.Sprintf("%v", dd[i])
		a = a + a2
	}
	return a
}

func main() {
	var a = IpAddr([4]byte{1, 0, 0, 1})
	fmt.Println(a)
}
