package toolService

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	model "github.com/DhakshidMurali/tara/model/tool"
	toolRepository "github.com/DhakshidMurali/tara/repository/tool"

	"github.com/gin-gonic/gin"
)

func GetTopToolsByCommunication(context *gin.Context) {
	response, err := toolRepository.GetTopToolsByCommunication()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to fetch toresponseols: %v", err),
		})
		return
	}
	context.JSON(http.StatusOK, response)
}

func GetToolByDomainDetails(context *gin.Context) {
	toolByDomainDetails, err := toolRepository.GetToolByDomainDetails()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to fetch response: %v", err),
		})
		return
	}

	var response = []model.ToolDetailsByDomainPayload{}
	var domainList = []string{}
	var subDomainList = []string{}

	for _, value := range toolByDomainDetails {
		if !slices.Contains(domainList, value.DomainName) {
			domainList = append(domainList, value.DomainName)
			newDomain := model.ToolDetailsByDomainPayload{
				DomainName: value.DomainName,
				SubDomain:  []model.SubDomain{},
			}
			response = append(response, newDomain)
		}
		joinArr := []string{value.DomainName, value.SubDomain}
		if !slices.Contains(subDomainList, strings.Join(joinArr, "-")) {
			subDomainList = append(subDomainList, strings.Join(joinArr, "-"))
			for i, domain := range response {
				if domain.DomainName == value.DomainName {
					response[i].SubDomain = append(domain.SubDomain, model.SubDomain{
						SubDomainName: value.SubDomain,
						Tool:          []string{value.ToolName},
					})
				}
			}
		}
		for i, domain := range response {
			if domain.DomainName == value.DomainName {
				for j, subDomain := range domain.SubDomain {
					if subDomain.SubDomainName == value.SubDomain {
						response[i].SubDomain[j].Tool = append(subDomain.Tool, value.ToolName)
					}
				}
			}
		}
	}

	context.JSON(http.StatusOK, response)
}

func GetLikesStatsByLikes(context *gin.Context) {
	response, err := toolRepository.GetLikesStatsByLikes()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to fetch toresponseols: %v", err),
		})
		return
	}
	context.JSON(http.StatusOK, response)
}
