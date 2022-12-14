package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	//1. standard if statement
	if x < 0 {
		//这给你厉害的，还返回一个虚数...
		return sqrt(-x) + "i"
	}
	//Sprint系列函数会把传入的数据生成并返回一个字符串。
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	//2. a short statement before if
	//pow 返回x的y次幂
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g \n", v, lim)
	}
	return lim
}

func main() {
	//sa
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	//3. switch statement
	// 也可以写成 switch os:= runtime.GOOS; os {...} 即switch也可以又一个赋值语句
	// 每个case都是默认带着break的
	os := runtime.GOOS
	fmt.Print("Go runs on : ")
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s. \n", os)
	}

	//4. switch statement-2
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too Far away")
	}

	//4. switch statement-3 : if   else if    else if
	//这里switch 后是不跟判断条件的
	//因为go只有if else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
		fmt.Println("Good Good.")
	default:
		fmt.Println("Good evening.")
	}
}
