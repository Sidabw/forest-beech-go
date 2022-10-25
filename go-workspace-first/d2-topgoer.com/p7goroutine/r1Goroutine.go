package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//goroutine就是协程
	//协程，又叫微线程，调度仅在用户态完成，类似Java的NIO。
	//Java的线程在用户态和内核态是一对一对应的，肯定是比协程慢。
	//https://cloud.tencent.com/developer/article/1839604
	//https://mp.weixin.qq.com/s?__biz=MzUyNzgyNzAwNg==&mid=2247484266&idx=1&sn=203fa240c91642b7eef11d3f34374526&scene=21#wechat_redirect

	//learn1()
	learn2()

}

func learn1() {
	//创建一个协程，相比与Java的NIO，使用难度上不是一个层次啊。。。
	go hello()
	fmt.Printf("hello2 \n")
	time.Sleep(time.Second * 5)
}

func hello() {
	fmt.Printf("hello1 \n")
}

//原来struct还可以这样用？ 都不用构建啊....
var wg sync.WaitGroup

func learn2() {
	//类似Java的CountDownLaunch，主协程等子协程全部执行玩了再执行
	for i := 0; i < 5; i++ {
		wg.Add(1)
		hello2(i)
	}
	wg.Wait()
}

func hello2(i int) {
	defer wg.Done()
	fmt.Println("hello + ", i)
}
