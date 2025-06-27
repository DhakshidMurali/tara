package domainModel

const (
	GET_DOMAIN_AND_SUBDOMAINS_DETAILS = `
	MATCH (N1:SUBDOMAIN)-[:PART_OF]->(N2:DOMAIN)
	WITH N2.domainName AS domainName, COLLECT(N1.subDomainName) AS subDomains
	RETURN {
		domainName: domainName,
		subDomains: subDomains
	} as response`

	GET_TOP_DOMAINS_ORDER_BY_TOOLS = `
	MATCH (N1:TOOL)-[:USED_FOR]->(N2:SUBDOMAIN)
    MATCH (N2:SUBDOMAIN)-[:PART_OF]->(N3:DOMAIN)
    WITH N3.domainName AS domainName, COUNT(N1) AS toolsCount
    ORDER BY toolsCount DESC
    RETURN {
      domainName:domainName,
      toolsCount:toolsCount
    } as response`

	GET_DELIVERY_FORMAT_GROUP_BY_DOMAINS = `
	    MATCH (N1:TOOL)-[:USED_FOR]->(N2:SUBDOMAIN)
    MATCH (N2:SUBDOMAIN)-[:PART_OF]->(N3:DOMAIN)
    WITH N3, COUNT (N1) AS countOfTools
     ORDER BY countOfTools DESC
    LIMIT 5
    WITH COLLECT(N3.domainName) AS Top5Domains
    MATCH (N1:TOOL)-[:USED_FOR]->(N2:SUBDOMAIN)
    MATCH (N2:SUBDOMAIN)-[:PART_OF]->(N3:DOMAIN)
    WHERE N3.domainName IN Top5Domains
    WITH N1.deliveryFormat AS deliveryFormat, N3, COUNT(N1.deliveryFormat) AS groupByDeliveryFormat
    RETURN {
      deliveryFormat:deliveryFormat,
      domainName:N3.domainName,
      groupByDeliveryFormat:groupByDeliveryFormat
      } AS result`
)
