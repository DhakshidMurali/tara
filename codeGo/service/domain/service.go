package domainService

import (
	domainRepository "github.com/DhakshidMurali/tara/repository/domain"
	"github.com/gin-gonic/gin"
)

func GetDomainPageDetails(context *gin.Context) {
	domainRepository.ListToolGroupByDomain(context)
	domainRepository.ListToolsGroupByDeliveryFormatForTop4Domain(context)

}
func ListToolGroupByDomain(context *gin.Context) {
	domainRepository.ListToolGroupByDomain(context)
}

func ListToolsGroupByDeliveryFormatForTop5Domain(context *gin.Context) {
	domainRepository.ListToolsGroupByDeliveryFormatForTop4Domain(context)
}
