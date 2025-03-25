package commonRepository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DhakshidMurali/tara/db"
	"github.com/DhakshidMurali/tara/model/common"
	"github.com/DhakshidMurali/tara/util"
	"github.com/gin-gonic/gin"
)

func ListTool(context *gin.Context) {
	var tool model.Tool
	var data model.Tool

	query := tool.MakeQuery("LIST_TOOL")
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
		toolList = append(toolList, data)
	}
	context.JSON(http.StatusOK, toolList)
	toolList = nil
}

func ListCommunity(context *gin.Context) {
	var community model.Community
	var data model.Community

	query := community.MakeQuery("LIST_COMMUNITY")
	fmt.Println(query)
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
		communityList = append(communityList, data)
	}
	context.JSON(http.StatusOK, communityList)
	communityList = nil
}

func ListDomain(context *gin.Context) {
	var domain model.Domain
	var data model.Domain

	query := domain.MakeQuery("LIST_DOMAIN")
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
		domainList = append(domainList, data)
	}
	context.JSON(http.StatusOK, domainList)
	domainList = nil
}

func UpdateUser(node *model.UpdateNode, user *model.User, context *gin.Context) {

	query := model.Node{
		NodeName: "User",
	}.MakeQuery()

	params := user.MakeParams(node.Key, "UPDATE")
	db.Execute(query, params)
}

func UpdateTool(node *model.UpdateNode, tool *model.Tool, context *gin.Context) {
	query := model.Node{
		NodeName: "Tool",
	}.MakeQuery()

	params := tool.MakeParams(node.Key, "UPDATE")

	db.Execute(query, params)
}

func UpdateDomain(node *model.UpdateNode, domain *model.Domain, context *gin.Context) {
	query := model.Node{
		NodeName: "Domain",
	}.MakeQuery()

	params := domain.MakeParams(node.Key, "UPDATE")

	db.Execute(query, params)
}

func UpdateCommunity(node *model.UpdateNode, community *model.Community, context *gin.Context) {
	query := model.Node{
		NodeName: "Community",
	}.MakeQuery()

	params := community.MakeParams(node.Key, "UPDATE")

	db.Execute(query, params)
}

func UpdateCommunication(node *model.UpdateNode, communication *model.Communication, context *gin.Context) {
	query := model.Node{
		NodeName: "Communication",
	}.MakeQuery()

	params := communication.MakeParams(node.Key, "UPDATE")

	db.Execute(query, params)
}

func UpdateSkills(node *model.UpdateNode, skills *model.Skills, context *gin.Context) {
	query := model.Node{
		NodeName: "Skills",
	}.MakeQuery()

	params := skills.MakeParams(node.Key, "UPDATE")

	db.Execute(query, params)
}
