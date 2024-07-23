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

func listTool(context *gin.Context) {
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

func listCommunity(context *gin.Context) {
	var community model.Community
	var data model.Community

	query := community.MakeQuery("LIST_COMMUNITY")
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
	context.JSON(http.StatusOK, community)
	communityList = nil
}

func listDepartment(context *gin.Context) {
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
func updateEmployee(context *gin.Context) {
	var node model.UpdateNode
	var employee model.Employee
	err := context.ShouldBindJSON(&node)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
		return
	}

	jsonData, err := json.Marshal(node.Node.(map[string]interface{}))

	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(jsonData, &employee); err != nil {
		panic(err)
	}

	query := model.Node{
		NodeName: "Hello",
	}.MakeQuery()

	params := employee.MakeParams(node.Key, "UPDATE")

	fmt.Println(query)
	fmt.Println(params)

	db.Execute(query, params)
}

func updateTool(context *gin.Context) {
	var node model.UpdateNode
	var tool model.Tool
	err := context.ShouldBindJSON(&node)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
		return
	}

	jsonData, err := json.Marshal(node.Node.(map[string]interface{}))

	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(jsonData, &tool); err != nil {
		panic(err)
	}

	query := model.Node{
		NodeName: "Tool",
	}.MakeQuery()

	params := tool.MakeParams(node.Key, "UPDATE")

	db.Execute(query, params)
}

func updateDepartment(context *gin.Context) {
	var node model.UpdateNode
	var department model.Department
	err := context.ShouldBindJSON(&node)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
		return
	}

	jsonData, err := json.Marshal(node.Node.(map[string]interface{}))

	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(jsonData, &department); err != nil {
		panic(err)
	}

	query := model.Node{
		NodeName: "Department",
	}.MakeQuery()

	params := department.MakeParams(node.Key, "UPDATE")

	db.Execute(query, params)

}

func updateCommunity(context *gin.Context) {
	var node model.UpdateNode
	var community model.Community
	err := context.ShouldBindJSON(&node)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
		return
	}

	jsonData, err := json.Marshal(node.Node.(map[string]interface{}))

	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(jsonData, &community); err != nil {
		panic(err)
	}

	query := model.Node{
		NodeName: "Community",
	}.MakeQuery()

	params := community.MakeParams(node.Key, "UPDATE")

	db.Execute(query, params)

}

func updateCommunication(context *gin.Context) {
	var node model.UpdateNode
	var communication model.Communication
	err := context.ShouldBindJSON(&node)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
		return
	}

	jsonData, err := json.Marshal(node.Node.(map[string]interface{}))

	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(jsonData, &communication); err != nil {
		panic(err)
	}

	query := model.Node{
		NodeName: "Communication",
	}.MakeQuery()

	params := communication.MakeParams(node.Key, "UPDATE")

	db.Execute(query, params)

}
func updateSkills(context *gin.Context) {
	var node model.UpdateNode
	var skills model.Skills
	err := context.ShouldBindJSON(&node)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
		return
	}

	jsonData, err := json.Marshal(node.Node.(map[string]interface{}))

	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(jsonData, &skills); err != nil {
		panic(err)
	}

	query := model.Node{
		NodeName: "Skills",
	}.MakeQuery()

	params := skills.MakeParams(node.Key, "UPDATE")

	db.Execute(query, params)

}
