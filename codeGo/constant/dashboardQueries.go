package constant

const (
	GET_LIST_TOOLS_GROUPBY_DOMAIN = `
	MATCH (N1:TOOL)-[r:COMESUNDER]->(N2:DOMAIN)
	WITH N2, COUNT(N2.DomainName) AS CountToolsGroupByDomain
	ORDER BY CountToolsGroupByDomain DESC
	RETURN N2{
		DomainName:N2.DomainName,
		CountToolsGroupByDomain:CountToolsGroupByDomain
	}
	`
)
