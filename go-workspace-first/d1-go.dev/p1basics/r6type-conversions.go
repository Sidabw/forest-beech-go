package main

import (
  "fmt"
  "math"
)

func main() {
  var x, y int = 3, 4
  //需要显式的做类型转换，直接转会抛异常
  var f float64 = math.Sqrt(float64(x*x + y*y))
  //强转得到一个无符号int。
  //当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型
  var z uint = uint(f)
  fmt.Println(x, y, f, z)
}

