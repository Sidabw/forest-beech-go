package main

import (
  "fmt"
  "math"
)

//可以传递函数
func compute(fn func(float64, float64) float64) float64 {
  return fn(3, 4)
}


//可以返回一个函数；下面是创建一个闭包(closures)
func adder() func(int) int {
  //这个sum属于闭包对象里
  sum := 0
  return func(x int) int {
    sum += x
    return sum
  }
}

func main() {
  hypot := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
  }
  fmt.Println(hypot(5, 12))

  fmt.Println(compute(hypot))
  fmt.Println(compute(math.Pow))


  fmt.Println("闭包closures test---------------------------------->")
//  pos, neg := adder(), adder()
  pos := adder()
  for i := 0; i < 10; i++ {
    fmt.Println(pos(i))
  }
}
