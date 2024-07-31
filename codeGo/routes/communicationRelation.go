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

func createCommunicationPostedInCommunity(context *gin.Context) {
	var communicationPostedInCommunity model.CommunicationPostedInCommunity
	err := context.ShouldBindJSON(&communicationPostedInCommunity)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
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

func createCommunicationPostedByEmployee(context *gin.Context) {
	var communicationPostedByEmployee model.CommunicationPostedByEmployee
	err := context.ShouldBindJSON(&communicationPostedByEmployee)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := communicationPostedByEmployee.MakeQuery("CREATE")
	params := communicationPostedByEmployee.MakeParams("CREATE")
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

func listCommunicationPostedInCommunity(context *gin.Context) {
	var communicationPostedInCommunity model.CommunicationPostedInCommunity
	var data model.Communication
	err := context.ShouldBindJSON(&communicationPostedInCommunity)

	if err != nil {
		panic(err)
	}

	query := communicationPostedInCommunity.MakeQuery("LIST_COMMUNICATION_POSTED_IN_COMMUNITY")
	params := communicationPostedInCommunity.MakeParams("LIST_COMMUNICATION_POSTED_IN_COMMUNITY")
	result := db.Execute(query, params)

	for _, record := range result.Records {
		mapRecord, _ := record.Get("n1")
		byteData := util.MarshalData(mapRecord)
		err = json.Unmarshal(byteData, &data)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}
		communicationList = append(communicationList, data)
	}
	context.JSON(http.StatusOK, communicationList)
	communicationList = nil
}

func listCommunicationPostedByEmployee(context *gin.Context) {
	var communicationPostedByEmployee model.CommunicationPostedByEmployee
	var data model.Communication
	err := context.ShouldBindJSON(&communicationPostedByEmployee)

	if err != nil {
		panic(err)
	}

	query := communicationPostedByEmployee.MakeQuery("LIST_COMMUNICATION_POSTED_BY_EMPLOYEE")
	fmt.Println(query)
	params := communicationPostedByEmployee.MakeParams("LIST_COMMUNICATION_POSTED_BY_EMPLOYEE")
	result := db.Execute(query, params)

	for _, record := range result.Records {
		mapRecord, _ := record.Get("n1")
		byteData := util.MarshalData(mapRecord)
		err = json.Unmarshal(byteData, &data)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}
		communicationList = append(communicationList, data)
	}
	context.JSON(http.StatusOK, communicationList)
	communicationList = nil
}
