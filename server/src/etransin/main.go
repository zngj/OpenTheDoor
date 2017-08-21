package main

import (
	"common/ginx"
	"github.com/gin-gonic/gin"
	"etransin/controller"
	"common/tokenutil"
)

var DB = make(map[string]string)

func main() {

	r := gin.New()
	r.Use(ginx.Logger(), ginx.Recovery())

	authorized := r.Group("/", tokenutil.VerifyToken)

	wallet := authorized.Group("wallet")
	{
		wallet.GET("/info", controller.WalletInfo)
		wallet.POST("/charge", controller.WalletCharge)
	}

	evidence := authorized.Group("evidence")
	{
		evidence.GET("/in", controller.GetEvidenceIn)
		evidence.GET("/out", controller.GetEvidenceOut)
		evidence.PUT("/discard/:id", controller.DiscardEvidence)
		//evidence.PUT("/consume/:id", controller.DiscardEvidence)
	}

	router := authorized.Group("router")
	{
		router.GET("/status", controller.RouterStatus)
		router.GET("/evidence/in", controller.GetEvidenceIn) //deprecate
		router.GET("/evidence/out", controller.GetEvidenceOut)
		router.GET("/in/list", controller.RouterInList)
		router.GET("/out/list", controller.RouterOutList)
		router.GET("/list", controller.MyRouters)
	}

	notification := authorized.Group("notification")
	{
		notification.GET("/current", controller.CurrentNotification)
		notification.PUT("/consume", controller.ConsumeRouterNotification)
	}

	test := r.Group("/test")
	{
		test.GET("/router/in/:userid", controller.TestRouterIn)
		test.GET("/router/out/:userid", controller.TestRouterOut)
	}

	// Listen and Server in 0.0.0.0:8082
	r.Run(":8082")
}
