package domainService

import (
	"fmt"
	"net/http"

	domainRepository "github.com/DhakshidMurali/tara/repository/domain"
	"github.com/gin-gonic/gin"
)

func GetDomainDetails(context *gin.Context) {
	response, err := domainRepository.GetDomainDetails()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to fetch toresponseols: %v", err),
		})
		return
	}
	context.JSON(http.StatusOK, response)
}

func GetTopDomainsByTool(context *gin.Context) {
	response, err := domainRepository.GetTopDomainsByTool()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to fetch toresponseols: %v", err),
		})
		return
	}
	context.JSON(http.StatusOK, response)
}

func GetDeliveryFormatByDomains(context *gin.Context) {
	response, err := domainRepository.GetDeliveryFormatByDomains()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to fetch toresponseols: %v", err),
		})
		return
	}
	context.JSON(http.StatusOK, response)
}
