package toolModel

const (
	GET_TOOL_MENU = `
	MATCH (n1:TOOL)-[r1:COMESUNDER]->(n2:Domain)
	WITH n1, n1.COMMUNICATIONCOUNT AS totalCommunicationCount,n2
	ORDER BY totalCommunicationCount DESC
	SKIP $page
	LIMIT 15
	RETURN n2.DomainName, n1
	`
	GET_TOP_COMMNICATION_LISTED_TOOL = `
	MATCH (N1:TOOL)-[r:COMESUNDER]->(N2:Domain)
	WITH N2, COUNT(N2.DomainName) AS CountToolsGroupByDomain,N1
	ORDER BY CountToolsGroupByDomain DESC
	RETURN N2{
		ToolName:N1.NAME,
		DomainName:N2.DomainName,
		CountToolsGroupByDomain:CountToolsGroupByDomain
	}
	`
)

// Match (N1:TOOL)-[R1:COMESUNDER]->(N2:Domain)
// with N1,N2
// order by N1.COMMUNICATIONCOUNT desc
// return N2.DomainName as DOMAINNAME, COLLECT(N1) as Tools

//
