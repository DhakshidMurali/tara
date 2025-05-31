package toolModel

type ToolPageDetails struct {
	ToolName           string   `json:"ToolName,omitempty"`
	CommunicationCount int      `json:"CommunicationCount,omitempty"`
	DomainName         []string `json:"DomainName,omitempty"`
}

func (v ToolPageDetails) MakeQuery(queryConst string) string {
	switch queryConst {
	case "GET_LIST_TOOLS_ORDERBY_COMMUNICATION":
		return GET_TOOL_MENU
	default:
		return ""
	}
}
