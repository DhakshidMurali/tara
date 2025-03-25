package commonService

import (
	"encoding/json"
	"net/http"

	"github.com/DhakshidMurali/tara/model/common"
	repository "github.com/DhakshidMurali/tara/repository/common"
	"github.com/gin-gonic/gin"
)

var toolList = []model.Tool{}
var domainList = []model.Domain{}

func ListTool(context *gin.Context) {
	repository.ListTool(context)
}

func ListCommunity(context *gin.Context) {
	repository.ListCommunity(context)
}

func ListDomain(context *gin.Context) {
	repository.ListDomain(context)
}

func UpdateUser(context *gin.Context) {
	var node model.UpdateNode
	var user model.User
	err := context.ShouldBindJSON(&node)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
		return
	}
	jsonData, err := json.Marshal(node.Node.(map[string]interface{}))

	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(jsonData, &user); err != nil {
		panic(err)
	}
	repository.UpdateUser(&node, &user, context)

}

func UpdateTool(context *gin.Context) {
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
	repository.UpdateTool(&node, &tool, context)
}

func UpdateDomain(context *gin.Context) {
	var node model.UpdateNode
	var domain model.Domain
	err := context.ShouldBindJSON(&node)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
		return
	}

	jsonData, err := json.Marshal(node.Node.(map[string]interface{}))

	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(jsonData, &domain); err != nil {
		panic(err)
	}
	repository.UpdateDomain(&node, &domain, context)

}

func UpdateCommunity(context *gin.Context) {
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
	repository.UpdateCommunity(&node, &community, context)

}

func UpdateCommunication(context *gin.Context) {
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
	repository.UpdateCommunication(&node, &communication, context)

}

func UpdateSkills(context *gin.Context) {
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
	repository.UpdateSkills(&node, &skills, context)

}
