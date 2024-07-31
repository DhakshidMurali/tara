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

func (v CommunicationPostedInCommunity) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.COMMUNICATION,
			"%r":  constant.POSTEDIN,
			"%n2": constant.COMMUNITY,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_COMMUNICATION_POSTED_IN_COMMUNITY":
		returnData := constant.RETURNDATA_COMMUNICATION
		query := constant.RETRIEVE_DATA_NODE_WHERE
		mapData := map[string]string{
			"%n1":        "Communication",
			"%r":         "PostedIn",
			"%n2":        "Community",
			"%condition": "elementId(n2)=$NodeId",
			"%node":      "n1",
			"%return":    returnData,
		}
		query = util.DoubleReplaceQuery(query, mapData)
		return query

	default:
		return ""
	}

}

func (v CommunicationPostedInCommunity) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"CommunicationType":    v.Communication.Type,
			"CommunicationContent": v.Communication.Content,
			"PostedInPostedInDate": v.PostedIn.PostedInDate,
			"CommunityName":        v.Community.Name,
			"CommunityDescription": v.Community.Description,
			"CommunityAccessType":  v.Community.AccessType,
		}
	case "LIST_COMMUNICATION_POSTED_IN_COMMUNITY":
		return map[string]any{
			"NodeId": v.Community.Key,
		}
	default:
		return map[string]any{}
	}
}

func (v CommunicationPostedByEmployee) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.COMMUNICATION,
			"%r":  constant.POSTEDBY,
			"%n2": constant.EMPLOYEEEN1,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_COMMUNICATION_POSTED_BY_EMPLOYEE":
		returnData := constant.RETURNDATA_COMMUNICATION
		query := constant.RETRIEVE_DATA_NODE_WHERE
		mapData := map[string]string{
			"%n1":        "Communication",
			"%r":         "PostedBy",
			"%n2":        "Employee",
			"%condition": "elementId(n2)=$NodeId",
			"%node":      "n1",
			"%return":    returnData,
		}
		query = util.DoubleReplaceQuery(query, mapData)
		return query
	default:
		return ""
	}
}

func (v CommunicationPostedByEmployee) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"CommunicationType":    v.Communication.Type,
			"CommunicationContent": v.Communication.Content,
			"PostedByPostedByDate": v.PostedBy.PostedByDate,
			"EmployeeName":         v.Employee.Name,
			"EmployeeMailAddress":  v.Employee.MailAddress,
			"EmployeeDomain":       v.Employee.Domain,
			"EmployeePassword":     v.Employee.Password,
			"EmployeePhoneNumber":  v.Employee.PhoneNumber,
			"EmployeeEmployeeRole": v.Employee.Role,
			"EmployeeLocation":     v.Employee.Location,
		}
	case "LIST_COMMUNICATION_POSTED_BY_EMPLOYEE":
		return map[string]any{
			"NodeId": v.Employee.Key,
		}
	default:
		return map[string]any{}
	}
}
