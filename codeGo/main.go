package main

import (
	"github.com/DhakshidMurali/tara/db"
	"github.com/DhakshidMurali/tara/handler"
	"github.com/gin-gonic/gin"
)

// type nodeProp struct {
// 	Name string `json:"name"`
// 	// class string `json:"class"`
// 	// labels []string `json:"labels"`
// }

func main() {

	db.Init()

	// for i := 10; i < 25; i++ {
	// 	var applicationName string = "Perform Application Matters " + strconv.Itoa(i)
	// 	var communicationCount int = 30 - i
	// 	var query string = `
	// 	MERGE (n1:TOOL{
	// 			NAME:$applicationName,
	// 			DELIVERYFORMAT:'IDE Plugin',
	// 			COMMUNICATIONCOUNT:$communicationCount
	// 		})
	// 	MERGE(n2:Domain{
	// 			DomainName:'Dev Operation'
	// 		})
	// 	MERGE (n1)-[r:COMESUNDER]->(n2)`

	// 	db.Execute(query,
	// 		map[string]any{"applicationName": applicationName, "communicationCount": communicationCount})

	// }

	server := gin.Default()

	handler.APIRoutes(server)

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
