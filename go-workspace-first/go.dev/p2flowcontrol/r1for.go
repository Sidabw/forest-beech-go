package main

import "fmt"

func main() {
	var sum = 0

	// 1. 标准写法
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//2. 省略初始语句、后执行语句
	for sum < 1000 {
		sum += 1
	}
	fmt.Println(sum)

	//3. 同上,继续省略
	for sum < 2000 {
		sum += 2
	}
	fmt.Println(sum)

	//4. forever loop

	/* for {
	   fmt.Println(1)
	 }*/
	//TODO where is goto? break? continue?
}

/*
1. Go has only one looping construct, the for loop.
2. the variables declared there are visible only in the scope of the for statement
*/
