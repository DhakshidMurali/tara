package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DhakshidMurali/tara/db"
	"github.com/DhakshidMurali/tara/model"
	"github.com/DhakshidMurali/tara/util"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
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
func listToolRequestByEmployee(context *gin.Context) {
	var toolRequestByEmployee model.ToolRequestByEmployee
	err := context.ShouldBindJSON(&toolRequestByEmployee)

	var toolData model.Tool
	var employeeData model.Employee

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}

	retrieveNode := context.Param("node")
	var result *neo4j.EagerResult

	if retrieveNode == "Tool" {
		query := toolRequestByEmployee.MakeQuery("LIST_TOOL_REQUESTED_BY_EMPLOYEE")
		params := toolRequestByEmployee.MakeParams("LIST_TOOL_REQUESTED_BY_EMPLOYEE")
		result = db.Execute(query, params)
		for _, record := range result.Records {
			mapRecord, _ := record.Get("n1")
			byteData := util.MarshalData(mapRecord)
			err := json.Unmarshal(byteData, &toolData)
			if err != nil {
				fmt.Println("Error Unmarshalling Json")
				panic(err)
			}
			toolList = append(toolList, toolData)
		}
		context.JSON(http.StatusOK, toolList)
		toolList = nil
	}

	if retrieveNode == "Employee" {
		query := toolRequestByEmployee.MakeQuery("LIST_EMPLOYEES_REQUESTED_TO_TOOL")
		params := toolRequestByEmployee.MakeParams("LIST_EMPLOYEES_REQUESTED_TO_TOOL")
		result = db.Execute(query, params)
		for _, record := range result.Records {
			mapRecord, _ := record.Get("n2")
			byteData := util.MarshalData(mapRecord)
			err := json.Unmarshal(byteData, &employeeData)
			if err != nil {
				fmt.Println("Error Unmarshalling Json")
				panic(err)
			}
			employeeList = append(employeeList, employeeData)
		}
		context.JSON(http.StatusOK, employeeList)
		employeeList = nil
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
func listToolAccessToEmployee(context *gin.Context) {
	var toolAccessToEmployee model.ToolAccessToEmployee
	err := context.ShouldBindJSON(&toolAccessToEmployee)

	var toolData model.Tool
	var employeeData model.Employee

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}

	retrieveNode := context.Param("node")
	var result *neo4j.EagerResult
	if retrieveNode == "TOOL" {
		query := toolAccessToEmployee.MakeQuery("LIST_TOOLS_ACCESS_BY_EMPLOYEE")
		params := toolAccessToEmployee.MakeParams("LIST_TOOLS_ACCESS_BY_EMPLOYEE")
		result = db.Execute(query, params)
		for _, record := range result.Records {
			mapRecord, _ := record.Get("n1")
			byteData := util.MarshalData(mapRecord)
			err := json.Unmarshal(byteData, &toolData)
			if err != nil {
				fmt.Println("Error Unmarshalling Json")
				panic(err)
			}
			toolList = append(toolList, toolData)
		}
		context.JSON(http.StatusOK, toolList)
		toolList = nil
	}
	if retrieveNode == "EMPLOYEE" {
		query := toolAccessToEmployee.MakeQuery("LIST_EMPLOYEES_ACCESS_TO_TOOL")
		params := toolAccessToEmployee.MakeParams("LIST_EMPLOYEES_ACCESS_TO_TOOL")
		result = db.Execute(query, params)
		for _, record := range result.Records {
			mapRecord, _ := record.Get("n2")
			byteData := util.MarshalData(mapRecord)
			err := json.Unmarshal(byteData, &employeeData)
			if err != nil {
				fmt.Println("Error Unmarshalling Json")
				panic(err)
			}
			employeeList = append(employeeList, employeeData)
		}
		context.JSON(http.StatusOK, employeeList)
		employeeList = nil
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

func listToolManagedByEmployee(context *gin.Context) {
	var toolManagedByEmployee model.ToolManagedByEmployee
	err := context.ShouldBindJSON(&toolManagedByEmployee)

	var toolData model.Tool
	var employeeData model.Employee

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}

	retrieveNode := context.Param("node")
	var result *neo4j.EagerResult
	if retrieveNode == "Tool" {
		query := toolManagedByEmployee.MakeQuery("LIST_TOOLS_MANAGED_BY_EMPLOYEE")
		params := toolManagedByEmployee.MakeParams("LIST_TOOLS_MANAGED_BY_EMPLOYEE")
		result = db.Execute(query, params)
		for _, record := range result.Records {
			mapRecord, _ := record.Get("n1")
			byteData := util.MarshalData(mapRecord)
			err := json.Unmarshal(byteData, &toolData)
			if err != nil {
				fmt.Println("Error Unmarshalling Json")
				panic(err)
			}
			toolList = append(toolList, toolData)
		}
		context.JSON(http.StatusOK, toolList)
		employeeList = nil
	}
	if retrieveNode == "Employee" {
		query := toolManagedByEmployee.MakeQuery("LIST_EMPLOYEES_MANAGES_TOOL")
		params := toolManagedByEmployee.MakeParams("LIST_EMPLOYEES_MANAGES_TOOL")
		result = db.Execute(query, params)
		for _, record := range result.Records {
			mapRecord, _ := record.Get("n2")
			byteData := util.MarshalData(mapRecord)
			err := json.Unmarshal(byteData, &employeeData)
			if err != nil {
				fmt.Println("Error Unmarshalling Json")
				panic(err)
			}
			employeeList = append(employeeList, employeeData)
		}
		context.JSON(http.StatusOK, employeeList)
		employeeList = nil
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

func listToolComesUnderDepartment(context *gin.Context) {
	var toolComesUnderDepartment model.ToolComesUnderDepartment
	err := context.ShouldBindJSON(&toolComesUnderDepartment)

	var toolData model.Tool
	var departmentData model.Department

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}

	retrieveNode := context.Param("node")
	var result *neo4j.EagerResult
	if retrieveNode == "Tool" {
		query := toolComesUnderDepartment.MakeQuery("LIST_TOOLS_COMES_UNDER_DEPARTMENT")
		params := toolComesUnderDepartment.MakeParams("LIST_TOOLS_COMES_UNDER_DEPARTMENT")
		result = db.Execute(query, params)
		for _, record := range result.Records {
			mapRecord, _ := record.Get("n1")
			byteData := util.MarshalData(mapRecord)
			err := json.Unmarshal(byteData, &toolData)
			if err != nil {
				fmt.Println("Error Unmarshalling Json")
				panic(err)
			}
			toolList = append(toolList, toolData)
		}
		context.JSON(http.StatusOK, toolList)
		employeeList = nil
	}
	if retrieveNode == "Department" {
		query := toolComesUnderDepartment.MakeQuery("LIST_DEPARTMENT_MAINTAIN_TOOL")
		params := toolComesUnderDepartment.MakeParams("LIST_DEPARTMENT_MAINTAIN_TOOL")
		result = db.Execute(query, params)
		for _, record := range result.Records {
			mapRecord, _ := record.Get("n2")
			byteData := util.MarshalData(mapRecord)
			err := json.Unmarshal(byteData, &departmentData)
			if err != nil {
				fmt.Println("Error Unmarshalling Json")
				panic(err)
			}
			departmentList = append(departmentList, departmentData)
		}
		context.JSON(http.StatusOK, departmentList)
		departmentList = nil
	}

}
