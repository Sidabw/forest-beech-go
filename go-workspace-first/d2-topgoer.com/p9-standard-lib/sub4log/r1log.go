package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//part-1，初识log
	//part1()
	//part2()
	part3()
	//....
	//Go内置的log库功能有限，例如无法满足记录不同级别日志的情况，我们在实际的项目中根据自己的需要选择使用第三方的日志库，如logrus、zap等
	//....

}

func part1() {
	log.Println("这是一条很普通的日志。") //2022/10/20 14:22:04 这是一条很普通的日志。
	log.Printf("这是一条%s日志。\n", " \"很普通的\"")
	log.Fatalln("这是一条会触发fatal的日志。") //会调用os.Exit(1)
	log.Panicln("这是一条会触发panic的日志。") //会出发panic
}

func part2() {
	//设置一些标签，比较有用的就Lshortfile，打出来文件和行号
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	//设置前缀
	log.SetPrefix("[pprof]")
	log.Println("这是一条很普通的日志。")
}

func part3() {
	logFile, err := os.OpenFile("/var/www/dream/webapp/logs/nasus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	//只往文件里打了。。。
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	log.Println("哈哈哈哈哈哈哈哈哈凯凯凯凯")
}
