package main

import (
  "fmt"
  "math"
)

func main() {
  var x, y int = 3, 4
  //需要显式的做类型转换，直接转会抛异常
  var f float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(f)
  fmt.Println(x, y, f, z)
}
