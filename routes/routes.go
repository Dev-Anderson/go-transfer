package routes

import (
	"go-transfer/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		home := main.Group("home")
		{
			home.GET("/", controllers.Home)
		}
		upload := main.Group("upload")
		{
			upload.POST("/", controllers.PostUpload)
		}
	}

	return router
}
