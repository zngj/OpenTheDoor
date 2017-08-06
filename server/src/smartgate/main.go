package main

import (
	"github.com/gin-gonic/gin"
	"usercenter/token"
	"smartgate/controller"
)

var DB = make(map[string]string)

func main() {

	r := gin.Default()

	authorized := r.Group("/", token.VerifyTokenFn)
	authorized.GET("wallet/info", controller.WalletInfo)
	authorized.GET("router/status", controller.RouterStatus)
	authorized.GET("router/evidence", controller.RouterEvidence)

	// Listen and Server in 0.0.0.0:8082
	r.Run(":8082")
}
