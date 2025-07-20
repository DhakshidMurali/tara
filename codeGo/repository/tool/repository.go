package toolRepository

import (
	"encoding/json"
	"fmt"

	"github.com/DhakshidMurali/tara/db"
	model "github.com/DhakshidMurali/tara/model/tool"
	"github.com/DhakshidMurali/tara/util"
)

var toolDataList = []model.ToolCountByCommunication{}
var toolDetailsByDomainList = []model.ToolDetailsByDomain{}
var toolStatsByLikesList = []model.ToolsStatsByLikes{}

func GetTopToolsByCommunication() ([]model.ToolCountByCommunication, error) {
	toolDataList = nil
	var data model.ToolCountByCommunication
	query := data.MakeQuery("GET_TOP_TOOLS_ORDER_BY_COMMUNICATION")
	params := map[string]any{}
	result, err := db.Execute(query, params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, record := range result.Records {
		var toolPageDetails model.ToolCountByCommunication
		mapRecord, _ := record.Get("response")
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &toolPageDetails)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			return nil, err
		}
		toolDataList = append(toolDataList, toolPageDetails)
	}
	return toolDataList, nil
}

func GetToolByDomainDetails() ([]model.ToolDetailsByDomain, error) {
	toolDetailsByDomainList = nil
	var data model.ToolDetailsByDomain
	query := data.MakeQuery("GET_TOOL_BY_DOMAIN_AND_SUBDOMAINS_DETAILS")
	params := map[string]any{}
	result, err := db.Execute(query, params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, record := range result.Records {
		var toolDetailsByDomain model.ToolDetailsByDomain
		mapRecord, _ := record.Get("response")
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &toolDetailsByDomain)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}
		toolDetailsByDomainList = append(toolDetailsByDomainList, toolDetailsByDomain)
	}
	return toolDetailsByDomainList, nil
}

func GetLikesStatsByLikes() ([]model.ToolDetailsByDomain, error) {
	toolStatsByLikesList = nil
	var data model.ToolsStatsByLikes
	query := data.MakeQuery("GET_LIKES_STATS_BY_TOOL")
	params := map[string]any{}
	result, err := db.Execute(query, params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, record := range result.Records {
		var toolsStatsByLikes model.ToolsStatsByLikes
		mapRecord, _ := record.Get("response")
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &toolsStatsByLikes)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}
		toolStatsByLikesList = append(toolStatsByLikesList, toolsStatsByLikes)
	}
	return toolDetailsByDomainList, nil
}
