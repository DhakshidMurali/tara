package toolModel

type ToolCountByCommunication struct {
	ToolName           string `json:"toolName,omitempty"`
	SubDomain          string `json:"subDomain,omitempty"`
	DomainName         string `json:"domainName,omitempty"`
	CommunicationCount int    `json:"communicationCount,omitempty"`
}

type ToolDetailsByDomain struct {
	DomainName string `json:"domainName,omitempty"`
	SubDomain  string `json:"subDomain,omitempty"`
	ToolName   string `json:"toolName,omitempty"`
	Likes      int    `json:"likes,omitempty"`
}

type ToolDetailsByDomainPayload struct {
	DomainName string      `json:"domainName,omitempty"`
	SubDomain  []SubDomain `json:"subDomain,omitempty"`
}

type SubDomain struct {
	SubDomainName string   `json:"subDomainName,omitempty"`
	Tool          []string `json:"tool,omitempty"`
}

type ToolsStatsByLikes struct {
	ToolName string `json:"toolName,omitempty"`
	Likes    int    `json:"likes,omitempty"`
}

func (v ToolCountByCommunication) MakeQuery(queryConst string) string {
	switch queryConst {
	case "GET_TOP_TOOLS_ORDER_BY_COMMUNICATION":
		return GET_TOP_TOOLS_ORDER_BY_COMMUNICATION
	default:
		return ""
	}
}

func (v ToolsStatsByLikes) MakeQuery(queryConst string) string {
	switch queryConst {
	case "GET_LIKES_STATS_BY_TOOL":
		return GET_LIKES_STATS_BY_TOOL
	default:
		return ""
	}
}

func (v ToolDetailsByDomain) MakeQuery(queryConst string) string {
	switch queryConst {
	case "GET_TOOL_BY_DOMAIN_AND_SUBDOMAINS_DETAILS":
		return GET_TOOL_BY_DOMAIN_AND_SUBDOMAINS_DETAILS
	default:
		return ""
	}
}
