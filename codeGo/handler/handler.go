package handler

import (
	domainHandler "github.com/DhakshidMurali/tara/handler/domain"
	commonService "github.com/DhakshidMurali/tara/service/common"
	"github.com/gin-gonic/gin"
)

func APIRoutes(server *gin.Engine) {

	authHandler := server.Group("/")

	/*
	* Create Relation
	 */
	authHandler.POST("/create/relation/toolRequestByUser", commonService.CreateToolRequestByUser)
	authHandler.POST("/create/relation/toolAccessToUser", commonService.CreateToolAccessToUser)
	authHandler.POST("/create/relation/toolManagedByUser", commonService.CreateToolManagedByUser)
	authHandler.POST("/create/relation/toolComesUnderDomain", commonService.CreateToolComesUnderDomain)

	authHandler.POST("/create/relation/userCollaboratedWithUser", commonService.CreateUserCollaboratedWithUser)
	authHandler.POST("/create/relation/userWorksInTools", commonService.CreateUserWorksInTools)
	authHandler.POST("/create/relation/userSkilledInSkills", commonService.CreateUserSkilledInSkills)

	authHandler.POST("/create/relation/communityMemberUser", commonService.CreateCommunityMemberUser)
	authHandler.POST("/create/relation/communityCreatedByUser", commonService.CreateCommunityCreatedByUser)

	authHandler.POST("/create/relation/communicationPostedInCommunity", commonService.CreateCommunicationPostedInCommunity)
	authHandler.POST("/create/relation/communiationPostedByUser", commonService.CreateCommunicationPostedByUser)

	/*
	* Update Node
	 */
	authHandler.POST("/update/node/user", commonService.UpdateUser)
	authHandler.POST("/update/node/tool", commonService.UpdateTool)
	authHandler.POST("/update/node/domain", commonService.UpdateDomain)
	authHandler.POST("/update/node/community", commonService.UpdateCommunity)
	authHandler.POST("/update/node/communication", commonService.UpdateCommunication)
	authHandler.POST("/update/node/skills", commonService.UpdateSkills)

	/*
	* Retrieve data for Communication
	 */
	authHandler.POST("/get/list/communicationPostedInCommunity", commonService.ListCommunicationPostedInCommunity)
	authHandler.POST("/get/list/communicationPostedByUser", commonService.ListCommunicationPostedByUser)

	/*
	* Retrieve data for Community
	 */
	authHandler.POST("/get/list/communityCreatedByEployee", commonService.ListCommunityCreatedByUser)
	authHandler.POST("/get/list/communityMemberUser", commonService.ListCommunityMemberUser)
	authHandler.POST("/get/list/community", commonService.ListCommunity)

	/*
	* Retrieve data for Domain
	 */
	authHandler.POST("/get/list/domain", commonService.ListDomain)

	/*
	* Retrieve data for User
	 */
	authHandler.POST("/get/list/userCollaboratedWithUser", commonService.ListUserCollaboratedWithUser)
	authHandler.POST("/get/list/userWorksInTools/:node", commonService.ListUserWorksInTools)
	authHandler.POST("/get/list/userSkilledInSkills/:node", commonService.ListUserSkilledInSkills)
	/*
	* Retrieve data for Tool
	 */
	authHandler.POST("/get/list/tool", commonService.ListTool)
	authHandler.POST("/get/list/toolRequestByUser/:node", commonService.ListToolRequestByUser)
	authHandler.POST("/get/list/toolAccessToUser/:node", commonService.ListToolAccessToUser)
	authHandler.POST("/get/list/toolManagedByUser/:node", commonService.ListToolManagedByUser)
	authHandler.POST("/get/list/toolComesUnderDomain/:node", commonService.ListToolComesUnderDomain)

	/*
	* Tool Page Handler
	 */
	domainHandler.APIRoutes(server)

}
