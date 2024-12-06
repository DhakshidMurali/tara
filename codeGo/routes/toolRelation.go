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
func createToolComesUnderDomain(context *gin.Context) {
	var toolComesUnderDomain model.ToolComesUnderDomain
	err := context.ShouldBindJSON(&toolComesUnderDomain)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := toolComesUnderDomain.MakeQuery("CREATE")
	params := toolComesUnderDomain.MakeParams("CREATE")
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

func listToolComesUnderDomain(context *gin.Context) {
	var toolComesUnderDomain model.ToolComesUnderDomain
	err := context.ShouldBindJSON(&toolComesUnderDomain)

	var toolData model.Tool
	var domainData model.Domain

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}

	retrieveNode := context.Param("node")
	var result *neo4j.EagerResult
	if retrieveNode == "TOOL" {
		query := toolComesUnderDomain.MakeQuery("LIST_TOOLS_COMES_UNDER_DOMAIN")
		fmt.Println(query)
		params := toolComesUnderDomain.MakeParams("LIST_TOOLS_COMES_UNDER_DOMAIN")
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
	if retrieveNode == "DOMAIN" {
		query := toolComesUnderDomain.MakeQuery("LIST_DOMAIN_MAINTAIN_TOOL")
		params := toolComesUnderDomain.MakeParams("LIST_DOMAIN_MAINTAIN_TOOL")
		result = db.Execute(query, params)
		for _, record := range result.Records {
			mapRecord, _ := record.Get("n2")
			byteData := util.MarshalData(mapRecord)
			err := json.Unmarshal(byteData, &domainData)
			if err != nil {
				fmt.Println("Error Unmarshalling Json")
				panic(err)
			}
			domainList = append(domainList, domainData)
		}
		context.JSON(http.StatusOK, domainList)
		domainList = nil
	}

}
