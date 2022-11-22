package main

import (
	"github.com/labstack/echo"

	"github.com/sidabw/go-w6-echo/router"
)

func main() {
	e := echo.New()
	// 这里使用的是闭包函数
	e.GET("/hello", router.Test1HelloWorld)
	// curl 'localhost:8081/h2/aa?a=b'
	e.GET("/h2/:id", router.Test2GetParams)

	// Test3PostForm
	// Test4PostJson

	// localhost:8081/static/grassland.png
	e.Static("/static", "static")
	// 启动http server, 并监听8080端口，冒号（:）前面为空的意思就是绑定网卡所有Ip地址，本机支持的所有ip地址都可以访问。
	e.Start(":8081")
}
