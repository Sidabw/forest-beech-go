package main

import (
	"fmt"
	"w7_widgets/amqp"
	"w7_widgets/mem"
)

func main() {

	mem.DoTest()
	fmt.Println("------------------------------------------")
	fmt.Println("------------------------------------------")
	fmt.Println("------------------------------------------")

	amqp.DoTest2()

}
