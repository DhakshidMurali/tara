package util

import "strings"

func ReplaceQuery(query string, mapData map[string]string) string {
	for key, value := range mapData {
		query = strings.Replace(query,key,value,-1)
	}
	return query
}