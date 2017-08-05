package main

import (
	"github.com/gin-gonic/gin"
	"wxapp/codec"
	"github.com/carsonsx/log4g"
	"user/token"
)

var DB = make(map[string]string)

func main() {

	ciphertext, err := codec.Encrypt("bbbbb")
	if err != nil {
		panic(err)
	}
	log4g.Debug(ciphertext)

	origData, err := codec.Decrypt(ciphertext)
	if err != nil {
		panic(err)
	}
	log4g.Debug(origData)

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	authorized := r.Group("/", token.VerifyTokenFn)

	authorized.POST("route/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "0"})
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8082")
}
