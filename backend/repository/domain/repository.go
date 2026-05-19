package domainRepository

import (
	"encoding/json"
	"fmt"

	"github.com/DhakshidMurali/tara/db"
	model "github.com/DhakshidMurali/tara/model/domain"
	"github.com/DhakshidMurali/tara/util"
)

var domainDetailsSlice []model.DomainDetails
var topDomainsSlice []model.ToolCountByDomain
var deliveryStatsByDomainSlice []model.DeliveryStatsByDomain

func GetDomainDetails() ([]model.DomainDetails, error) {
	var data model.DomainDetails
	query := data.GetQuery("GET_DOMAIN_AND_SUBDOMAINS_DETAILS")
	params := map[string]any{}
	response, err := db.Execute(query, params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	domainDetailsSlice = nil
	for _, record := range response.Records {
		var domainDetailModel model.DomainDetails
		mapRecord, _ := record.Get("response")
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &domainDetailModel)
		domainDetailsSlice = append(domainDetailsSlice, domainDetailModel)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			return nil, err
		}
	}
	return domainDetailsSlice, nil
}

func GetTopDomainsByTool() ([]model.ToolCountByDomain, error) {
	var data model.ToolCountByDomain
	query := data.GetQuery("GET_TOP_DOMAINS_ORDER_BY_TOOLS")
	params := map[string]any{}
	response, err := db.Execute(query, params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	topDomainsSlice = nil
	for _, record := range response.Records {
		var toolCountByDomain model.ToolCountByDomain
		mapRecord, _ := record.Get("response")
		fmt.Println(record)
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &toolCountByDomain)
		topDomainsSlice = append(topDomainsSlice, toolCountByDomain)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			return nil, err
		}

	}
	return topDomainsSlice, nil
}

func GetDeliveryFormatByDomains() ([]model.DeliveryStatsByDomain, error) {
	var data model.DeliveryStatsByDomain
	query := data.GetQuery("GET_DELIVERY_FORMAT_GROUP_BY_DOMAINS")
	params := map[string]any{}
	response, err := db.Execute(query, params)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	deliveryStatsByDomainSlice = nil
	for _, record := range response.Records {
		var deliveryStatsByDomain model.DeliveryStatsByDomain
		mapRecord, _ := record.Get("response")
		fmt.Println(record)
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &deliveryStatsByDomain)
		deliveryStatsByDomainSlice = append(deliveryStatsByDomainSlice, deliveryStatsByDomain)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			return nil, err
		}
	}
	return deliveryStatsByDomainSlice, nil
}
