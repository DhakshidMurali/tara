package toolHandler

import (
	toolService "github.com/DhakshidMurali/tara/service/tool"
	"github.com/gin-gonic/gin"
)

func APIRoutes(server *gin.Engine) {
	toolsRoutes := server.Group("/")

	toolsRoutes.GET("get/list/topToolsByCommunication", toolService.ListToolGroupByDomain)

}
