package main

import (
	"github.com/gin-gonic/gin"
	"usercenter/controller"
	"usercenter/token"
)

var DB = make(map[string]string)

func main() {

	r := gin.Default()

	// Get user value
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to youstars' user center api service")
	})

	r.POST("/wxapp/login", controller.WxappLogin)
	r.GET("/verifytoken", token.VerifyToken)

	// Listen and Server in 0.0.0.0:8081
	r.Run(":8081")
}