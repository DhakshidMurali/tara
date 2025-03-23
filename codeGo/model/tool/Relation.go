package toolModel

type ToolPageDetails struct {
	ToolList      []ToolList
	ToolAnalytics ToolAnalytics
}
type ToolList struct {
	ToolName           string
	CommunicationCount int
	DomainName         []string
}

type ToolAnalytics struct {
	DomainToolListMap map[string][]tool
}

type tool struct {
	Name               string
	CommunicationCount int
	BookmarkCount      int
}

func (v ToolPageDetails) getToolPageDetails(selectedTools []string) ToolList {
	if len(selectedTools) == 0 {
		return constant.
	}
}
