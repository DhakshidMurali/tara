package domainModel

const (
	GET_LIST_TOOLS_GROUPBY_DOMAIN = `
	MATCH (N1:TOOL)-[r:COMESUNDER]->(N2:DOMAIN)
	WITH N2, COUNT(N2.domainName) AS CountToolsGroupByDomain
	ORDER BY CountToolsGroupByDomain DESC
	RETURN N2{
		DomainName:N2.DomainName,
		CountToolsGroupByDomain:CountToolsGroupByDomain
	}
	`

	GET_LIST_TOOLS_GROUPBY_DELIVERYFORMAT_FOR_TOP4DOMAINS = `
	MATCH (N1:TOOL)-[r:COMESUNDER]->(N2:Domain)
	WITH N2,COUNT (r) AS NoOfToolsUnderDomain
	ORDER BY NoOfToolsUnderDomain DESC
	LIMIT 5
	WITH COLLECT(N2.DomainName) AS Top4Domains
	match (N1:TOOL)-[r:COMESUNDER]->(N2:Domain)
	where N2.DomainName in Top4Domains
	with N1.DELIVERYFORMAT as DeliveryFormat,N2, COUNT(N1.DELIVERYFORMAT) as NoOfToolsInDomain
	return N2{
		DeliveryFormat:DeliveryFormat,
		DomainName:N2.DomainName,
		DeliveryFormatCount:NoOfToolsInDomain
	}`

	GET_LIST_TOOLS_ORDERBY_COMMUNICATION = `
	MATCH (N1:TOOL)
	RETURN N1
	ORDER BY N1.COMMUNICATIONCOUNT desc
	LIMIT 25
	`
)
