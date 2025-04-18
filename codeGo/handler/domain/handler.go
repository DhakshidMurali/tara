package domainRoutes

import (
	domainService "github.com/DhakshidMurali/tara/service/domain"
	"github.com/gin-gonic/gin"
)

func APIRoutes(server *gin.Engine) {

	domainRoutes := server.Group("/")

	domainRoutes.GET("get/list/toolGroupByDomain", domainService.ListToolGroupByDomain)

	domainRoutes.GET("get/list/toolsGroupByDeliveryFormatForTopDomains", domainService.ListToolsGroupByDeliveryFormatForTop5Domain)

	// Other handler used for Domain Page Routes
	// authHandler.POST("/get/list/domain", service.ListDomain)
}
