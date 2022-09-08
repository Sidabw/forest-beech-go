package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

/*
1. GO没有类，但是可以通过这种方式给“给定Type”定义个方法
*/

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


/*
2. 第二种定义method的方式，定义一个type，不需要struct。有人会这么用吗？包括上面的，直接定义一个function不香吗
官方的说法同样是：methods is function
*/
type MyFloat float64

func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

/*
3.对于第一种方式，使用的是值传递，所以当我们想在方法中改变struct的字段的值时，需要使用引用传递，如下：
*/
func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}


func main(){
  v := Vertex{2, 4}
  fmt.Println(v.Abs())


  f := MyFloat(-math.Sqrt2)
  fmt.Println(f.Abs())

  v2 := Vertex{5, 8}
  v2.Scale(10)
  fmt.Println(v2.X, v2.Y)
}

