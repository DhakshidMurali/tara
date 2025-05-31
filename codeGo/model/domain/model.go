package domainModel

type DashboardPayload struct {
	ToolCountByDomain             []ToolCountByDomain     `json:"toolCountByDomain,omitempty"`
	DomainDetails                 []DomainDetails         `json:"domainDetails,omitempty"`
	DomainGroupByToolDeliveryType []DeliveryStatsByDomain `json:"domainGroupByToolDeliveryType,omitempty"`
}

type ToolCountByDomain struct {
	DomainName string `json:"domainName,omitempty"`
	ToolCount  int    `json:"toolCount,omitempty"`
}

type DomainDetails struct {
	Domain     string   `json:"domain,omitempty"`
	SubDomains []string `json:"subDomains,omitempty"`
}

type DeliveryStatsByDomain struct {
	Domain       string            `json:"domain,omitempty"`
	DeliveryType []DeliveryDetails `json:"deliveryType,omitempty"`
}

type DeliveryDetails struct {
	DeliveryType         string `json:"deliveryType,omitempty"`
	CountPerDeliveryType []int  `json:"countPerDeliveryType,omitempty"`
}

/*
Learning :

	Here we are using an same common type struct, this struct referred by all Api ( Excluding Common API ) used for Dashboard.
	So, we need fields which are returned from database alone to be populated by struct and returned for API
	Say for Example : If API returns DomainName alone, then only domainname field to be returned as API response, no other fields included
	We added omitempty to make fields which are not populated to be omitted
*/
func (v DashboardPayload) GetQuery(queryConst string) string {
	switch queryConst {
	case "GET_LIST_TOOLS_GROUPBY_DOMAIN":
		return GET_LIST_TOOLS_GROUPBY_DOMAIN
	case "GET_LIST_TOOLS_GROUPBY_DELIVERYFORMAT_FOR_TOP4DOMAINS":
		return GET_LIST_TOOLS_GROUPBY_DELIVERYFORMAT_FOR_TOP4DOMAINS
	default:
		return ""
	}
}
