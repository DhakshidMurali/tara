package toolRoutes

import "github.com/gin-gonic/gin"

func APIRoutes(server *gin.Engine) {
	toolsRoutes := server.Group("/")

	toolsRoutes.GET("/get/buildToolPage", BuildToolPage)

}
