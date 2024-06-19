package main

import (
	"github.com/DhakshidMurali/tara/db"
	"github.com/DhakshidMurali/tara/routes"
	"github.com/gin-gonic/gin"
)

type nodeProp struct {
	Name string `json:"name"`
	// class string `json:"class"`
	// labels []string `json:"labels"`
}

func main() {

	db.Init();
	
	server:=gin.Default();

	routes.APIRoutes(server);

	server.Run(":8080");
	// for _, record := range result.Records {
	// 	mapRecord, _ := record.Get("arjun")
	// 	recordMap, ok := mapRecord.(map[string]interface{})
	// 	if !ok {
	// 		fmt.Println("Error unmarshalling JSON:", err)
	// 		return
	// 	}
	// 	jsonData, err := json.Marshal(recordMap)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		return
	// 	}
	// 	var data nodeProp
	// 	err = json.Unmarshal(jsonData, &data)
	// 	if err != nil {
	// 		fmt.Println("Error unmarshalling JSON:", err)
	// 		return
	// 	}
	// 	fmt.Printf("%+v\n", data)
	// 	fmt.Println(data.Name)
	// }

}
