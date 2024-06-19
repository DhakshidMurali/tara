package db

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var ctx context.Context
var driver neo4j.DriverWithContext
func Init() {
	ctx := context.Background()
	dbUri := "bolt://localhost:7687"
	dbUser := "arjun"
	dbPassword := "arjun@23"
	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))
	if err!=nil{
		panic("Error in Creating Context!")
	}
	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic("Error in Connecting to Database!")
	}
}

func Execute(query string) (*neo4j.EagerResult) {
	result, _ := neo4j.ExecuteQuery(ctx, driver,query,
		map[string]any{}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))
	return result;
}
