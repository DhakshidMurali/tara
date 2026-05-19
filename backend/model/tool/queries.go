package toolModel

const (
	GET_TOP_TOOLS_ORDER_BY_COMMUNICATION = `
	MATCH (N1:TOOL)-[:USED_FOR]->(N2:SUBDOMAIN)
	MATCH (N2:SUBDOMAIN)-[:PART_OF]->(N3:DOMAIN)
	WITH N1, N2, N3
	MATCH (N1)
	RETURN {
	toolName:N1.toolName,
	subDomain:N2.subDomainName,
	domainName:N3.domainName,
	communicationCount:N1.communicationCount
	} AS response
	ORDER BY N1.communicationCount DESC
    LIMIT 10
	`
	GET_TOP_COMMNICATION_LISTED_TOOL = `
	MATCH (N1:TOOL)-[r:COMESUNDER]->(N2:Domain)
	WITH N2, COUNT(N2.DomainName) AS CountToolsGroupByDomain,N1
	ORDER BY CountToolsGroupByDomain DESC
	RETURN N2{
		toolName:N1.NAME,
		domainName:N2.DomainName,
		countToolsGroupByDomain:CountToolsGroupByDomain
	}
	`

	GET_LIKES_STATS_BY_TOOL = `
	MATCH (N:TOOL)
	RETURN {
	toolName:N.toolName,
	likes:N.likes
	} AS response
	ORDER BY N.likes DESC
	LIMIT 15
	`

	GET_TOOL_BY_DOMAIN_AND_SUBDOMAINS_DETAILS = `
	MATCH (N1:TOOL)-[:USED_FOR]->(N2:SUBDOMAIN)
	MATCH (N2:SUBDOMAIN)-[:PART_OF]->(N3:DOMAIN)
	WITH N1, N2, N3
	MATCH (N1)
	RETURN {
	toolName:N1.toolName,
	subDomain:N2.subDomainName,
	domainName:N3.domainName,
	likes:N1.likes
	} AS response
	 `
)
