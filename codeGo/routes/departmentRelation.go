package routes

import (
	"net/http"

	"github.com/DhakshidMurali/tara/db"
	"github.com/DhakshidMurali/tara/model"
	"github.com/gin-gonic/gin"
)

func createDepartmentManagedByEmployee(context *gin.Context){
	var departmentManagedByEmployee model.DepartmentManagedByEmployee
	err:=context.ShouldBindJSON(&departmentManagedByEmployee)

	if err!=nil{
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := departmentManagedByEmployee.MakeQuery("CREATE")
	params := departmentManagedByEmployee.MakeParams("CREATE")
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