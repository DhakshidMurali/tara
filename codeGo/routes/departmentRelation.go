package routes

import (
	"net/http"

	"github.com/DhakshidMurali/tara/model"
	"github.com/gin-gonic/gin"
)

func createDepartmentManagedByEmployee(context *gin.Context){
	var departmentManagedByEmployee model.DepartmentManagedByEmployee
	err:=context.ShouldBindJSON(&departmentManagedByEmployee)

	if err!=nil{
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
}