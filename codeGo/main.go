package main

import (
	"encoding/json"
	"fmt"

	"github.com/DhakshidMurali/tara/db"
	"github.com/DhakshidMurali/tara/model"
)

// type nodeProp struct {
// 	Name string `json:"name"`
// 	// class string `json:"class"`
// 	// labels []string `json:"labels"`
// }

func main() {

	db.Init()

	// server := gin.Default()

	// routes.APIRoutes(server)

	// server.Run(":8080")

	result := db.Execute(`
			match (n:Communication)
			return n{Type:n.Type,Content:n.Content}`, map[string]any{})
	for _, record := range result.Records {
		// fmt.Println(record.Values...)
		mapRecord, _ := record.Get("n")
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
		var data model.Communication
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}
		fmt.Printf("%+v\n", data)
		fmt.Println(data.Content)
		fmt.Println(data.Type)
	}

}
