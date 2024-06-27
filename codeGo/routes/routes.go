package routes

import "github.com/gin-gonic/gin"

func APIRoutes(server *gin.Engine) {

	authRoutes := server.Group("/")

	// Create Relation
	authRoutes.POST("/create/relation/toolRequestByEmployee")
	authRoutes.POST("/create/relation/toolAccessToEmployee")
	authRoutes.POST("/create/relation/toolManagedByEmployee")
	authRoutes.POST("/create/relation/toolComesUnderDepartment")

	authRoutes.POST("/create/relation/employeeCollaboratedWithEmployee")
	authRoutes.POST("/create/relation/employeeWorksInTools")
	authRoutes.POST("/create/relation/employeeSkilledInSkills")
	authRoutes.POST("/create/relation/employeeReportToEmployee")

	authRoutes.POST("/create/relation/departmentManagedByEmployee")
	authRoutes.POST("/create/relation/communityMemberEmployee")
	authRoutes.POST("/create/relation/communityCreatedByEmployee")

	authRoutes.POST("/create/relation/communicationPostedInCommunity", createCommunicationPostedInCommunity)
	authRoutes.POST("/create/relation/communiationPostedByEmployee")

}
