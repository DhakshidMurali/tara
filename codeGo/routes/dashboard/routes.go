package dashboardRoutes

import "github.com/gin-gonic/gin"

func DashboardRoutes(server *gin.Engine) {
	authRoutes := server.Group("/")

	authRoutes.GET("/get/list/toolGroupByDomain", listToolGroupByDomain)
}
