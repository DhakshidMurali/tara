package model

import (
	"github.com/DhakshidMurali/tara/constant"
	"github.com/DhakshidMurali/tara/util"
)

type CommunicationPostedInCommunity struct {
	Communication Communication `json:"communication"`
	PostedIn      PostedIn
	Community     Community
}

type CommunicationPostedByEmployee struct {
	Communication Communication
	PostedBy      PostedBy
	Employee      Employee
}

func (v CommunicationPostedInCommunity) MakeQuery() string {
	query := constant.CREATE_NODE_TO_NODE_RELATION
	mapData := map[string]string{
		"%n1": constant.Communication,
		"%r":  constant.PostedIn,
		"%n2": constant.Community,
	}
	query = util.ReplaceQuery(query, mapData)
	return query
}

func (v CommunicationPostedInCommunity) MakeParams() map[string]any {
	return map[string]any{
		"CommunicationType":    v.Communication.Type,
		"CommunicationContent": v.Communication.Content,
		"PostedInPostedInDate": v.PostedIn.PostedInDate,
		"CommunityName":        v.Community.Name,
		"CommunityDescription": v.Community.Description,
		"CommunityAccessType":  v.Community.AccessType,
	}
}

func (v CommunicationPostedByEmployee) MakeQuery() string {
	query := constant.CREATE_NODE_TO_NODE_RELATION
	mapData := map[string]string{
		"%n1": constant.Communication,
		"%r":  constant.PostedBy,
		"%n2": constant.EmployeeN1,
	}
	query = util.ReplaceQuery(query, mapData)
	return query
}

func (v CommunicationPostedByEmployee) MakeParams() map[string]any {
	return map[string]any{
		"CommunicationType":    v.Communication.Type,
		"CommunicationContent": v.Communication.Content,
		"PostedInPostedInDate": v.PostedBy.PostedByDate,
		"EmployeeName":v.Employee.Name,
		"EmployeeMailAddress":v.Employee.MailAddress,
		"EmployeeDomain":v.Employee.Domain,
		"EmployeePassword":v.Employee.Password,
		"EmployeePhoneNumber":v.Employee.PhoneNumber,
		"EmployeeEmployeeRole":v.Employee.Role,
		"EmployeeLocation":v.Employee.Location,
	}
}
