package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//https://blog.csdn.net/lvjie13450/article/details/121855625

	//writeInBasics()
	//writeInBuffer()
	writeInIoUtil()
}

func writeInBasics() {
	//os.O_CREATE:创建
	//os.O_WRONLY:只写
	//os.O_APPEND:追加
	//os.O_RDONLY:只读
	//os.O_RDWR:读写
	//os.O_TRUNC:清空

	//0644:文件的权限
	nasusFile, err := os.OpenFile("/var/www/dream/webapp/logs/nasus2.log",
		os.O_TRUNC|os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer nasusFile.Close()

	nasusFile.Write([]byte("hello 啊 "))
	nasusFile.WriteString("my go 世界")
	nasusFile.Write([]byte{'\n'})
}

func writeInBuffer() {
	nasusFile, err := os.OpenFile("/var/www/dream/webapp/logs/nasus2.log",
		os.O_TRUNC|os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer nasusFile.Close()

	nasusWriter := bufio.NewWriter(nasusFile)
	nasusWriter.WriteString("hahahha哈哈哈\n")
	nasusWriter.Flush()
}

func writeInIoUtil() {
	str := "啊哈哈"
	err := ioutil.WriteFile("/var/www/dream/webapp/logs/nasus2.log", []byte(str), 0644)
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
}
