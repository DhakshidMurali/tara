package routes

import (
	"net/http"

	"github.com/DhakshidMurali/tara/model"
	"github.com/gin-gonic/gin"
)

func createToolRequestByEmployee(context *gin.Context) {
	var toolRequestByEmployee model.ToolRequestByEmployee
	err := context.ShouldBindJSON(&toolRequestByEmployee)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}

}
func createToolAccessToEmployee(context *gin.Context) {
	var toolAccessToEmployee model.ToolAccessToEmployee
	err := context.ShouldBindJSON(&toolAccessToEmployee)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}

}
func createToolManagedByEmployee(context *gin.Context) {
	var toolManagedByEmployee model.ToolManagedByEmployee
	err := context.ShouldBindJSON(&toolManagedByEmployee)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}

}
func createToolComesUnderDepartment(context *gin.Context) {
	var toolComesUnderDepartment model.ToolComesUnderDepartment
	err := context.ShouldBindJSON(&toolComesUnderDepartment)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}

}
