package domainModel

type DashboardPayload struct {
	ToolCountByDomain             []ToolCountByDomain     `json:"toolCountByDomain,omitempty"`
	DomainDetails                 []DomainDetails         `json:"domainDetails,omitempty"`
	DomainGroupByToolDeliveryType []DeliveryStatsByDomain `json:"domainGroupByToolDeliveryType,omitempty"`
}

type ToolCountByDomain struct {
	DomainName string `json:"domainName,omitempty"`
	ToolsCount int    `json:"toolsCount,omitempty"`
}

type DomainDetails struct {
	DomainName string   `json:"domainName,omitempty"`
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

func (v DomainDetails) GetQuery(queryConst string) string {
	switch queryConst {
	case "GET_DOMAIN_AND_SUBDOMAINS_DETAILS":
		return GET_DOMAIN_AND_SUBDOMAINS_DETAILS
	default:
		return ""
	}
}

func (v ToolCountByDomain) GetQuery(queryConst string) string {
	switch queryConst {
	case "GET_TOP_DOMAINS_ORDER_BY_TOOLS":
		return GET_TOP_DOMAINS_ORDER_BY_TOOLS
	default:
		return ""
	}
}
