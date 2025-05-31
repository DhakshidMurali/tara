package commonService

import (
	"net/http"

	"github.com/DhakshidMurali/tara/model/common"
	repository "github.com/DhakshidMurali/tara/repository/common"
	"github.com/gin-gonic/gin"
)

func CreateUserCollaboratedWithUser(context *gin.Context) {
	var userCollaboratedWithUser model.UserCollaboratedWithUser
	err := context.ShouldBindJSON((&userCollaboratedWithUser))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	repository.CreateUserCollaboratedWithUser(&userCollaboratedWithUser, context)
}

func CreateUserWorksInTools(context *gin.Context) {
	var userWorksInTools model.UserWorksInTools
	err := context.ShouldBindJSON((&userWorksInTools))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	repository.CreateUserWorksInTools(&userWorksInTools, context)
}

func CreateUserSkilledInSkills(context *gin.Context) {
	var userSkilledInSkills model.UserSkilledInSkills
	err := context.ShouldBindJSON((&userSkilledInSkills))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	repository.CreateUserSkilledInSkills(&userSkilledInSkills, context)
}

func ListUserCollaboratedWithUser(context *gin.Context) {
	var userCollaboratedWithUser model.UserCollaboratedWithUser
	err := context.ShouldBindJSON((&userCollaboratedWithUser))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	repository.ListUserCollaboratedWithUser(&userCollaboratedWithUser, context)

}
func ListUserWorksInTools(context *gin.Context) {
	var userWorksInTools model.UserWorksInTools
	err := context.ShouldBindJSON((&userWorksInTools))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	retrieveNode := context.Param("node")

	if retrieveNode == "USER" {
		repository.ListUserWorksInToolsByUser(&userWorksInTools, context)
	}
	if retrieveNode == "TOOL" {
		repository.ListUserWorksInToolsByTool(&userWorksInTools, context)
	}
}

func ListUserSkilledInSkills(context *gin.Context) {
	var userSkilledInSkills model.UserSkilledInSkills
	err := context.ShouldBindJSON((&userSkilledInSkills))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	retrieveNode := context.Param("node")
	if retrieveNode == "USER" {
		repository.ListUserSkilledInSkillsByUser(&userSkilledInSkills, context)
	}
	if retrieveNode == "SKILLS" {
		repository.ListUserSkilledInSkillsBySkills(&userSkilledInSkills, context)

	}
}
