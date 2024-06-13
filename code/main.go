package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type nodeProp struct {
	Name string `json:"name"`
	// class string `json:"class"`
	// labels []string `json:"labels"`
}

func main() {
	ctx := context.Background()
	dbUri := "bolt://localhost:7687"
	dbUser := "arjun"
	dbPassword := "arjun@23"
	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}
	result, _ := neo4j.ExecuteQuery(ctx, driver,
		`
  match (arjun:PERSON{Id:'2'})
  RETURN arjun
  {
      name: arjun.name
  }
		`,

		map[string]any{}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))
		for _, record := range result.Records {
		fmt.Println(record.Get("arjun"))
		mapRecord, _ := record.Get("arjun")
		recordMap, ok := mapRecord.(map[string]interface{})
		if !ok {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}
		jsonData, err := json.Marshal(recordMap)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		var data nodeProp
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}
		fmt.Printf("%+v\n", data)
		fmt.Println(data.Name)
	}
	defer driver.Close(ctx)

}
