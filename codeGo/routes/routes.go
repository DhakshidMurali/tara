package routes

import "github.com/gin-gonic/gin"

func APIRoutes(server *gin.Engine) {

	authRoutes := server.Group("/")

	// Create
	// Relation
	authRoutes.POST("/create/relation/toolRequestByEmployee")
	authRoutes.POST("/create/relation/toolAccessToEmployee")
	authRoutes.POST("/create/relation/toolManagedByEmployee")
	authRoutes.POST("/create/relation/toolManagedByDepartment")

	authRoutes.POST("/create/relation/employeeCollaboratedWithEmployee")
	authRoutes.POST("/create/relation/employeeWorksOnProduct")
	authRoutes.POST("/create/relation/employeeSkilledSkills")
	authRoutes.POST("/create/relation/employeeReportToEmployee")

	authRoutes.POST("/create/relation/departmentManagedByEmployee")
	authRoutes.POST("/create/relation/communityMemberEmployee")
	authRoutes.POST("/create/relation/communityAccessRequiredEmployee")

	authRoutes.POST("/create/relation/communicationPostedInCommunity")
	authRoutes.POST("/create/relation/communiationPostedByEmployee")

}
