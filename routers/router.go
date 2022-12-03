package routers

import (
	"FoT/controllers"
	"github.com/gin-gonic/gin"
)



func registerRouter() *gin.Engine {
	//config := initConfig()
	//fmt.Println(config.Get("settings.application.host"))
	//host := config.Get("settings.application.host")
	//port := config.Get("settings.application.port")

	router := gin.Default()

	menuRouter := router.Group("/menu")
	menuRouter.GET("/", controllers.QueryMenu)
	menuRouter.POST("/", controllers.CreateMenu)

	return router

	//server.GET("/api", func(context *gin.Context) {
	//	context.JSON(200, gin.H{"msg": "api创建成功"})
	//})
	//err := server.Run(host)
	//if err != nil {
	//	fmt.Println("服务器启动失败！")
	//}
}
