package domainRoutes

import (
	domainService "github.com/DhakshidMurali/tara/service/domain"
	"github.com/gin-gonic/gin"
)

func APIRoutes(server *gin.Engine) {
	domainRoutes := server.Group("/")

	domainRoutes.GET("/get/domainpage/details", domainService.GetDomainPageDetails)

	domainRoutes.GET("/get/list/toolGroupByDomain", domainService.ListToolGroupByDomain)

	domainRoutes.GET("get/list/toolsGroupByDeliveryFormatForTop4Domains", domainService.ListToolsGroupByDeliveryFormatForTop4Domain)
}
