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

var communicationList = []model.Communication{}

func CreateCommunicationPostedInCommunity(communicationPostedInCommunity *model.CommunicationPostedInCommunity, context *gin.Context) {
	query := communicationPostedInCommunity.MakeQuery("CREATE")
	params := communicationPostedInCommunity.MakeParams("CREATE")
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

func CreateCommunicationPostedByUser(communicationPostedByUser *model.CommunicationPostedByUser, context *gin.Context) {
	query := communicationPostedByUser.MakeQuery("CREATE")
	params := communicationPostedByUser.MakeParams("CREATE")
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

func ListCommunicationPostedInCommunity(communicationPostedInCommunity *model.CommunicationPostedInCommunity, context *gin.Context) {
	var data model.Communication
	query := communicationPostedInCommunity.MakeQuery("LIST_COMMUNICATION_POSTED_IN_COMMUNITY")
	params := communicationPostedInCommunity.MakeParams("LIST_COMMUNICATION_POSTED_IN_COMMUNITY")
	result := db.Execute(query, params)

	for _, record := range result.Records {
		mapRecord, _ := record.Get("n1")
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &data)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}
		communicationList = append(communicationList, data)
	}
	context.JSON(http.StatusOK, communicationList)
	communicationList = nil
}

func ListCommunicationPostedByUser(communicationPostedByUser *model.CommunicationPostedByUser, context *gin.Context) {
	var data model.Communication
	query := communicationPostedByUser.MakeQuery("LIST_COMMUNICATION_POSTED_BY_USER")
	params := communicationPostedByUser.MakeParams("LIST_COMMUNICATION_POSTED_BY_USER")
	result := db.Execute(query, params)

	for _, record := range result.Records {
		mapRecord, _ := record.Get("n1")
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &data)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}
		communicationList = append(communicationList, data)
	}
	context.JSON(http.StatusOK, communicationList)
	communicationList = nil
}
