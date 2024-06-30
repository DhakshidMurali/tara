package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DhakshidMurali/tara/db"
	"github.com/DhakshidMurali/tara/model"
	"github.com/gin-gonic/gin"
)

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

	params := employee.MakeParams(node.Key)

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

	params := tool.MakeParams(node.Key)

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

	params := department.MakeParams(node.Key)

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

	params := community.MakeParams(node.Key)

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

	params := communication.MakeParams(node.Key)

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

	params := skills.MakeParams(node.Key)

	db.Execute(query, params)

}
