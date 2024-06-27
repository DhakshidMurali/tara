package routes

import (
	"net/http"

	"github.com/DhakshidMurali/tara/model"
	"github.com/gin-gonic/gin"
)

func createCommunityMemberEmployee(context *gin.Context){
	var communityMemberEmployee model.CommunityMemberEmployee
	err:=context.ShouldBindJSON(&communityMemberEmployee)

	if err!=nil{
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
}

func createCommunityCreatedByEmployee(context * gin.Context){
	var communityCreatedByEmployee model.CommunityCreatedByEmployee
	err:=context.ShouldBindJSON(&communityCreatedByEmployee)

	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"Description": "Received data can't be parsed"})
	}
}