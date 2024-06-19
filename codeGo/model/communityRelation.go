package model

type CommunityMemberEmployee struct{
	Community Community
	AddedToCommunityDate string
	Employee Employee
}

type CommunityCreatedByEmployee struct{
	Community Community
	timeStamp string
	Employee Employee
}

