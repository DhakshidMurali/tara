package routes

import (
	domainRoutes "github.com/DhakshidMurali/tara/routes/domain"
	"github.com/gin-gonic/gin"
)

func APIRoutes(server *gin.Engine) {

	authRoutes := server.Group("/")

	/*
	* Create Relation
	 */
	authRoutes.POST("/create/relation/toolRequestByUser", createToolRequestByUser)
	authRoutes.POST("/create/relation/toolAccessToUser", createToolAccessToUser)
	authRoutes.POST("/create/relation/toolManagedByUser", createToolManagedByUser)
	authRoutes.POST("/create/relation/toolComesUnderDomain", createToolComesUnderDomain)

	authRoutes.POST("/create/relation/userCollaboratedWithUser", createUserCollaboratedWithUser)
	authRoutes.POST("/create/relation/userWorksInTools", createUserWorksInTools)
	authRoutes.POST("/create/relation/userSkilledInSkills", createUserSkilledInSkills)

	authRoutes.POST("/create/relation/communityMemberUser", createCommunityMemberUser)
	authRoutes.POST("/create/relation/communityCreatedByUser", createCommunityCreatedByUser)

	authRoutes.POST("/create/relation/communicationPostedInCommunity", createCommunicationPostedInCommunity)
	authRoutes.POST("/create/relation/communiationPostedByUser", createCommunicationPostedByUser)

	/*
	* Update Node
	 */
	authRoutes.POST("/update/node/user", updateUser)
	authRoutes.POST("/update/node/tool", updateTool)
	authRoutes.POST("/update/node/domain", updateDomain)
	authRoutes.POST("/update/node/community", updateCommunity)
	authRoutes.POST("/update/node/communication", updateCommunication)
	authRoutes.POST("/update/node/skills", updateSkills)

	/*
	* Retrieve data for Communication
	 */
	authRoutes.POST("/get/list/communicationPostedInCommunity", listCommunicationPostedInCommunity)
	authRoutes.POST("/get/list/communicationPostedByUser", listCommunicationPostedByUser)

	/*
	* Retrieve data for Community
	 */
	authRoutes.POST("/get/list/communityCreatedByEployee", listCommunityCreatedByUser)
	authRoutes.POST("/get/list/communityMemberUser", listCommunityMemberUser)
	authRoutes.POST("/get/list/community", listCommunity)

	/*
	* Retrieve data for Domain
	 */
	authRoutes.POST("/get/list/domain", listDomain)

	/*
	* Retrieve data for User
	 */
	authRoutes.POST("/get/list/userCollaboratedWithUser", listUserCollaboratedWithUser)
	authRoutes.POST("/get/list/userWorksInTools/:node", listUserWorksInTools)
	authRoutes.POST("/get/list/userSkilledInSkills/:node", listUserSkilledInSkills)
	/*
	* Retrieve data for Tool
	 */
	authRoutes.POST("/get/list/tool", listTool)
	authRoutes.POST("/get/list/toolRequestByUser/:node", listToolRequestByUser)
	authRoutes.POST("/get/list/toolAccessToUser/:node", listToolAccessToUser)
	authRoutes.POST("/get/list/toolManagedByUser/:node", listToolManagedByUser)
	authRoutes.POST("/get/list/toolComesUnderDomain/:node", listToolComesUnderDomain)

	/*
	* Tool Page Routes
	 */
	domainRoutes.APIRoutes(server)

}
