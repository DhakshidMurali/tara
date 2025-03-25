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
	MATCH (n1:%n1)-[r:%r]->(n2:%n2)
	RETURN %return
	`

	RETRIEVE_DATA_NODE_WHERE = `
	MATCH (n1:%n1)-[r:%rel]->(n2:%n2)
	WHERE %condition
	RETURN %return
	`

	USEREN1 = `
	USER{
		Name:$UserName,
        MailAddress:$UserMailAddress,
        Domain:$UserDomain,
        Password:$UserPassword,
        PhoneNumber:$UserPhoneNumber,
        Role:$UserUserRole,
        Location:$UserLocation
	}
	`

	USEREN2 = `
	USER{
		Name:$UserNameN2,
        MailAddress:$UserMailAddressN2,
        Domain:$UserDomainN2,
        Password:$UserPasswordN2,
        PhoneNumber:$UserPhoneNumberN2,
        Role:$UserUserRoleN2,
        Location:$UserLocationN2
	}
	`
	TOOL = `
	TOOL{
		Name:$ToolName,
		DeliveryFormat:$ToolDeliveryFormat
	}
	`
	DOMAIN = `
	DOMAIN{
		DomainName:$DomainName
	}
	`

	COMMUNITY = `
	COMMUNITY{
		Name:$CommunityName,
		Description:$CommunityDescription,
		AccessType:$CommunityAccessType
	}
	`

	COMMUNICATION = `
	COMMUNICATION{
		Link:$CommunicationLink,
		Content:$CommunicationContent
	}
	`

	SKILLS = `
	SKILLS{
		SkillName:$SkillsSkills
	}
	`

	POSTEDIN = `
	POSTEDIN{
		PostedInDate:$PostedInPostedInDate
	}
	`
	POSTEDBY = `
	POSTEDBY{
		PostedByDate:$PostedByPostedByDate
	}
	`

	COLLABORATEDWITH = `
	COLLABORATEDWITH{
		TaskDescription:$CollaboratedWithTaskDescription
	}
	`

	WORKSIN = `
	WORKSIN{
		Role:$WorksInWorksInRole
	}
	`

	SKILLEDIN = `
	SKILLEDIN{
		Experience:$SkilledInExperience
	}
	`

	CREATEDBY = `
	CREATEDBY{
		CommunityCreatedDate:$CreatedByCommunityCreatedDate
	}
	`

	MEMBER = `
	MEMBER{
		AddedToCommunityDate:$MemberAddedToCommunityDate
	}
	`

	REQUESTBY = `
	REQUESTBY{
		RequestedDate:$RequestByRequestedDate,
		AccessType:$RequestByRequestByAccessType,
		Lvl1Approval:$RequestByLvl1Approval,
		Lvl2Approval:$RequestByLvl2Approval
	}
	`

	ACCESSTO = `
	ACCESSTO{
		AccessType:$RequestByAccessType
	}`

	MANAGEDBY = `MANAGEDBY{}`

	COMESUNDER = `COMESUNDER{}`

	REPORTTO = `REPORTTO{}`

	RETURNDATA_USER = `%node{
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

	RETURNDATA_DOMAIN = `%node{
		DomainName:%node.DomainName,
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
