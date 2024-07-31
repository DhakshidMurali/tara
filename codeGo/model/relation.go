package model

type PostedIn struct {
	PostedInDate string
}

type PostedBy struct {
	PostedByDate string
}

type CollaboratedWith struct {
	TaskDescription string
}

type WorksIn struct {
	Role string
}

type SkilledIn struct {
	Experience string
}

type CreatedBy struct {
	CommunityCreatedDate string
}

type Member struct {
	AddedToCommunityDate string
}

type RequestBy struct {
	RequestedDate string
	AccessType    string
	Lvl1Approval  string
	Lvl2Approval  string
}

type AccessTo struct {
	AccessType string
}

