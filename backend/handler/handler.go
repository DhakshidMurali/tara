package handler

import (
	domainHandler "github.com/DhakshidMurali/tara/handler/domain"
	toolHandler "github.com/DhakshidMurali/tara/handler/tool"

	"github.com/gin-gonic/gin"
)

func APIRoutes(server *gin.Engine) {
	domainHandler.APIRoutes(server)
	toolHandler.APIRoutes(server)
}
