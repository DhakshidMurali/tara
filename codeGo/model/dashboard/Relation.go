package dashboardModel

import "github.com/DhakshidMurali/tara/constant"

type ToolGroupByDomain struct {
	DomainName              string
	CountToolsGroupByDomain int
}

func (v ToolGroupByDomain) MakeQuery() string {
	query := constant.GET_LIST_TOOLS_GROUPBY_DOMAIN
	return query
}
