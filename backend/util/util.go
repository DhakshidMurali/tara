package util

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ReplaceQuery(query string, mapData map[string]string) string {
	for key, value := range mapData {
		query = strings.Replace(query, key, value, -1)
	}
	return query
}

func DoubleReplaceQuery(query string, mapData map[string]string) string {
	for key, value := range mapData {
		query = strings.Replace(query, key, value, -1)
	}
	for key, value := range mapData {
		query = strings.Replace(query, key, value, -1)
	}
	return query
}

func AddToParams(value string, mapKey string, params map[string]string) {
	if value != "" {
		params[mapKey] = value
	}
}

func MarshalData(mapRecord interface{}) []byte {
	recordMap, ok := mapRecord.(map[string]interface{})
	if !ok {
		fmt.Println("Error Type Assertion")
		panic(ok)
	}
	jsonData, err := json.Marshal(recordMap)
	if err != nil {
		fmt.Println("Error Marshal Json")
		panic(err)
	}
	return jsonData
}
