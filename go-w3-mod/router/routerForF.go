package router

import (
	"fmt"

	myService "github.com/sidabw/go-w-third-mod/service"
)

func HaHaHa() {
	fmt.Println("HaHaHa from go-w3-mod")

	//函数调用这样没问题
	//如果要用service包下面的type xx struct的话，对应文件名也得大写
	myService.HeHeHe()
}
