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

	//5. break? continue? in for and label
Label1:
	for sum < 3000 {
		sum += 2
		if sum == 2500 {
			fmt.Println("break for loop")
			//continue同理不再演示
			break Label1
		}
	}

	for sum < 4000 {
		sum += 1
		if sum == 3333 {
			fmt.Println("break for loop using a simple break")
			break
		}
	}

	//6. goto?
	//goto用于在判断语句中跳到指定行。不推荐使用。
	if sum == 2500 {
		goto Label2
	}
	fmt.Println("1 Has this proceeded ?")
	fmt.Println("2 Has this proceeded ?")
	fmt.Println("3 Has this proceeded ?")
Label2:
	fmt.Println("This has proceeded.")
}

/*
1. Go has only one looping construct, the for loop.
2. same as Java: the variables declared there are visible only in the scope of the for statement
*/
