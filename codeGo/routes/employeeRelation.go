package routes

import (
	"net/http"

	"github.com/DhakshidMurali/tara/db"
	"github.com/DhakshidMurali/tara/model"
	"github.com/gin-gonic/gin"
)

func createEmployeeCollaboratedWithEmployee(context *gin.Context) {
	var employeeCollaboratedWithEmployee model.EmployeeCollaboratedWithEmployee
	err := context.ShouldBindJSON((&employeeCollaboratedWithEmployee))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := employeeCollaboratedWithEmployee.MakeQuery("CREATE")
	params := employeeCollaboratedWithEmployee.MakeParams("CREATE")
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

func createEmployeeWorksInTools(context *gin.Context) {
	var employeeWorksInTools model.EmployeeWorksInTools
	err := context.ShouldBindJSON((&employeeWorksInTools))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := employeeWorksInTools.MakeQuery("CREATE")
	params := employeeWorksInTools.MakeParams("CREATE")
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

func createEmployeeSkilledInSkills(context *gin.Context) {
	var employeeSkilledInSkills model.EmployeeSkilledInSkills
	err := context.ShouldBindJSON((&employeeSkilledInSkills))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := employeeSkilledInSkills.MakeQuery("CREATE")
	params := employeeSkilledInSkills.MakeParams("CREATE")
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

func createEmployeeReportToEmployee(context *gin.Context) {
	var employeeReportToEmployee model.EmployeeReportToEmployee
	err := context.ShouldBindJSON((&employeeReportToEmployee))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := employeeReportToEmployee.MakeQuery("CREATE")
	params := employeeReportToEmployee.MakeParams("CREATE")
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
