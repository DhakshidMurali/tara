package commonRepository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DhakshidMurali/tara/db"
	"github.com/DhakshidMurali/tara/model/common"
	"github.com/DhakshidMurali/tara/util"
	"github.com/gin-gonic/gin"
)

var communityList = []model.Community{}

func CreateCommunityMemberUser(communityMemberUser *model.CommunityMemberUser, context *gin.Context) {
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

func CreateCommunityCreatedByUser(communityCreatedByUser *model.CommunityCreatedByUser, context *gin.Context) {
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

func ListCommunityCreatedByUser(communityCreatedByUser *model.CommunityCreatedByUser, context *gin.Context) {
	var data model.Community
	query := communityCreatedByUser.MakeQuery("LIST_COMMUNITY_CREATED_BY_USER")
	params := communityCreatedByUser.MakeParams("LIST_COMMUNITY_CREATED_BY_USER")
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
	context.JSON(http.StatusOK, communityList)
	communityList = nil
}

func ListCommunityMemberUser(communityMemberUser *model.CommunityMemberUser, context *gin.Context) {
	var data model.User
	query := communityMemberUser.MakeQuery("LIST_USERS_MEMBER_OF_COMMUNITY")
	params := communityMemberUser.MakeParams("LIST_USERS_MEMBER_OF_COMMUNITY")
	result := db.Execute(query, params)

	for _, record := range result.Records {
		mapRecord, _ := record.Get("n2")
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &data)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}
		userList = append(userList, data)
	}
	context.JSON(http.StatusOK, userList)
	userList = nil
}
