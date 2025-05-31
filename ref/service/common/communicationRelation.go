package commonService

import (
	"net/http"

	"github.com/DhakshidMurali/tara/model/common"
	repository "github.com/DhakshidMurali/tara/repository/common"
	"github.com/gin-gonic/gin"
)

var CommunicationList = []model.Communication{}

func CreateCommunicationPostedInCommunity(context *gin.Context) {
	var communicationPostedInCommunity model.CommunicationPostedInCommunity
	err := context.ShouldBindJSON(&communicationPostedInCommunity)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	repository.CreateCommunicationPostedInCommunity(&communicationPostedInCommunity, context)

}

func CreateCommunicationPostedByUser(context *gin.Context) {
	var communicationPostedByUser model.CommunicationPostedByUser
	err := context.ShouldBindJSON(&communicationPostedByUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	repository.CreateCommunicationPostedByUser(&communicationPostedByUser, context)

}

func ListCommunicationPostedInCommunity(context *gin.Context) {
	var communicationPostedInCommunity model.CommunicationPostedInCommunity
	err := context.ShouldBindJSON(&communicationPostedInCommunity)

	if err != nil {
		panic(err)
	}
	repository.ListCommunicationPostedInCommunity(&communicationPostedInCommunity, context)

}

func ListCommunicationPostedByUser(context *gin.Context) {
	var communicationPostedByUser model.CommunicationPostedByUser
	err := context.ShouldBindJSON(&communicationPostedByUser)

	if err != nil {
		panic(err)
	}
}
