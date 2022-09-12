package main

import "fmt"
import "math"

/* 也可以写成分组导入的形式。分组导入是推荐的
import (
  "fmt"
  "math"
)
*/

func main() {
  fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
  fmt.Println(math.Pi)
  
  /* 
  引用一个包时，只能引用其‘已导出’的名字；如果一个名字已大写开头那它就是已导出的；
  类似math.Pi    fmt.Printf
  */

}


