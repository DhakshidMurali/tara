package routes

import "github.com/gin-gonic/gin"

func APIRoutes(server *gin.Engine) {

	authRoutes := server.Group("/")

	/*
	* Create Relation
	 */
	authRoutes.POST("/create/relation/toolRequestByEmployee", createToolRequestByEmployee)
	authRoutes.POST("/create/relation/toolAccessToEmployee", createToolAccessToEmployee)
	authRoutes.POST("/create/relation/toolManagedByEmployee", createToolManagedByEmployee)
	authRoutes.POST("/create/relation/toolComesUnderDepartment", createToolComesUnderDepartment)

	authRoutes.POST("/create/relation/employeeCollaboratedWithEmployee", createEmployeeCollaboratedWithEmployee)
	authRoutes.POST("/create/relation/employeeWorksInTools", createEmployeeWorksInTools)
	authRoutes.POST("/create/relation/employeeSkilledInSkills", createEmployeeSkilledInSkills)
	authRoutes.POST("/create/relation/employeeReportToEmployee", createEmployeeReportToEmployee)
	authRoutes.POST("/create/relation/employeeComesUnderDepartment", createEmployeeComesUnderDepartment)

	authRoutes.POST("/create/relation/departmentManagedByEmployee", createDepartmentManagedByEmployee)
	authRoutes.POST("/create/relation/communityMemberEmployee", createCommunityMemberEmployee)
	authRoutes.POST("/create/relation/communityCreatedByEmployee", createCommunityCreatedByEmployee)

	authRoutes.POST("/create/relation/communicationPostedInCommunity", createCommunicationPostedInCommunity)
	authRoutes.POST("/create/relation/communiationPostedByEmployee", createCommunicationPostedByEmployee)

	/*
	* Update Node
	 */
	authRoutes.POST("/update/node/employee", updateEmployee)
	authRoutes.POST("/update/node/tool", updateTool)
	authRoutes.POST("/update/node/department", updateDepartment)
	authRoutes.POST("/update/node/community", updateCommunity)
	authRoutes.POST("/update/node/communication", updateCommunication)
	authRoutes.POST("/update/node/skills", updateSkills)

	/*
	* Retrieve data for Communication
	 */
	authRoutes.POST("/get/list/communicationPostedInCommunity", listCommunicationPostedInCommunity)
	authRoutes.POST("/get/list/communicationPostedByEmployee", listCommunicationPostedByEmployee)

	/*
	* Retrieve data for Community
	 */
	authRoutes.POST("/get/list/communityCreatedByEployee", listCommunityCreatedByEmployee)
	authRoutes.POST("/get/list/communityMemberEmployee", listCommunityMemberEmployee)
	authRoutes.POST("/get/list/community", listCommunity)

	/*
	* Retrieve data for Department
	 */
	authRoutes.POST("/get/list/department", listDepartment)
	authRoutes.POST("/get/list/departmentManagedByEmployee", listDepartmentManagedByEmployee)

	/*
	* Retrieve data for Employee
	 */
	authRoutes.POST("/get/list/employeeCollaboratedWithEmployee", listEmployeeCollaboratedWithEmployee)
	authRoutes.POST("/get/list/employeeWorksInTools/:node", listEmployeeWorksInTools)
	authRoutes.POST("/get/list/employeeSkilledInSkills/:node", listEmployeeSkilledInSkills)
	authRoutes.POST("/get/list/employeeReportToEmployee/:node", listEmployeeReportToEmployee)
	authRoutes.POST("/get/list/employeeComesUnderDepartment/:node", listEmployeeComesUnderDepartment)

	/*
	* Retrieve data for Tool
	 */
	authRoutes.POST("/get/list/tool", listTool)
	authRoutes.POST("/get/list/toolRequestByEmployee/:node", listToolRequestByEmployee)
	authRoutes.POST("/get/list/toolAccessToEmployee/:node", listToolAccessToEmployee)
	authRoutes.POST("/get/list/toolManagedByEmployee/:node", listToolManagedByEmployee)
	authRoutes.POST("/get/list/toolComesUnderDepartment/:node", listToolComesUnderDepartment)

	/*
	* Dashboard Page Api
	 */
	authRoutes.POST("/get/list/employeeGroupByDepartment", listEmployeeGroupByDepartment)

}
