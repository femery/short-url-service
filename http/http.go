package http

import (
	"2022/short-url-service/controller"
	"github.com/gin-gonic/gin"
)

//SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	router := gin.New()

	router.GET("/ping", func(context *gin.Context) {
		context.String(200, "Pong!")
	})
	router.POST("/surl/gen", controller.GenShortUrl)
	router.GET("/surl/v/:surl", controller.VisitShortUrl)
	return router

}
