package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DhakshidMurali/tara/db"
	"github.com/DhakshidMurali/tara/model"
	"github.com/DhakshidMurali/tara/util"
	"github.com/gin-gonic/gin"
)


func createDepartmentManagedByEmployee(context *gin.Context) {
	var departmentManagedByEmployee model.DepartmentManagedByEmployee
	err := context.ShouldBindJSON(&departmentManagedByEmployee)

	if err != nil {
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

func listAllDepartment(context *gin.Context) {
	var department model.Department
	var data model.Department

	query := department.MakeQuery("LIST_DEPARTMENT")
	params := map[string]any{}
	result := db.Execute(query, params)

	for _, record := range result.Records {
		mapRecord, _ := record.Get("n1")
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &data)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}
		departmentList = append(departmentList, data)
	}
	context.JSON(http.StatusOK, departmentList)
	departmentList = nil
}

func listDepartmentManagedByEmployee(context *gin.Context) {
	var departmentManagedByEmployee model.DepartmentManagedByEmployee
	var data model.Department

	query := departmentManagedByEmployee.MakeQuery("LIST_DEPARTMENT")
	params := departmentManagedByEmployee.MakeParams("LIST_DEPARTMENT")
	result := db.Execute(query, params)

	for _, record := range result.Records {
		mapRecord, _ := record.Get("n1")
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &data)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}
		departmentList = append(departmentList, data)
	}
	context.JSON(http.StatusOK, departmentList)
	departmentList = nil
}
