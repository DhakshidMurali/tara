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

func listEmployeeCollaboratedWithEmployee(context *gin.Context) {
	var employeeCollaboratedWithEmployee model.EmployeeCollaboratedWithEmployee
	var data model.Employee
	err := context.ShouldBindJSON((&employeeCollaboratedWithEmployee))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := employeeCollaboratedWithEmployee.MakeQuery("LIST_EMPLOYEE_COLLABORATED_WITH")
	params := employeeCollaboratedWithEmployee.MakeParams("LIST_EMPLOYEE_COLLABORATED_WITH")
	result := db.Execute(query, params)

	for _, record := range result.Records {
		mapRecord, _ := record.Get("n2")
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &data)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}
		employeeList = append(employeeList, data)
	}
	context.JSON(http.StatusOK, employeeList)
	employeeList = nil
}
func listEmployeeWorksInTools(context *gin.Context) {
	var employeeWorksInTools model.EmployeeWorksInTools
	var employeeData model.Employee
	var toolData model.Tool
	err := context.ShouldBindJSON((&employeeWorksInTools))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	retrieveNode := context.Param("node")
	var result *neo4j.EagerResult

	if retrieveNode == "Employee" {
		query := employeeWorksInTools.MakeQuery("LIST_EMPLOYEES_WORKS_IN_TOOL")
		params := employeeWorksInTools.MakeParams("LIST_EMPLOYEES_WORKS_IN_TOOL")
		result = db.Execute(query, params)

		for _, record := range result.Records {
			mapRecord, _ := record.Get("n1")
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
	if retrieveNode == "Tool" {
		query := employeeWorksInTools.MakeQuery("LIST_TOOLS_EMPLOYEE_WORKS_IN")
		params := employeeWorksInTools.MakeParams("LIST_TOOLS_EMPLOYEE_WORKS_IN")
		result = db.Execute(query, params)

		for _, record := range result.Records {
			mapRecord, _ := record.Get("n2")
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
}

func listEmployeeSkilledInSkills(context *gin.Context) {
	var employeeSkilledInSkills model.EmployeeSkilledInSkills
	var employeeData model.Employee
	var skillsData model.Skills
	err := context.ShouldBindJSON((&employeeSkilledInSkills))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	retrieveNode := context.Param("node")
	var result *neo4j.EagerResult
	if retrieveNode == "Employee" {
		query := employeeSkilledInSkills.MakeQuery("LIST_EMPLOYEE_SKILLED_IN")
		params := employeeSkilledInSkills.MakeParams("LIST_EMPLOYEE_SKILLED_IN")
		result = db.Execute(query, params)

		for _, record := range result.Records {
			mapRecord, _ := record.Get("n1")
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
	if retrieveNode == "Skills" {
		query := employeeSkilledInSkills.MakeQuery("LIST_SKILLS_SKILLED_BY_EMPLOYEE")
		params := employeeSkilledInSkills.MakeParams("LIST_SKILLS_SKILLED_BY_EMPLOYEE")
		result = db.Execute(query, params)

		for _, record := range result.Records {
			mapRecord, _ := record.Get("n2")
			byteData := util.MarshalData(mapRecord)
			err := json.Unmarshal(byteData, &skillsData)
			if err != nil {
				fmt.Println("Error Unmarshalling Json")
				panic(err)
			}
			skillsList = append(skillsList, skillsData)
		}
		context.JSON(http.StatusOK, skillsList)
		skillsList = nil
	}
}

func listEmployeeReportToEmployee(context *gin.Context) {
	var employeeReportToEmployee model.EmployeeReportToEmployee
	var data model.Employee
	err := context.ShouldBindJSON(&employeeReportToEmployee)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	retrieveNode := context.Param("node")
	var result *neo4j.EagerResult

	if retrieveNode == "MANAGER" {
		query := employeeReportToEmployee.MakeQuery("LIST_MANAGER_HIERARCHY_OF_EMPLOYEE")
		params := employeeReportToEmployee.MakeParams("LIST_MANAGER_HIERARCHY_OF_EMPLOYEE")
		result = db.Execute(query, params)

		for _, record := range result.Records {
			mapRecord, _ := record.Get("n2")
			byteData := util.MarshalData(mapRecord)
			err := json.Unmarshal(byteData, &data)
			if err != nil {
				fmt.Println("Error Unmarshalling Json")
				panic(err)
			}
			employeeList = append(employeeList, data)
		}
		context.JSON(http.StatusOK, employeeList)
		employeeList = nil
	}
	if retrieveNode == "EMPLOYEE" {
		query := employeeReportToEmployee.MakeQuery("LIST_EMPLOYEE_REPORTEE_OF_MANAGER")
		params := employeeReportToEmployee.MakeParams("LIST_EMPLOYEE_REPORTEE_OF_MANAGER")
		result = db.Execute(query, params)

		for _, record := range result.Records {
			mapRecord, _ := record.Get("n1")
			byteData := util.MarshalData(mapRecord)
			err := json.Unmarshal(byteData, &data)
			if err != nil {
				fmt.Println("Error Unmarshalling Json")
				panic(err)
			}
			employeeList = append(employeeList, data)
		}
		context.JSON(http.StatusOK, employeeList)
		employeeList = nil
	}
}
