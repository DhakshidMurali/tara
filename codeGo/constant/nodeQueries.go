package constant

const (
	CREATE_NODE_TO_NODE_RELATION = `
	MERGE (n1:%n1)
	MERGE (n2:%n2)
	MERGE (n1)-[r:%r]->(n2)
	ON CREATE SET r.createdAt = timestamp()
	ON MATCH SET r.lastCheckedAt = timestamp()
	RETURN CASE
		WHEN r.createdAt=timestamp() THEN TRUE
		ELSE FALSE
	END AS isCreated
	`

	UPDATE_NODE = `
	MATCH (n1:%n)
	WHERE elementId(n1)=$NodeId
	SET n1+=$v
	RETURN n1
	`

	RETRIEVE_DATA_NODES_ALL=`
	MATCH (n1:%n1)
	RETURN %return
	`
	RETRIEVE_DATA_NODES=`
	MATCH (n1:%n1)-(r:%r)->(n2:%n2)
	RETURN %return
	`

	RETRIEVE_DATA_NODE_WHERE=`
	MATCH (n1:%n1)-[r:%r]->(n2:%n2)
	WHERE %condition
	RETURN %return
	`

	EmployeeN1 = `
	Employee{
		Name:$EmployeeName,
        MailAddress:$EmployeeMailAddress,
        Domain:$EmployeeDomain,
        Password:$EmployeePassword,
        PhoneNumber:$EmployeePhoneNumber,
        Role:$EmployeeEmployeeRole,
        Location:$EmployeeLocation
	}
	`

	EmployeeN2 = `
	Employee{
		Name:$EmployeeNameN2,
        MailAddress:$EmployeeMailAddressN2,
        Domain:$EmployeeDomainN2,
        Password:$EmployeePasswordN2,
        PhoneNumber:$EmployeePhoneNumberN2,
        Role:$EmployeeEmployeeRoleN2,
        Location:$EmployeeLocationN2
	}
	`
	Tool = `
	Tool{
		Name:$ToolName,
		ApprovalType:$ToolApprovalType
	}
	`
	Department = `
	Department{
		DepartmentName:$DepartmentName
	}
	`

	Community = `
	Community{
		Name:$CommunityName,
		Description:$CommunityDescription,
		AccessType:$CommunityAccessType
	}
	`

	Communication = `
	Communication{
		Type:$CommunicationType,
		Content:$CommunicationContent
	}
	`

	Skills = `
	Skills{
		Skills:$SkillsSkills
	}
	`

	PostedIn = `
	PostedIn{
		PostedInDate:$PostedInPostedInDate
	}
	`
	PostedBy = `
	PostedBy{
		PostedByDate:$PostedByPostedByDate
	}
	`

	CollaboratedWith = `'
	CollaboratedWith{
		TaskDescription:$CollaboratedWithTaskDescription
	}
	`

	WorksIn = `
	WorksIn{
		Role:$WorksInWorksInRole
	}
	`

	SkilledIn = `
	SkilledIn{
		Experience:$SkilledInExperience
	}
	`

	CreatedBy = `
	CreatedBy{
		CommunityCreatedDate:$CreatedByCommunityCreatedDate
	}
	`

	Member = `
	Member{
		AddedToCommunityDate:$MemberAddedToCommunityDate
	}
	`

	RequestBy = `
	RequestBy{
		RequestedDate:$RequestByRequestedDate,
		AccessType:$RequestByRequestByAccessType,
		Lvl1Approval:$RequestByLvl1Approval,
		Lvl2Approval:$RequestByLvl2Approval
	}
	`

	AccessTo = `
	AccessTo{
		AccessType:$RequestByAccessType
	}`

	ManagedBy = `ManagedBy{}`

	ComesUnder = `ComesUnder{}`

	ReportTo = `ReportTo{}`
)
