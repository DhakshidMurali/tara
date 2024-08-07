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

	RETRIEVE_DATA_NODES_ALL = `
	MATCH (n1:%n1)
	RETURN %return
	`
	RETRIEVE_DATA_NODES = `
	MATCH (n1:%n1)-(r:%r)->(n2:%n2)
	RETURN %return
	`

	RETRIEVE_DATA_NODE_WHERE = `
	MATCH (n1:%n1)-[r:%rel]->(n2:%n2)
	WHERE %condition
	RETURN %return
	`

	EMPLOYEEEN1 = `
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

	EMPLOYEEEN2 = `
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
	TOOL = `
	Tool{
		Name:$ToolName,
		ApprovalType:$ToolApprovalType
	}
	`
	DEPARTMENT = `
	Department{
		DepartmentName:$DepartmentName
	}
	`

	COMMUNITY = `
	Community{
		Name:$CommunityName,
		Description:$CommunityDescription,
		AccessType:$CommunityAccessType
	}
	`

	COMMUNICATION = `
	Communication{
		Type:$CommunicationType,
		Content:$CommunicationContent
	}
	`

	SKILLS = `
	Skills{
		SkillName:$SkillsSkills
	}
	`

	POSTEDIN = `
	PostedIn{
		PostedInDate:$PostedInPostedInDate
	}
	`
	POSTEDBY = `
	PostedBy{
		PostedByDate:$PostedByPostedByDate
	}
	`

	COLLABORATEDWITH = `
	CollaboratedWith{
		TaskDescription:$CollaboratedWithTaskDescription
	}
	`

	WORKSIN = `
	WorksIn{
		Role:$WorksInWorksInRole
	}
	`

	SKILLEDIN = `
	SkilledIn{
		Experience:$SkilledInExperience
	}
	`

	CREATEDBY = `
	CreatedBy{
		CommunityCreatedDate:$CreatedByCommunityCreatedDate
	}
	`

	MEMBER = `
	Member{
		AddedToCommunityDate:$MemberAddedToCommunityDate
	}
	`

	REQUESTBY = `
	RequestBy{
		RequestedDate:$RequestByRequestedDate,
		AccessType:$RequestByRequestByAccessType,
		Lvl1Approval:$RequestByLvl1Approval,
		Lvl2Approval:$RequestByLvl2Approval
	}
	`

	ACCESSTO = `
	AccessTo{
		AccessType:$RequestByAccessType
	}`

	MANAGEDBY = `ManagedBy{}`

	COMESUNDER = `ComesUnder{}`

	REPORTTO = `ReportTo{}`

	RETURNDATA_EMPLOYEE = `%node{
		Name:%node.Name,
		MailAddress:%node.MailAddress,
		Domain:%node.Domain,
		Password:%node.Password,
		PhoneNumber:%node.PhoneNumber,
		Role:%node.Role,
		Location:%node.Location,
		Key:elementId(%node)
		}
	`
	RETURNDATA_TOOL = `%node{
		Name:%node.Name,
		ApprovalType:%node.ApprovalType,
		Key:elementId(%node)
	}`

	RETURNDATA_DEPARTMENT = `%node{
		DepartmentName:%node.DepartmentName,
		Key:elementId(%node)
	}`

	RETURNDATA_COMMUNICATION = `%node{
		Type:%node.Type,
		Content:%node.Content,
		Key:elementId(%node)
	}`

	RETURNDATA_COMMUNITY = `%node{
		Name:%node.Name,
		Description:%node.Description,
		AccessType:%node.AccessType,
		Key:elementId(%node)
	}`
	RETURNDATA_SKILLS = `%node{
		SkillName:%node.SkillName,
		Key:elementId(%node)
	}`
)
