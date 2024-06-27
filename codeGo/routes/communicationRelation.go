package routes

import (
	"net/http"

	"github.com/DhakshidMurali/tara/db"
	"github.com/DhakshidMurali/tara/model"
	"github.com/gin-gonic/gin"
)

func createCommunicationPostedInCommunity(context *gin.Context) {
	var communicationPostedInCommunity model.CommunicationPostedInCommunity
	err := context.ShouldBindJSON(&communicationPostedInCommunity)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := communicationPostedInCommunity.MakeQuery()
	params := communicationPostedInCommunity.MakeParams()
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

func createCommunicationPostedByEmployee(context *gin.Context) {
	var communicationPostedByEmployee model.CommunicationPostedByEmployee
	err := context.ShouldBindJSON(&communicationPostedByEmployee)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := communicationPostedByEmployee.MakeQuery()
	params := communicationPostedByEmployee.MakeParams()
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
