package main

import (
	"github.com/gin-gonic/gin"

	"github.com/sidabw/go-w-third-mod/router"
)

func main() {
	router.HaHaHa()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
