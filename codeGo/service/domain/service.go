package domainService

import (
	domainRepository "github.com/DhakshidMurali/tara/repository/domain"
	"github.com/gin-gonic/gin"
)

func GetDashBoardDetails(context *gin.Context) {
	domainRepository.GetDashBoardDetails(context)
}

func ListToolsGroupByDeliveryFormatForTop5Domain(context *gin.Context) {
	domainRepository.ListToolsGroupByDeliveryFormatForTop4Domain(context)
}
