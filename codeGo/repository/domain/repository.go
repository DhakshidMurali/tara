package domainRepository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DhakshidMurali/tara/db"
	model "github.com/DhakshidMurali/tara/model/domain"
	"github.com/DhakshidMurali/tara/util"
	"github.com/gin-gonic/gin"
)

var domainDetailsSlice []model.DomainDetails
var topDomainsSlice []model.ToolCountByDomain

func GetDomainDetails(context *gin.Context) {
	var data model.DomainDetails
	query := data.GetQuery("GET_DOMAIN_AND_SUBDOMAINS_DETAILS")
	params := map[string]any{}
	response := db.Execute(query, params)
	for _, record := range response.Records {
		mapRecord, _ := record.Get("response")
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &data)
		domainDetailsSlice = append(domainDetailsSlice, data)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}

	}
	context.JSON(http.StatusOK, domainDetailsSlice)
}

func GetTopDomainsByTool(context *gin.Context) {
	var data model.ToolCountByDomain
	query := data.GetQuery("GET_TOP_DOMAINS_ORDER_BY_TOOLS")
	params := map[string]any{}
	response := db.Execute(query, params)
	for _, record := range response.Records {
		mapRecord, _ := record.Get("response")
		fmt.Println(record)
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &data)
		topDomainsSlice = append(topDomainsSlice, data)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}

	}
	context.JSON(http.StatusOK, topDomainsSlice)
}

func GetDeliveryFormatByDomains(context *gin.Context) {
	var data model.ToolCountByDomain
	query := data.GetQuery("GET_TOP_DOMAINS_ORDER_BY_TOOLS")
	params := map[string]any{}
	response := db.Execute(query, params)
	for _, record := range response.Records {
		mapRecord, _ := record.Get("response")
		fmt.Println(record)
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &data)
		topDomainsSlice = append(topDomainsSlice, data)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}

	}
	context.JSON(http.StatusOK, topDomainsSlice)
}
