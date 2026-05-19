package domainRoutes

import (
	domainService "github.com/DhakshidMurali/tara/service/domain"
	"github.com/gin-gonic/gin"
)

func APIRoutes(server *gin.Engine) {

	domainRoutes := server.Group("/domain")

	domainRoutes.GET("/get/list/domain-details", domainService.GetDomainDetails)

	domainRoutes.GET("/get/list/count-tools-by-domains", domainService.GetTopDomainsByTool)

	domainRoutes.GET("/get/list/group-delivery-format-by-domains", domainService.GetDeliveryFormatByDomains)


}
