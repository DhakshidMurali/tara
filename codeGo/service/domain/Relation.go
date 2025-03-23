package domainService

import (
	domainService "github.com/DhakshidMurali/tara/repository/domain"
	"github.com/gin-gonic/gin"
)

func ListToolGroupByDomain(context *gin.Context) {
	domainService.ListToolGroupByDomain(context)
}

func ListToolsGroupByDeliveryFormatForTop4Domain(context *gin.Context) {
	domainService.ListToolGroupByDomain(context)
}
