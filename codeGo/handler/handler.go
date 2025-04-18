package handler

import (
	domainHandler "github.com/DhakshidMurali/tara/handler/domain"
	toolHandler "github.com/DhakshidMurali/tara/handler/tool"
	service "github.com/DhakshidMurali/tara/service/common"
	"github.com/gin-gonic/gin"
)

func APIRoutes(server *gin.Engine) {

	authHandler := server.Group("/")

	/*
	* Create Relation
	 */
	authHandler.POST("/create/relation/toolRequestByUser", service.CreateToolRequestByUser)
	authHandler.POST("/create/relation/toolAccessToUser", service.CreateToolAccessToUser)
	authHandler.POST("/create/relation/toolManagedByUser", service.CreateToolManagedByUser)
	authHandler.POST("/create/relation/toolComesUnderDomain", service.CreateToolComesUnderDomain)

	authHandler.POST("/create/relation/userCollaboratedWithUser", service.CreateUserCollaboratedWithUser)
	authHandler.POST("/create/relation/userWorksInTools", service.CreateUserWorksInTools)
	authHandler.POST("/create/relation/userSkilledInSkills", service.CreateUserSkilledInSkills)

	authHandler.POST("/create/relation/communityMemberUser", service.CreateCommunityMemberUser)
	authHandler.POST("/create/relation/communityCreatedByUser", service.CreateCommunityCreatedByUser)

	authHandler.POST("/create/relation/communicationPostedInCommunity", service.CreateCommunicationPostedInCommunity)
	authHandler.POST("/create/relation/communiationPostedByUser", service.CreateCommunicationPostedByUser)

	/*
	* Update Node
	 */
	authHandler.POST("/update/node/user", service.UpdateUser)
	authHandler.POST("/update/node/tool", service.UpdateTool)
	authHandler.POST("/update/node/domain", service.UpdateDomain)
	authHandler.POST("/update/node/community", service.UpdateCommunity)
	authHandler.POST("/update/node/communication", service.UpdateCommunication)
	authHandler.POST("/update/node/skills", service.UpdateSkills)

	/*
	* Retrieve data for Communication
	 */
	authHandler.POST("/get/list/communicationPostedInCommunity", service.ListCommunicationPostedInCommunity)
	authHandler.POST("/get/list/communicationPostedByUser", service.ListCommunicationPostedByUser)

	/*
	* Retrieve data for Community
	 */
	authHandler.POST("/get/list/communityCreatedByEployee", service.ListCommunityCreatedByUser)
	authHandler.POST("/get/list/communityMemberUser", service.ListCommunityMemberUser)
	authHandler.POST("/get/list/community", service.ListCommunity)

	/*
	* Retrieve data for Domain
	 */
	authHandler.POST("/get/list/domain", service.ListDomain)

	/*
	* Retrieve data for User
	 */
	authHandler.POST("/get/list/userCollaboratedWithUser", service.ListUserCollaboratedWithUser)
	authHandler.POST("/get/list/userWorksInTools/:node", service.ListUserWorksInTools)
	authHandler.POST("/get/list/userSkilledInSkills/:node", service.ListUserSkilledInSkills)
	/*
	* Retrieve data for Tool
	 */
	authHandler.POST("/get/list/tool", service.ListTool)
	authHandler.POST("/get/list/toolRequestByUser/:node", service.ListToolRequestByUser)
	authHandler.POST("/get/list/toolAccessToUser/:node", service.ListToolAccessToUser)
	authHandler.POST("/get/list/toolManagedByUser/:node", service.ListToolManagedByUser)
	authHandler.POST("/get/list/toolComesUnderDomain/:node", service.ListToolComesUnderDomain)

	/*
	* Tool Page Handler
	 */
	domainHandler.APIRoutes(server)
	toolHandler.APIRoutes(server)
}
