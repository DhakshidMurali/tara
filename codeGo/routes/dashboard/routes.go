package dashboardRoutes

import "github.com/gin-gonic/gin"

func APIRoutes(server *gin.Engine) {
	authRoutes := server.Group("/")

	authRoutes.GET("/get/list/toolGroupByDomain", listToolGroupByDomain)
	authRoutes.GET("get/list/toolsGroupByDeliveryFormatForTop4Domains", listToolsGroupByDeliveryFormatForTop4Domain)
}
