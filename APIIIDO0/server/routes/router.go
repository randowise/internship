package routes

import (
	"example/APIIIDO0/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		pro := main.Group("pro")
		{
			pro.POST("/", controllers.CreatePro)
			pro.GET("/:id", controllers.ShowPro)
			pro.GET("/", controllers.ShowPros)

		}
	}
	return router
}
