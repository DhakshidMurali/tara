package toolService

import (
	toolRepository "github.com/DhakshidMurali/tara/repository/tool"
	"github.com/gin-gonic/gin"
)

func ListToolGroupByDomain(context *gin.Context) {
	toolRepository.ListToolGroupByCommunication(context)
}
