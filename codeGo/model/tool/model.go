package toolModel

import "github.com/DhakshidMurali/tara/constant"

type ToolPageDetails struct {
	ToolName           string   `json:"ToolName,omitempty"`
	CommunicationCount int      `json:"CommunicationCount,omitempty"`
	DomainName         []string `json:"DomainName,omitempty"`
}

func (v ToolPageDetails) MakeQuery(queryConst string) string {
	switch queryConst {
	case "GET_LIST_TOOLS_ORDERBY_COMMUNICATION":
		return constant.GET_LIST_TOOLS_ORDERBY_COMMUNICATION
	default:
		return ""
	}
}
