package main

import "fmt"

func main() {
  //dt() 函数会立即执行，但当前defer修饰的函数，则会在其包裹函数return后执行
  defer fmt.Println(dt())
  
  fmt.Println("hello")
}

func dt() string {
  fmt.Println("dt method invoked")
  return "world"
}
