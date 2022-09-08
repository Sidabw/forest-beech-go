package main

import "fmt"

//类型在形参后面
func add(x int, y int) int {
  return x + y
}

//连续两个或多个相同时，可以简写
func add2(x, y int) int {
  return x + y
}

//返回值可以有多个
func swap(x, y string) (string, string) {
  return y, x
}

/*
返回值可以被命名：可以理解被命名的返回值就是定义在函数顶部的变量，用法就是‘简写return’;
不过不推荐使用，除非函数很短
*/
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

func main() {
  fmt.Println(add(42, 11))
  fmt.Println(add2(32, 11))
  
  a, b := swap("hello", "world")
  fmt.Println(a, b)

  fmt.Println(split(17))

}

