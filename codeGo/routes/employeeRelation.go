package routes

import (
	"net/http"

	"github.com/DhakshidMurali/tara/model"
	"github.com/gin-gonic/gin"
)

func createEmployeeCollaboratedWithEmployee(context *gin.Context){
	var employeeCollaboratedWithEmployee model.EmployeeCollaboratedWithEmployee
	err:=context.ShouldBindJSON((&employeeCollaboratedWithEmployee))

	if err!=nil{
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
}

func createEmployeeWorksInTools (context *gin.Context){
	var employeeWorksInTools model.EmployeeWorksInTools
	err:=context.ShouldBindJSON((&employeeWorksInTools))

	if err!=nil{
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
}

func createEmployeeSkilledInSkills (context *gin.Context){
	var employeeSkilledInSkills model.EmployeeSkilledInSkills
	err:=context.ShouldBindJSON((&employeeSkilledInSkills))

	if err!=nil{
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
}

func createEmployeeReportToEmployee (context *gin.Context){
	var employeeReportToEmployee model.EmployeeReportToEmployee
	err:=context.ShouldBindJSON((&employeeReportToEmployee))

	if err!=nil{
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
}