package main

import (
	"common/ginx"
	"github.com/gin-gonic/gin"
	"usercenter/controller"
)

var DB = make(map[string]string)

func main() {

	r := gin.New()
	r.Use(ginx.Logger(), ginx.Recovery())

	// Get user value
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to youstars' user center api service")
	})

	r.POST("/wxapp/login", controller.WeappLogin) // deprecated
	r.GET("/verifytoken", controller.CheckToken)  // deprecated

	r.GET("/check_phone_number", controller.CheckPhoneNumber)
	r.POST("/signup", controller.SignUp)
	r.POST("/login", controller.Login)
	r.POST("/login_weapp", controller.WeappLogin)
	r.POST("/login_wxapi", controller.WxapiLogin)
	r.GET("/check_token", controller.CheckToken)

	// Listen and Server in 0.0.0.0:8081
	r.Run(":8081")
}
