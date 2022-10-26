package routers

import (
	"assignment-3_AzmiFarisM/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	routers := gin.Default()
	routers.LoadHTMLGlob("templates/*.html")
	routers.GET("/", controllers.GetStatus)
	return routers
}
