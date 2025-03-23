package domainRoutes

import "github.com/gin-gonic/gin"

func APIRoutes(server *gin.Engine) {
	domainRoutes := server.Group("/")

	domainRoutes.GET("/get/list/toolGroupByDomain", listToolGroupByDomain)

	domainRoutes.GET("get/list/toolsGroupByDeliveryFormatForTop4Domains", listToolsGroupByDeliveryFormatForTop4Domain)
}
