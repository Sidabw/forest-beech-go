package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// 这里使用的是闭包函数
	e.GET("/hello", func(context echo.Context) error {
		return context.String(http.StatusCreated, "echo framework success.")
	})
	// 启动http server, 并监听8080端口，冒号（:）前面为空的意思就是绑定网卡所有Ip地址，本机支持的所有ip地址都可以访问。
	e.Start(":8081")
}
