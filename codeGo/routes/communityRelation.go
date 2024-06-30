package routes

import (
	"net/http"

	"github.com/DhakshidMurali/tara/db"
	"github.com/DhakshidMurali/tara/model"
	"github.com/gin-gonic/gin"
)

func createCommunityMemberEmployee(context *gin.Context){
	var communityMemberEmployee model.CommunityMemberEmployee
	err:=context.ShouldBindJSON(&communityMemberEmployee)

	if err!=nil{
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := communityMemberEmployee.MakeQuery()
	params := communityMemberEmployee.MakeParams()
	result := db.Execute(query, params)

	isCreated := result.Records[0].Values[0].(bool)
	if isCreated {
		context.JSON(http.StatusOK, gin.H{
			"message": "Data Created in Database",
		})
	}
	if !isCreated {
		context.JSON(http.StatusOK, gin.H{
			"message": "Data already Exist in Database",
		})
	}
}

func createCommunityCreatedByEmployee(context * gin.Context){
	var communityCreatedByEmployee model.CommunityCreatedByEmployee
	err:=context.ShouldBindJSON(&communityCreatedByEmployee)

	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"Description": "Received data can't be parsed"})
	}
	query := communityCreatedByEmployee.MakeQuery()
	params := communityCreatedByEmployee.MakeParams()
	result := db.Execute(query, params)

	isCreated := result.Records[0].Values[0].(bool)
	if isCreated {
		context.JSON(http.StatusOK, gin.H{
			"message": "Data Created in Database",
		})
	}
	if !isCreated {
		context.JSON(http.StatusOK, gin.H{
			"message": "Data already Exist in Database",
		})
	}
}