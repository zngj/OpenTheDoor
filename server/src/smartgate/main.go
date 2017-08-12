package main

import (
	"github.com/gin-gonic/gin"
	"usercenter/token"
	"smartgate/controller"
	"common/ginx"
)

var DB = make(map[string]string)

func main() {

	r := gin.New()
	r.Use(ginx.Logger(), ginx.Recovery())

	authorized := r.Group("/", token.VerifyTokenFn)
	authorized.GET("wallet/info", controller.WalletInfo)
	authorized.POST("wallet/charge", controller.WalletCharge)
	authorized.GET("router/status", controller.RouterStatus)
	authorized.GET("router/evidence", controller.RouterEvidence)
	authorized.GET("router/evidence/in", controller.RouterEvidenceIn)
	authorized.GET("router/evidence/out", controller.RouterEvidenceOut)
	authorized.GET("notification/:category", controller.GetRouterNotification)
	authorized.PUT("notification/consume/:id", controller.ConsumeRouterNotification)

	r.GET("/test/router/in/:userid", controller.TestRouterIn)
	r.GET("/test/router/out/:userid", controller.TestRouterOut)

	// Listen and Server in 0.0.0.0:8082
	r.Run(":8082")
}
