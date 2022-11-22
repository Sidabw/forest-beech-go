package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//https://blog.csdn.net/lvjie13450/article/details/121855625

	//readAllInBasics()
	//readAllInBuffer()
	readAllInIoUtil()
}

func readAllInBasics() {
	readmeFile, err := os.Open("/var/www/dream/webapp/logs/nasus.log")
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer readmeFile.Close()

	for true {
		tmpBytes := make([]byte, 128)
		//n为读到了多少字节
		n, err := readmeFile.Read(tmpBytes)
		if n > 0 {
			fmt.Println("读到内容：", string(tmpBytes[:n]))
		}
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("读取文件失败", err)
			return
		}
	}
}

func readAllInBuffer() {
	readmeFile, err := os.Open("/var/www/dream/webapp/logs/nasus.log")
	if err != nil {
		fmt.Println("打开文件失败，", err)
		return
	}
	defer readmeFile.Close()

	for true {
		bufioReader := bufio.NewReader(readmeFile)
		//以\n作为终止符(delimit)，读字符串
		//这里是要一个byte
		readString, err := bufioReader.ReadString('\n')
		fmt.Println("读到字符串：", readString)
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
	}
}

func readAllInIoUtil() {
	file, err := ioutil.ReadFile("/var/www/dream/webapp/logs/nasus.log")
	if err != nil {
		fmt.Println("打开文件失败，", err)
		return
	}
	if err == io.EOF {
		fmt.Println("empty file")
		return
	}
	fmt.Println(string(file))
}
