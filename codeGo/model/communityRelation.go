package model

import (
	"github.com/DhakshidMurali/tara/constant"
	"github.com/DhakshidMurali/tara/util"
)

type CommunityMemberEmployee struct {
	Community Community
	Member    Member
	Employee  Employee
}

type CommunityCreatedByEmployee struct {
	Community Community
	CreatedBy CreatedBy
	Employee  Employee
}

func (v CommunityMemberEmployee) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.COMMUNITY,
			"%r":  constant.MEMBER,
			"%n2": constant.EMPLOYEEEN1,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_EMPLOYEES_MEMBER_OF_COMMUNITY":
		returnData := constant.RETURNDATA_EMPLOYEE
		query := constant.RETRIEVE_DATA_NODE_WHERE
		mapData := map[string]string{
			"%n1":        "Community",
			"%r":         "Member",
			"%n2":        "Employee",
			"%condition": "elementId(n1)=$NodeId",
			"%node":      "n2",
			"%return":    returnData,
		}
		query = util.ReplaceQuery(query, mapData)
		return query

	default:
		return ""
	}
}

func (v CommunityMemberEmployee) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"CommunityName":              v.Community.Name,
			"CommunityDescription":       v.Community.Description,
			"CommunityAccessType":        v.Community.AccessType,
			"MemberAddedToCommunityDate": v.Member.AddedToCommunityDate,
			"EmployeeName":               v.Employee.Name,
			"EmployeeMailAddress":        v.Employee.MailAddress,
			"EmployeeDomain":             v.Employee.Domain,
			"EmployeePassword":           v.Employee.Password,
			"EmployeePhoneNumber":        v.Employee.PhoneNumber,
			"EmployeeEmployeeRole":       v.Employee.Role,
			"EmployeeLocation":           v.Employee.Location,
		}
	case "LIST_EMPLOYEES_MEMBER_OF_COMMUNITY":
		return map[string]any{
			"NodeId": v.Community.Key,
		}
	default:
		return map[string]any{}
	}
}

func (v CommunityCreatedByEmployee) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.COMMUNITY,
			"%r":  constant.CREATEDBY,
			"%n2": constant.EMPLOYEEEN1,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_COMMUNITY_CREATED_BY_EMPLOYEE":
		returnData := constant.RETURNDATA_COMMUNITY
		query := constant.RETRIEVE_DATA_NODE_WHERE
		mapData := map[string]string{
			"%n1":        "Community",
			"%r":         "CreatedBy",
			"%n2":        "Employee",
			"%condition": "elementId(n2)=$NodeId",
			"%node":      "n1",
			"%return":    returnData,
		}
		query = util.ReplaceQuery(query, mapData)
		return query

	default:
		return ""
	}
}

func (v CommunityCreatedByEmployee) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"CommunityName":                 v.Community.Name,
			"CommunityDescription":          v.Community.Description,
			"CommunityAccessType":           v.Community.AccessType,
			"CreatedByCommunityCreatedDate": v.CreatedBy.CommunityCreatedDate,
			"EmployeeName":                  v.Employee.Name,
			"EmployeeMailAddress":           v.Employee.MailAddress,
			"EmployeeDomain":                v.Employee.Domain,
			"EmployeePassword":              v.Employee.Password,
			"EmployeePhoneNumber":           v.Employee.PhoneNumber,
			"EmployeeEmployeeRole":          v.Employee.Role,
			"EmployeeLocation":              v.Employee.Location,
		}
	case "LIST_COMMUNITY_CREATED_BY_EMPLOYEE":
		return map[string]any{
			"NodeId": v.Employee.Key,
		}
	default:
		return map[string]any{}
	}
}
