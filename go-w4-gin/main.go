package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {

		//获取query string
		name := context.DefaultQuery("name", "hhh")
		fmt.Println("get name from qs: ", name)

		context.String(http.StatusOK, "hello world")

	})
	r.Run(":8011")
}
