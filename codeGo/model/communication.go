package model

type CommunicationPostedInCommunity struct {
	Communication Communication
	PostedIn      string
	Community     Community
}

type CommunicationPostedByEmployee struct {
	Communication Communication
	PostedBy      string
	Community     Community
}
