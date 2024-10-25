package main

import (
	"encoding/json"
	"fmt"

	"github.com/DhakshidMurali/tara/db"
	"github.com/DhakshidMurali/tara/model"
	"github.com/DhakshidMurali/tara/routes"
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

	// server.Run(":8080")

	result := db.Execute(`
			Match (n1:Employee)-[r:ComesUnder]->(n2:Department)
WITH count(n1) as employeeCount, n2.DepartmentName as departmentName
RETURN {
    DepartmentName:departmentName,
    EmployeeCount:employeeCount
} as data`, map[string]any{})
	for _, record := range result.Records {
		// fmt.Println(record.Values...)
		mapRecord, _ := record.Get("data")
		fmt.Println(mapRecord)
		recordMap, ok := mapRecord.(map[string]interface{})
		if !ok {
			fmt.Println("Error unmarshalling JSON:")
			return
		}
		jsonData, err := json.Marshal(recordMap)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		var data model.EmployeeGroupByDepartment
		err = json.Unmarshal(jsonData, &data)
		fmt.Println(jsonData)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}
		fmt.Printf("%+v\n", data)
		fmt.Println(data.EmployeeCount)
		fmt.Println(data.DepartmentName)
	}

}

/*
* Building list employeeGroupByDepartment Api and  Create API testing for that
 */
