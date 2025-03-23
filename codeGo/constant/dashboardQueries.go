package constant

const (
	GET_LIST_TOOLS_GROUPBY_DOMAIN = `
	MATCH (N1:Tool)-[r:ComesUnder]->(N2:Domain)
	WITH N2, COUNT(N2.DomainName) AS CountToolsGroupByDomain
	ORDER BY CountToolsGroupByDomain DESC
	RETURN N2{
		DomainName:N2.DomainName,
		CountToolsGroupByDomain:CountToolsGroupByDomain
	}
	`

	GET_LIST_TOOLS_GROUPBY_DELIVERYFORMAT_FOR_TOP4DOMAINS = `
	MATCH (N1:Tool)-[r:ComesUnder]->(N2:Domain)
	WITH N2,COUNT (r) AS NoOfToolsUnderDomain
	ORDER BY NoOfToolsUnderDomain DESC
	LIMIT 4
	WITH COLLECT(N2.DomainName) AS Top4Domains
	match (N1:Tool)-[r:ComesUnder]->(N2:Domain)
	where N2.DomainName in Top4Domains
	with N1.DeliveryFormat as DeliveryFormat,N2, COUNT(N1.DeliveryFormat) as NoOfToolsInDomain
	order by N2.DomainName
	return N2{
		DeliveryFormat:DeliveryFormat,
		DomainName:N2.DomainName,
		DeliveryFormatCount:NoOfToolsInDomain
	}`
)
