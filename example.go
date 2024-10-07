package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/welcome", func(c *gin.Context) {
		c.String(http.StatusOK, "值:%v", "welcome to gin web service!")
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080

	//var a = 1
	//var b = 2
	//
	//c := a + b
	//fmt.Println(c)

}
