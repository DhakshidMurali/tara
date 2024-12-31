package main

import (
	"github.com/DhakshidMurali/tara/db"
	"github.com/DhakshidMurali/tara/routes"
	dashboardRoutes "github.com/DhakshidMurali/tara/routes/dashboard"
	"github.com/gin-gonic/gin"
)

// type nodeProp struct {
// 	Name string `json:"name"`
// 	// class string `json:"class"`
// 	// labels []string `json:"labels"`
// }

func main() {

	db.Init()

	server := gin.Default()

	routes.APIRoutes(server)
	dashboardRoutes.APIRoutes(server)

	server.Run(":8080")

	// result := db.Execute(`
	// 		MATCH (N1:TOOL)-[r:COMESUNDER]->(N2:DOMAIN)
	// RETURN N2.DomainName AS DomainName,COUNT(N2.DomainName) AS CountToolsGroupByDomain
	// ORDER BY CountToolsGroupByDomain DESC`, map[string]any{})
	// var data dashboardModel.ToolGroupByDomain

	// for _, record := range result.Records {
	// 	fmt.Println(record)
	// 	// byteData := util.MarshalData(record)
	// 	// err := json.Unmarshal(byteData, &data)
	// 	// if err != nil {
	// 	// 	fmt.Println("Error Unmarshalling Json")
	// 	// 	panic(err)
	// 	// }
	// 	// fmt.Println(data)
	// }

}

/*
* Building list userGroupByDomain Api and  Create API testing for that
 */
