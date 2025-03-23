package main

import (
	"strconv"

	"github.com/DhakshidMurali/tara/db"
)

// type nodeProp struct {
// 	Name string `json:"name"`
// 	// class string `json:"class"`
// 	// labels []string `json:"labels"`
// }

func main() {

	db.Init()

	for i := 20; i < 100; i++ {
		var applicationName string = "Perform Application Matters " + strconv.Itoa(i)
		var query string = `
		MERGE (n1:COMMUNICATION{
			CONTENT:'New content'
		})
		MERGE (n2:TOOL{
				NAME:$applicationName,
				DELIVERYFORMAT:'Open Source'
			})
		MERGE (n1)-[r:RELATED_TO]->(n2)`

		db.Execute(query,
			map[string]any{"applicationName": applicationName})

	}

	// server := gin.Default()

	// routes.APIRoutes(server)

	// server.Run(":8080")

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
