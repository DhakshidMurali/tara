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
			"%n1": constant.Community,
			"%r":  constant.Member,
			"%n2": constant.EmployeeN1,
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
	default:
		return map[string]any{}
	}
}

func (v CommunityCreatedByEmployee) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.Community,
			"%r":  constant.CreatedBy,
			"%n2": constant.EmployeeN1,
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
	default:
		return map[string]any{}
	}
}
