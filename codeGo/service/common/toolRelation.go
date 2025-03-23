package commonService

import (
	"net/http"

	"github.com/DhakshidMurali/tara/model"
	repository "github.com/DhakshidMurali/tara/repository/common"
	"github.com/gin-gonic/gin"
)

func CreateToolRequestByUser(context *gin.Context) {
	var toolRequestByUser model.ToolRequestByUser
	err := context.ShouldBindJSON(&toolRequestByUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	repository.CreateToolRequestByUser(&toolRequestByUser, context)

}

func CreateToolAccessToUser(context *gin.Context) {
	var toolAccessToUser model.ToolAccessToUser
	err := context.ShouldBindJSON(&toolAccessToUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	repository.CreateToolAccessToUser(&toolAccessToUser, context)

}

func CreateToolManagedByUser(context *gin.Context) {
	var toolManagedByUser model.ToolManagedByUser
	err := context.ShouldBindJSON(&toolManagedByUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	repository.CreateToolManagedByUser(&toolManagedByUser, context)

}

func CreateToolComesUnderDomain(context *gin.Context) {
	var toolComesUnderDomain model.ToolComesUnderDomain
	err := context.ShouldBindJSON(&toolComesUnderDomain)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	repository.CreateToolComesUnderDomain(&toolComesUnderDomain, context)

}

func ListToolRequestByUser(context *gin.Context) {
	var toolRequestByUser model.ToolRequestByUser
	err := context.ShouldBindJSON(&toolRequestByUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	retrieveNode := context.Param("node")
	if retrieveNode == "TOOL" {
		repository.ListToolRequestByUserByTool(&toolRequestByUser, context)
	}

	if retrieveNode == "USER" {
		repository.ListToolRequestByUserByUser(&toolRequestByUser, context)
	}
}

func ListToolAccessToUser(context *gin.Context) {
	var toolAccessToUser model.ToolAccessToUser
	err := context.ShouldBindJSON(&toolAccessToUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}

	retrieveNode := context.Param("node")
	if retrieveNode == "TOOL" {
		repository.ListToolAccessToUserByTool(&toolAccessToUser, context)
	}
	if retrieveNode == "USER" {
		repository.ListToolAccessToUserByUser(&toolAccessToUser, context)
	}
}

func ListToolManagedByUser(context *gin.Context) {
	var toolManagedByUser model.ToolManagedByUser
	err := context.ShouldBindJSON(&toolManagedByUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	retrieveNode := context.Param("node")
	if retrieveNode == "TOOL" {
		repository.ListToolManagedByUserByTool(&toolManagedByUser, context)
	}
	if retrieveNode == "USER" {
		repository.ListToolManagedByUserByToolByUser(&toolManagedByUser, context)
	}
}

func ListToolComesUnderDomain(context *gin.Context) {
	var toolComesUnderDomain model.ToolComesUnderDomain
	err := context.ShouldBindJSON(&toolComesUnderDomain)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	retrieveNode := context.Param("node")
	if retrieveNode == "TOOL" {
		repository.ListToolComesUnderDomainByTool(&toolComesUnderDomain, context)
	}
	if retrieveNode == "DOMAIN" {
		repository.ListToolComesUnderDomainByDomain(&toolComesUnderDomain, context)
	}
}
