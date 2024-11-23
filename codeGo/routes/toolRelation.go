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

func createToolRequestByUser(context *gin.Context) {
	var toolRequestByUser model.ToolRequestByUser
	err := context.ShouldBindJSON(&toolRequestByUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := toolRequestByUser.MakeQuery("CREATE")
	params := toolRequestByUser.MakeParams("CREATE")
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
func listToolRequestByUser(context *gin.Context) {
	var toolRequestByUser model.ToolRequestByUser
	err := context.ShouldBindJSON(&toolRequestByUser)

	var toolData model.Tool
	var userData model.User

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}

	retrieveNode := context.Param("node")
	var result *neo4j.EagerResult

	if retrieveNode == "TOOL" {
		query := toolRequestByUser.MakeQuery("LIST_TOOL_REQUESTED_BY_USER")
		params := toolRequestByUser.MakeParams("LIST_TOOL_REQUESTED_BY_USER")
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

	if retrieveNode == "USER" {
		query := toolRequestByUser.MakeQuery("LIST_USERS_REQUESTED_TO_TOOL")
		params := toolRequestByUser.MakeParams("LIST_USERS_REQUESTED_TO_TOOL")
		result = db.Execute(query, params)
		for _, record := range result.Records {
			mapRecord, _ := record.Get("n2")
			byteData := util.MarshalData(mapRecord)
			err := json.Unmarshal(byteData, &userData)
			if err != nil {
				fmt.Println("Error Unmarshalling Json")
				panic(err)
			}
			userList = append(userList, userData)
		}
		context.JSON(http.StatusOK, userList)
		userList = nil
	}
}
func createToolAccessToUser(context *gin.Context) {
	var toolAccessToUser model.ToolAccessToUser
	err := context.ShouldBindJSON(&toolAccessToUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := toolAccessToUser.MakeQuery("CREATE")
	params := toolAccessToUser.MakeParams("CREATE")
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
func listToolAccessToUser(context *gin.Context) {
	var toolAccessToUser model.ToolAccessToUser
	err := context.ShouldBindJSON(&toolAccessToUser)

	var toolData model.Tool
	var userData model.User

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}

	retrieveNode := context.Param("node")
	var result *neo4j.EagerResult
	if retrieveNode == "TOOL" {
		query := toolAccessToUser.MakeQuery("LIST_TOOLS_ACCESS_BY_USER")
		params := toolAccessToUser.MakeParams("LIST_TOOLS_ACCESS_BY_USER")
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
	if retrieveNode == "USER" {
		query := toolAccessToUser.MakeQuery("LIST_USERS_ACCESS_TO_TOOL")
		params := toolAccessToUser.MakeParams("LIST_USERS_ACCESS_TO_TOOL")
		result = db.Execute(query, params)
		for _, record := range result.Records {
			mapRecord, _ := record.Get("n2")
			byteData := util.MarshalData(mapRecord)
			err := json.Unmarshal(byteData, &userData)
			if err != nil {
				fmt.Println("Error Unmarshalling Json")
				panic(err)
			}
			userList = append(userList, userData)
		}
		context.JSON(http.StatusOK, userList)
		userList = nil
	}

}
func createToolManagedByUser(context *gin.Context) {
	var toolManagedByUser model.ToolManagedByUser
	err := context.ShouldBindJSON(&toolManagedByUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := toolManagedByUser.MakeQuery("CREATE")
	params := toolManagedByUser.MakeParams("CREATE")
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

func listToolManagedByUser(context *gin.Context) {
	var toolManagedByUser model.ToolManagedByUser
	err := context.ShouldBindJSON(&toolManagedByUser)

	var toolData model.Tool
	var userData model.User

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}

	retrieveNode := context.Param("node")
	var result *neo4j.EagerResult
	if retrieveNode == "TOOL" {
		query := toolManagedByUser.MakeQuery("LIST_TOOLS_MANAGED_BY_USER")
		fmt.Println(query)
		params := toolManagedByUser.MakeParams("LIST_TOOLS_MANAGED_BY_USER")
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
		userList = nil
	}
	if retrieveNode == "USER" {
		query := toolManagedByUser.MakeQuery("LIST_USERS_MANAGES_TOOL")
		params := toolManagedByUser.MakeParams("LIST_USERS_MANAGES_TOOL")
		result = db.Execute(query, params)
		for _, record := range result.Records {
			mapRecord, _ := record.Get("n2")
			byteData := util.MarshalData(mapRecord)
			err := json.Unmarshal(byteData, &userData)
			if err != nil {
				fmt.Println("Error Unmarshalling Json")
				panic(err)
			}
			userList = append(userList, userData)
		}
		context.JSON(http.StatusOK, userList)
		userList = nil
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
	if retrieveNode == "TOOL" {
		query := toolComesUnderDepartment.MakeQuery("LIST_TOOLS_COMES_UNDER_DEPARTMENT")
		fmt.Println(query)
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
		userList = nil
	}
	if retrieveNode == "DEPARTMENT" {
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
