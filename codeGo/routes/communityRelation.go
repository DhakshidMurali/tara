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

func createCommunityMemberUser(context *gin.Context) {
	var communityMemberUser model.CommunityMemberUser
	err := context.ShouldBindJSON(&communityMemberUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := communityMemberUser.MakeQuery("CREATE")
	params := communityMemberUser.MakeParams("CREATE")
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

func createCommunityCreatedByUser(context *gin.Context) {
	var communityCreatedByUser model.CommunityCreatedByUser
	err := context.ShouldBindJSON(&communityCreatedByUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := communityCreatedByUser.MakeQuery("CREATE")
	params := communityCreatedByUser.MakeParams("CREATE")
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

func listCommunityCreatedByUser(context *gin.Context) {
	var communityCreatedByUser model.CommunityCreatedByUser
	var data model.Community
	err := context.ShouldBindJSON(&communityCreatedByUser)

	if err != nil {
		panic(err)
	}

	query := communityCreatedByUser.MakeQuery("LIST_COMMUNITY_CREATED_BY_USER")
	params := communityCreatedByUser.MakeParams("LIST_COMMUNITY_CREATED_BY_USER")
	result := db.Execute(query, params)

	for _, record := range result.Records {
		mapRecord, _ := record.Get("n1")
		byteData := util.MarshalData(mapRecord)
		err = json.Unmarshal(byteData, &data)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}
		communityList = append(communityList, data)
	}
	context.JSON(http.StatusOK, communityList)
	communityList = nil
}

func listCommunityMemberUser(context *gin.Context) {
	var communityMemberUser model.CommunityMemberUser
	var data model.User
	err := context.ShouldBindJSON(&communityMemberUser)

	if err != nil {
		panic(err)
	}

	query := communityMemberUser.MakeQuery("LIST_USERS_MEMBER_OF_COMMUNITY")
	params := communityMemberUser.MakeParams("LIST_USERS_MEMBER_OF_COMMUNITY")
	result := db.Execute(query, params)

	for _, record := range result.Records {
		mapRecord, _ := record.Get("n2")
		byteData := util.MarshalData(mapRecord)
		err = json.Unmarshal(byteData, &data)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}
		userList = append(userList, data)
	}
	context.JSON(http.StatusOK, userList)
	userList = nil
}
