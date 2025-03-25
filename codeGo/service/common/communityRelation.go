package commonService

import (
	"net/http"

	"github.com/DhakshidMurali/tara/model/common"
	repository "github.com/DhakshidMurali/tara/repository/common"
	"github.com/gin-gonic/gin"
)

var communityList = []model.Community{}
var userList = []model.User{}

func CreateCommunityMemberUser(context *gin.Context) {
	var communityMemberUser model.CommunityMemberUser
	err := context.ShouldBindJSON(&communityMemberUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	repository.CreateCommunityMemberUser(&communityMemberUser, context)
}

func CreateCommunityCreatedByUser(context *gin.Context) {
	var communityCreatedByUser model.CommunityCreatedByUser
	err := context.ShouldBindJSON(&communityCreatedByUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	repository.CreateCommunityCreatedByUser(&communityCreatedByUser, context)
}

func ListCommunityCreatedByUser(context *gin.Context) {
	var communityCreatedByUser model.CommunityCreatedByUser
	err := context.ShouldBindJSON(&communityCreatedByUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	repository.ListCommunityCreatedByUser(&communityCreatedByUser, context)

}

func ListCommunityMemberUser(context *gin.Context) {
	var communityMemberUser model.CommunityMemberUser
	err := context.ShouldBindJSON(&communityMemberUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	repository.ListCommunityMemberUser(&communityMemberUser, context)
}
