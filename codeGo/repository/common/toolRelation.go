package commonRepository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DhakshidMurali/tara/db"
	"github.com/DhakshidMurali/tara/model"
	"github.com/DhakshidMurali/tara/util"
	"github.com/gin-gonic/gin"
)

var domainList = []model.Domain{}

func CreateToolRequestByUser(toolRequestByUser *model.ToolRequestByUser, context *gin.Context) {
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

func CreateToolAccessToUser(toolAccessToUser *model.ToolAccessToUser, context *gin.Context) {
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

func CreateToolManagedByUser(toolManagedByUser *model.ToolManagedByUser, context *gin.Context) {
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

func CreateToolComesUnderDomain(toolComesUnderDomain *model.ToolComesUnderDomain, context *gin.Context) {
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

func ListToolRequestByUserByTool(toolRequestByUser *model.ToolRequestByUser, context *gin.Context) {
	var toolData model.Tool
	query := toolRequestByUser.MakeQuery("LIST_TOOL_REQUESTED_BY_USER")
	params := toolRequestByUser.MakeParams("LIST_TOOL_REQUESTED_BY_USER")
	result := db.Execute(query, params)
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
func ListToolRequestByUserByUser(toolRequestByUser *model.ToolRequestByUser, context *gin.Context) {
	var userData model.User
	query := toolRequestByUser.MakeQuery("LIST_USERS_REQUESTED_TO_TOOL")
	params := toolRequestByUser.MakeParams("LIST_USERS_REQUESTED_TO_TOOL")
	result := db.Execute(query, params)
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

func ListToolAccessToUserByTool(toolAccessToUser *model.ToolAccessToUser, context *gin.Context) {
	var toolData model.Tool
	query := toolAccessToUser.MakeQuery("LIST_TOOLS_ACCESS_BY_USER")
	params := toolAccessToUser.MakeParams("LIST_TOOLS_ACCESS_BY_USER")
	result := db.Execute(query, params)
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
func ListToolAccessToUserByUser(toolAccessToUser *model.ToolAccessToUser, context *gin.Context) {
	var userData model.User
	query := toolAccessToUser.MakeQuery("LIST_USERS_ACCESS_TO_TOOL")
	params := toolAccessToUser.MakeParams("LIST_USERS_ACCESS_TO_TOOL")
	result := db.Execute(query, params)
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

func ListToolManagedByUserByTool(toolManagedByUser *model.ToolManagedByUser, context *gin.Context) {
	var toolData model.Tool
	query := toolManagedByUser.MakeQuery("LIST_TOOLS_MANAGED_BY_USER")
	params := toolManagedByUser.MakeParams("LIST_TOOLS_MANAGED_BY_USER")
	result := db.Execute(query, params)
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
func ListToolManagedByUserByToolByUser(toolManagedByUser *model.ToolManagedByUser, context *gin.Context) {
	var userData model.User
	query := toolManagedByUser.MakeQuery("LIST_USERS_MANAGES_TOOL")
	params := toolManagedByUser.MakeParams("LIST_USERS_MANAGES_TOOL")
	result := db.Execute(query, params)
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

func ListToolComesUnderDomainByTool(toolComesUnderDomain *model.ToolComesUnderDomain, context *gin.Context) {
	var toolData model.Tool
	query := toolComesUnderDomain.MakeQuery("LIST_TOOLS_COMES_UNDER_DOMAIN")
	params := toolComesUnderDomain.MakeParams("LIST_TOOLS_COMES_UNDER_DOMAIN")
	result := db.Execute(query, params)
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
func ListToolComesUnderDomainByDomain(toolComesUnderDomain *model.ToolComesUnderDomain, context *gin.Context) {
	var domainData model.Domain
	query := toolComesUnderDomain.MakeQuery("LIST_DOMAIN_MAINTAIN_TOOL")
	params := toolComesUnderDomain.MakeParams("LIST_DOMAIN_MAINTAIN_TOOL")
	result := db.Execute(query, params)
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
