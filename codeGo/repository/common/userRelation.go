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

var skillsList = []model.Skills{}

var userList = []model.User{}

var toolList = []model.Tool{}

func CreateUserCollaboratedWithUser(userCollaboratedWithUser *model.UserCollaboratedWithUser, context *gin.Context) {
	query := userCollaboratedWithUser.MakeQuery("CREATE")
	params := userCollaboratedWithUser.MakeParams("CREATE")
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

func CreateUserWorksInTools(userWorksInTools *model.UserWorksInTools, context *gin.Context) {
	query := userWorksInTools.MakeQuery("CREATE")
	params := userWorksInTools.MakeParams("CREATE")
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

func CreateUserSkilledInSkills(userSkilledInSkills *model.UserSkilledInSkills, context *gin.Context) {
	query := userSkilledInSkills.MakeQuery("CREATE")
	params := userSkilledInSkills.MakeParams("CREATE")
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

func ListUserCollaboratedWithUser(userCollaboratedWithUser *model.UserCollaboratedWithUser, context *gin.Context) {
	var data model.User
	query := userCollaboratedWithUser.MakeQuery("LIST_USER_COLLABORATED_WITH")
	params := userCollaboratedWithUser.MakeParams("LIST_USER_COLLABORATED_WITH")
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
func ListUserWorksInToolsByUser(userWorksInTools *model.UserWorksInTools, context *gin.Context) {
	var userData model.User
	query := userWorksInTools.MakeQuery("LIST_USERS_WORKS_IN_TOOL")
	params := userWorksInTools.MakeParams("LIST_USERS_WORKS_IN_TOOL")
	result := db.Execute(query, params)

	for _, record := range result.Records {
		mapRecord, _ := record.Get("n1")
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &userData)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}
		userList = append(userList, userData)
	}
	context.JSON(http.StatusOK, userList)
	userList = nil
}
func ListUserWorksInToolsByTool(userWorksInTools *model.UserWorksInTools, context *gin.Context) {
	var toolData model.Tool
	query := userWorksInTools.MakeQuery("LIST_TOOLS_USER_WORKS_IN")
	params := userWorksInTools.MakeParams("LIST_TOOLS_USER_WORKS_IN")
	result := db.Execute(query, params)

	for _, record := range result.Records {
		mapRecord, _ := record.Get("n2")
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &toolData)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}
		toolList = append(toolList, toolData)
	}
	context.JSON(http.StatusOK, toolList)
	toolList = nil
}

func ListUserSkilledInSkillsByUser(userSkilledInSkills *model.UserSkilledInSkills, context *gin.Context) {
	query := userSkilledInSkills.MakeQuery("LIST_USER_SKILLED_IN")
	params := userSkilledInSkills.MakeParams("LIST_USER_SKILLED_IN")
	result := db.Execute(query, params)
	var userData model.User

	for _, record := range result.Records {
		mapRecord, _ := record.Get("n1")
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &userData)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}
		userList = append(userList, userData)
	}
	context.JSON(http.StatusOK, userList)
	userList = nil
}
func ListUserSkilledInSkillsBySkills(userSkilledInSkills *model.UserSkilledInSkills, context *gin.Context) {
	var skillsData model.Skills
	query := userSkilledInSkills.MakeQuery("LIST_SKILLS_SKILLED_BY_USER")
	params := userSkilledInSkills.MakeParams("LIST_SKILLS_SKILLED_BY_USER")
	result := db.Execute(query, params)

	for _, record := range result.Records {
		mapRecord, _ := record.Get("n2")
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &skillsData)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}
		skillsList = append(skillsList, skillsData)
	}
	context.JSON(http.StatusOK, skillsList)
	skillsList = nil

}
