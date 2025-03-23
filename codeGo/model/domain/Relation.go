package domainModel

import "github.com/DhakshidMurali/tara/queryRepository"

type DashboardFields struct {
	DomainName              string `json:"DomainName,omitempty"`
	CountToolsGroupByDomain int    `json:"CountToolsGroupByDomain,omitempty"`
	DeliveryFormat          string `json:"DeliveryFormat,omitempty"`
	DeliveryFormatCount     int    `json:"DeliveryFormatCount,omitempty"`
}

/*
Learning :
	Here we are using an same common type struct, this struct referred by all Api ( Excluding Common API ) used for Dashboard.
	So, we need fields which are returned from database alone to be populated by struct and returned for API
	Say for Example : If API returns DomainName alone, then only domainname field to be returned as API response, no other fields included
	We added omitempty to make fields which are not populated to be omitted
*/
func (v DashboardFields) MakeQuery(queryConst string) string {
	switch queryConst {
	case "GET_LIST_TOOLS_GROUPBY_DOMAIN":
		return queryRepository.GET_LIST_TOOLS_GROUPBY_DOMAIN
	case "GET_LIST_TOOLS_GROUPBY_DELIVERYFORMAT_FOR_TOP4DOMAINS":
		return queryRepository.GET_LIST_TOOLS_GROUPBY_DELIVERYFORMAT_FOR_TOP4DOMAINS
	default:
		return ""
	}
}
