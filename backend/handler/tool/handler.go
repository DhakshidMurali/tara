package toolHandler

import (
	toolService "github.com/DhakshidMurali/tara/service/tool"
	"github.com/gin-gonic/gin"
)

func APIRoutes(server *gin.Engine) {
	toolsRoutes := server.Group("/tool")

	toolsRoutes.GET("get/list/count-tools-by-communication", toolService.GetTopToolsByCommunication)

	toolsRoutes.GET("/get/list/tools-by-domain-details", toolService.GetToolByDomainDetails)

	toolsRoutes.GET("/get/list/group-likes-stats-by-domains", toolService.GetLikesStatsByLikes)


}
