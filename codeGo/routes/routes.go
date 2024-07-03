package routes

import "github.com/gin-gonic/gin"

func APIRoutes(server *gin.Engine) {

	authRoutes := server.Group("/")

	// Create Relation
	authRoutes.POST("/create/relation/toolRequestByEmployee", createToolRequestByEmployee)
	authRoutes.POST("/create/relation/toolAccessToEmployee", createToolAccessToEmployee)
	authRoutes.POST("/create/relation/toolManagedByEmployee", createToolManagedByEmployee)
	authRoutes.POST("/create/relation/toolComesUnderDepartment", createToolComesUnderDepartment)

	authRoutes.POST("/create/relation/employeeCollaboratedWithEmployee", createEmployeeCollaboratedWithEmployee)
	authRoutes.POST("/create/relation/employeeWorksInTools", createEmployeeWorksInTools)
	authRoutes.POST("/create/relation/employeeSkilledInSkills", createEmployeeSkilledInSkills)
	authRoutes.POST("/create/relation/employeeReportToEmployee", createEmployeeReportToEmployee)

	authRoutes.POST("/create/relation/departmentManagedByEmployee", createDepartmentManagedByEmployee)
	authRoutes.POST("/create/relation/communityMemberEmployee", createCommunityMemberEmployee)
	authRoutes.POST("/create/relation/communityCreatedByEmployee", createCommunityCreatedByEmployee)

	authRoutes.POST("/create/relation/communicationPostedInCommunity", createCommunicationPostedInCommunity)
	authRoutes.POST("/create/relation/communiationPostedByEmployee", createCommunicationPostedByEmployee)

	// Update Node
	authRoutes.POST("/update/node/employee", updateEmployee)
	authRoutes.POST("/update/node/tool", updateTool)
	authRoutes.POST("/update/node/department", updateDepartment)
	authRoutes.POST("/update/node/community", updateCommunity)
	authRoutes.POST("/update/node/communication", updateCommunication)
	authRoutes.POST("/update/node/skills", updateSkills)

	// Retrieve
	authRoutes.POST("/get/list/community/communication", listCommunicationPostedInCommunity)
	authRoutes.POST("/get/list/employee/communication", listCommunicationPostedByEmployee)
	authRoutes.POST("/get/list/community", updateEmployee)
	authRoutes.POST("/get/detail/community", updateEmployee)
	authRoutes.POST("/get/list/department", updateEmployee)
	authRoutes.POST("/get/detail/department", updateEmployee)

}
