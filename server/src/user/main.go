package main

import (
	"github.com/gin-gonic/gin"
	"user/controller"
	"user/token"
)

var DB = make(map[string]string)

func main() {

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Get user value
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to youstars' user center api service")
	})

	// Get user value
	r.POST("/wxapp/login", controller.WxappLogin)

	// Get user value
	r.POST("/verifytoken", token.VerifyToken)


	// Listen and Server in 0.0.0.0:8080
	r.Run(":8081")
}
