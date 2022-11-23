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

	// curl --location --request POST 'localhost:8081/h3' --header 'Content-Type: application/x-www-form-urlencoded' --data-urlencode 'a=哈哈'
	// --location 处理重定向
	e.POST("/h3", router.Test3PostForm)
	// curl --request POST 'localhost:8081/h4' --header 'Content-Type: application/json' --data-raw '{ "id": 11112, "name":"是布鲁啊" }'
	e.POST("/h4", router.Test4PostJson)

	e.GET("/h5", router.Test5Download)
	e.POST("/h6", router.Test6Upload)

	// localhost:8081/static/grassland.png
	e.Static("/static", "static")
	// 启动http server, 并监听8080端口，冒号（:）前面为空的意思就是绑定网卡所有Ip地址，本机支持的所有ip地址都可以访问。
	e.Start(":8081")

	//cookie & session
	//https://www.cnblogs.com/remixnameless/p/14318225.html
	//https://www.cnblogs.com/remixnameless/p/14318227.html

	//中间件 middleware
	//https://www.cnblogs.com/remixnameless/p/14318287.html

}
