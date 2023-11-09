package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {

		//获取query string
		//name := context.DefaultQuery("name", "hhh")
		//fmt.Printf("get name from qs: %v when %v \n", name, time.Now())

		context.String(http.StatusOK, "hello world..")

	})
	r.Run(":8011")
}
