package main

import (
	"fmt"
	"os"
)

func main() {
	//part-1 Print 系列
	//func Print(a ...interface{}) (n int, err error)
	//func Printf(format string, a ...interface{}) (n int, err error)
	//func Println(a ...interface{}) (n int, err error)

	//part-2 Fprint 系列
	//Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
	//r1-server.go中有用到
	fmt.Fprint(os.Stdout, "向标准输出写如内容")
	//0777 -rwxrwxrwx
	//0666 -rw-rw-rw-
	//0644 -rw-r--r--
	//三组参数由前往后，文件所有者权限，文件所属用户组权限，其他人权限
	fileObj, _ := os.OpenFile("/var/www/dream/webapp/logs/nasus2.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	fmt.Fprint(fileObj, "\nkkkk\n")

	//part-3 Sprint系列
	//Sprint系列函数会把传入的数据生成并返回一个字符串。
	name := "枯藤"
	age := 18
	s1 := fmt.Sprintf("name:%s,age:%d", name, age)
	fmt.Println(s1)

	//part-4 生成一个Error
	myerr := fmt.Errorf("自定义的一个错误")
	fmt.Printf("Type: %T, value: %v \n", myerr, myerr)

	//其他补充：格式化占位符
	//part-1 通用占位符
	//%v	值的默认格式表示
	//%+v	类似%v，但输出结构体时会添加字段名
	//%#v	值的Go语法表示/ 没看到跟%v有啥区别....
	//%T	打印值的类型
	//%%	百分号

	//part-2 布尔
	//%t	true或false

	//part-3 整型
	//%b	表示为二进制
	//%c	该值对应的unicode码值
	//%d	表示为十进制
	//%x	表示为十六进制，使用a-f
	//...

	//part-4 浮点数与复数
	//%b	无小数部分、二进制指数的科学计数法，如-123456p-78
	//%e	科学计数法，如-1234.456e+78
	//%E	科学计数法，如-1234.456E+78
	//%f	有小数部分但无指数部分，如123.456
	//%F	等价于%f
	//%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	//%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出

	//part-5 字符串和[]byte
	//%s	直接输出字符串或者[]byte
	//...

}
