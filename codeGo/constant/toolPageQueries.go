package constant

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
	MATCH (n1:TOOL)-[r1:COMESUNDER]->(n2:Domain)
	WITH n1, n1.COMMUNICATIONCOUNT AS totalCommunicationCount,n2
	ORDER BY totalCommunicationCount DESC
	LIMIT 25
	RETURN n2.DomainName, n1
	`
)

// Match (N1:TOOL)-[R1:COMESUNDER]->(N2:Domain)
// with N1,N2
// order by N1.COMMUNICATIONCOUNT desc
// return N2.DomainName as DOMAINNAME, COLLECT(N1) as Tools

//
