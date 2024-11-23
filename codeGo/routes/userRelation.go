package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DhakshidMurali/tara/db"
	"github.com/DhakshidMurali/tara/model"
	"github.com/DhakshidMurali/tara/util"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func createUserCollaboratedWithUser(context *gin.Context) {
	var userCollaboratedWithUser model.UserCollaboratedWithUser
	err := context.ShouldBindJSON((&userCollaboratedWithUser))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	query := userCollaboratedWithUser.MakeQuery("CREATE")
	fmt.Println(query)
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

func createUserWorksInTools(context *gin.Context) {
	var userWorksInTools model.UserWorksInTools
	err := context.ShouldBindJSON((&userWorksInTools))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
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

func createUserSkilledInSkills(context *gin.Context) {
	var userSkilledInSkills model.UserSkilledInSkills
	err := context.ShouldBindJSON((&userSkilledInSkills))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
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

func listUserCollaboratedWithUser(context *gin.Context) {
	var userCollaboratedWithUser model.UserCollaboratedWithUser
	var data model.User
	err := context.ShouldBindJSON((&userCollaboratedWithUser))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
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
func listUserWorksInTools(context *gin.Context) {
	var userWorksInTools model.UserWorksInTools
	var userData model.User
	var toolData model.Tool
	err := context.ShouldBindJSON((&userWorksInTools))
	fmt.Println(err)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	retrieveNode := context.Param("node")
	var result *neo4j.EagerResult

	if retrieveNode == "USER" {
		query := userWorksInTools.MakeQuery("LIST_USERS_WORKS_IN_TOOL")
		params := userWorksInTools.MakeParams("LIST_USERS_WORKS_IN_TOOL")
		result = db.Execute(query, params)

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
	if retrieveNode == "TOOL" {
		query := userWorksInTools.MakeQuery("LIST_TOOLS_USER_WORKS_IN")
		params := userWorksInTools.MakeParams("LIST_TOOLS_USER_WORKS_IN")
		result = db.Execute(query, params)

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
}

func listUserSkilledInSkills(context *gin.Context) {
	var userSkilledInSkills model.UserSkilledInSkills
	var userData model.User
	var skillsData model.Skills
	err := context.ShouldBindJSON((&userSkilledInSkills))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Description": "Received data can't be parsed"})
	}
	retrieveNode := context.Param("node")
	var result *neo4j.EagerResult
	if retrieveNode == "USER" {
		query := userSkilledInSkills.MakeQuery("LIST_USER_SKILLED_IN")
		params := userSkilledInSkills.MakeParams("LIST_USER_SKILLED_IN")
		result = db.Execute(query, params)

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
	if retrieveNode == "SKILLS" {
		query := userSkilledInSkills.MakeQuery("LIST_SKILLS_SKILLED_BY_USER")
		params := userSkilledInSkills.MakeParams("LIST_SKILLS_SKILLED_BY_USER")
		result = db.Execute(query, params)

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
}
