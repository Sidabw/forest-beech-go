package main

import (
	"fmt"
	"strconv"
)

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func main() {
	//parseInt...
	atoi, err := strconv.Atoi("123")
	if err == nil {
		fmt.Println("haha,没有error", atoi)
	}
	fmt.Println(err)

	sqrt, err := Sqrt(1.2)
	if err != nil {
		return
	}
	fmt.Println(sqrt)
}
