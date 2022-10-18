package main

import (
	"fmt"
	"math"
)

/* 也可以写成分组导入的形式。分组导入是推荐的
import (
  "fmt"
  "math"
)
*/

/*
import f "fmt"，为标准库起一个别名，而后调用 fmt.Println("") 可以使用 f.Println("")
import . "fmt"，将 fmt 启用别名"."，这样就可以直接使用其内容，而不用再加 fmt，例如fmt.Println("") 可以直接写成 Println("")
import _ "samples/util"，表示不使用该包，而是只执行下该包的 init 函数

*/

func main() {
	//%g 浮点数
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
	fmt.Println(math.Pi)

	/*
	  引用一个包时，只能引用其‘已导出’的名字；如果一个名字已大写开头那它就是已导出的；
	  类似math.Pi    fmt.Printf
	*/

}
