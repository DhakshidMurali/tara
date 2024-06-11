package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type output struct {
	Properties string `json:"properties"`
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
		match (n1:PERSON{name:'Arjun'})

return n1{
    properties:n1{.*}
}
		`,
		map[string]any{}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))
	// Loop through results and do something with them
	fmt.Println(result.Records)
	for _, record := range result.Records {
		// if len(result.Data) > 0 {
		// 	// Here are your phones.
		// 	for _, d := range result.Data {
		// 		fmt.Println(d.Row)
		// 	}
		// }
		// fmt.Println(len(record.AsMap()))
		// fmt.Println(record.Get("n"))
		fmt.Println(record.Keys)
		// fmt.Println(record.Values)
		// fmt.Println(record.AsMap())
		mapRecord, _ := record.Get("n1")
		jsonData, err := json.Marshal(mapRecord)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		var result interface{}
		err = json.Unmarshal(jsonData, &result)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		personMap, ok := result.(map[string]interface{})
		if !ok {
			fmt.Println("Error: Couldn't convert to map[string]interface{}")
			return
		}
		fmt.Println(personMap)
		prop, nameOk := personMap["properties"].(map[string]interface{})
		if !nameOk {
			fmt.Println("Error: Couldn't extract name and age")
			return
		}
		fmt.Println(prop)
		fmt.Println(prop["name"])
		if !nameOk {
			fmt.Println("Error: Couldn't extract name and age")
			return
		}
		// fmt.Println(name)
		// var person output
		// jsonData := []byte(mapRecord)
		// err := json.Unmarshal(jsonData, &person)
		// fmt.Println(reflect.TypeOf(mapRecord).Elem())
		// fmt.Println(reflect.TypeOf(record).Key())
	}

	// Summary information
	fmt.Printf("The query `%v` returned %v records in %+v.\n",
		result.Summary.Query().Text(), len(result.Records),
		result.Summary.ResultAvailableAfter())
	defer driver.Close(ctx)

}
