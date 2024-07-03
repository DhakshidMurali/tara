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

var community = []model.Community{}

func createCommunityMemberEmployee(context *gin.Context) {
	var communityMemberEmployee model.CommunityMemberEmployee
	err := context.ShouldBindJSON(&communityMemberEmployee)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := communityMemberEmployee.MakeQuery("CREATE")
	params := communityMemberEmployee.MakeParams("CREATE")
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

func createCommunityCreatedByEmployee(context *gin.Context) {
	var communityCreatedByEmployee model.CommunityCreatedByEmployee
	err := context.ShouldBindJSON(&communityCreatedByEmployee)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := communityCreatedByEmployee.MakeQuery("CREATE")
	params := communityCreatedByEmployee.MakeParams("CREATE")
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

func listCommunityCreatedByEmployee(context *gin.Context) {
	var communityCreatedByEmployee model.CommunityCreatedByEmployee
	var data model.Community
	err := context.ShouldBindJSON(&communityCreatedByEmployee)

	if err != nil {
		panic(err)
	}

	query := communityCreatedByEmployee.MakeQuery("LIST_COMMUNICATION_BY_EMPLOYEE")
	params := communityCreatedByEmployee.MakeParams("LIST_COMMUNICATION_BY_EMPLOYEE")
	result := db.Execute(query, params)

	for _, record := range result.Records {
		mapRecord, _ := record.Get("n1")
		byteData := util.MarshalData(mapRecord)
		err = json.Unmarshal(byteData, &data)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}
		community = append(community, data)
	}
	context.JSON(http.StatusOK, communication)
	communication = nil
}
