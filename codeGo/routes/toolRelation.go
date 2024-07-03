package routes

import (
	"net/http"

	"github.com/DhakshidMurali/tara/db"
	"github.com/DhakshidMurali/tara/model"
	"github.com/gin-gonic/gin"
)

func createToolRequestByEmployee(context *gin.Context) {
	var toolRequestByEmployee model.ToolRequestByEmployee
	err := context.ShouldBindJSON(&toolRequestByEmployee)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := toolRequestByEmployee.MakeQuery("CREATE")
	params := toolRequestByEmployee.MakeParams("CREATE")
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
func createToolAccessToEmployee(context *gin.Context) {
	var toolAccessToEmployee model.ToolAccessToEmployee
	err := context.ShouldBindJSON(&toolAccessToEmployee)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := toolAccessToEmployee.MakeQuery("CREATE")
	params := toolAccessToEmployee.MakeParams("CREATE")
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
func createToolManagedByEmployee(context *gin.Context) {
	var toolManagedByEmployee model.ToolManagedByEmployee
	err := context.ShouldBindJSON(&toolManagedByEmployee)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := toolManagedByEmployee.MakeQuery("CREATE")
	params := toolManagedByEmployee.MakeParams("CREATE")
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
func createToolComesUnderDepartment(context *gin.Context) {
	var toolComesUnderDepartment model.ToolComesUnderDepartment
	err := context.ShouldBindJSON(&toolComesUnderDepartment)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := toolComesUnderDepartment.MakeQuery("CREATE")
	params := toolComesUnderDepartment.MakeParams("CREATE")
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
