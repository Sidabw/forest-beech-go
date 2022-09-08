package main

import "fmt"

func main() {
  fmt.Println("Counting")
  
  //defer的语句会被压到栈里，所以先进后出，最后打印的顺序是倒着的
  for i := 0; i < 10; i ++ {
    defer fmt.Println(i)
  }

  fmt.Println("Done.")
}

