package db

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var ctx context.Context
var driver neo4j.DriverWithContext
var err error

func Init() {
	ctx = context.Background()
	dbUri := "bolt://localhost:7687"
	dbUser := "arjun"
	dbPassword := "arjun@23"
	driver, err = neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))
	if err != nil {
		panic("Error in Creating Context!")
	}
	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic("Error in Connecting to Database!")
	}
}

func Execute(query string, params map[string]any) *neo4j.EagerResult {
	// properties := map[string]interface{}{
	// 	"name": "Alice",
	// 	"age":  31,
	// 	"city": "New York",
	// }
	// result, err := neo4j.ExecuteQuery(ctx, driver, query, params, neo4j.EagerResultTransformer,
	// 	neo4j.ExecuteQueryWithDatabase("neo4j"))
	result, err := neo4j.ExecuteQuery(ctx, driver, query, params, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		fmt.Println(err)
	}
	return result
}
