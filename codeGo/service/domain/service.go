package domainService

import (
	domainRepository "github.com/DhakshidMurali/tara/repository/domain"
	"github.com/gin-gonic/gin"
)

func GetDomainDetails(context *gin.Context) {
	domainRepository.GetDomainDetails(context)
}

func GetTopDomainsByTool(context *gin.Context) {
	domainRepository.GetTopDomainsByTool(context)
}

func GetDeliveryFormatByDomains(context *gin.Context) {
	domainRepository.GetDeliveryFormatByDomains(context)
}

