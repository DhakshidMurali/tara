package model

import (
	"github.com/DhakshidMurali/tara/constant"
	"github.com/DhakshidMurali/tara/util"
)

type EmployeeCollaboratedWithEmployee struct {
	EmployeeN1       Employee
	CollaboratedWith CollaboratedWith
	EmployeeN2       Employee
}

type EmployeeWorksInTools struct {
	Employee Employee
	WorksIn  WorksIn
	Tool     Tool
}

type EmployeeSkilledInSkills struct {
	Employee  Employee
	SkilledIn SkilledIn
	Skills    Skills
}

type EmployeeReportToEmployee struct {
	Employee Employee
	Manager  Employee
}

func (v EmployeeCollaboratedWithEmployee) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.EmployeeN1,
			"%r":  constant.CollaboratedWith,
			"%n2": constant.EmployeeN2,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	default:
		return ""
	}
}

func (v EmployeeCollaboratedWithEmployee) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"EmployeeName":                    v.EmployeeN1.Name,
			"EmployeeMailAddress":             v.EmployeeN1.MailAddress,
			"EmployeeDomain":                  v.EmployeeN1.Domain,
			"EmployeePassword":                v.EmployeeN1.Password,
			"EmployeePhoneNumber":             v.EmployeeN1.PhoneNumber,
			"EmployeeEmployeeRole":            v.EmployeeN1.Role,
			"EmployeeLocation":                v.EmployeeN1.Location,
			"CollaboratedWithTaskDescription": v.CollaboratedWith.TaskDescription,
			"EmployeeNameN2":                  v.EmployeeN2.Name,
			"EmployeeMailAddressN2":           v.EmployeeN2.MailAddress,
			"EmployeeDomainN2":                v.EmployeeN2.Domain,
			"EmployeePasswordN2":              v.EmployeeN2.Password,
			"EmployeePhoneNumberN2":           v.EmployeeN2.PhoneNumber,
			"EmployeeEmployeeRoleN2":          v.EmployeeN2.Role,
			"EmployeeLocationN2":              v.EmployeeN2.Location,
		}
	default:
		return map[string]any{}
	}
}

func (v EmployeeWorksInTools) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.EmployeeN1,
			"%r":  constant.WorksIn,
			"%n2": constant.Tool,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	default:
		return ""
	}
}

func (v EmployeeWorksInTools) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"EmployeeName":         v.Employee.Name,
			"EmployeeMailAddress":  v.Employee.MailAddress,
			"EmployeeDomain":       v.Employee.Domain,
			"EmployeePassword":     v.Employee.Password,
			"EmployeePhoneNumber":  v.Employee.PhoneNumber,
			"EmployeeEmployeeRole": v.Employee.Role,
			"EmployeeLocation":     v.Employee.Location,
			"WorksInWorksInRole":   v.WorksIn.Role,
			"ToolName":             v.Tool.Name,
			"ToolApprovalType":     v.Tool.ApprovalType,
		}
	default:
		return map[string]any{}
	}
}

func (v EmployeeSkilledInSkills) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.EmployeeN1,
			"%r":  constant.SkilledIn,
			"%n2": constant.Skills,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	default:
		return ""
	}
}

func (v EmployeeSkilledInSkills) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"EmployeeName":         v.Employee.Name,
			"EmployeeMailAddress":  v.Employee.MailAddress,
			"EmployeeDomain":       v.Employee.Domain,
			"EmployeePassword":     v.Employee.Password,
			"EmployeePhoneNumber":  v.Employee.PhoneNumber,
			"EmployeeEmployeeRole": v.Employee.Role,
			"EmployeeLocation":     v.Employee.Location,
			"SkilledInExperience":  v.SkilledIn.Experience,
			"SkillsSkills":         v.Skills.SkillName,
		}
	default:
		return map[string]any{}
	}
}

func (v EmployeeReportToEmployee) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.EmployeeN1,
			"%r":  constant.ReportTo,
			"%n2": constant.EmployeeN2,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	default:
		return ""
	}
}

func (v EmployeeReportToEmployee) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"EmployeeName":           v.Employee.Name,
			"EmployeeMailAddress":    v.Employee.MailAddress,
			"EmployeeDomain":         v.Employee.Domain,
			"EmployeePassword":       v.Employee.Password,
			"EmployeePhoneNumber":    v.Employee.PhoneNumber,
			"EmployeeEmployeeRole":   v.Employee.Role,
			"EmployeeLocation":       v.Employee.Location,
			"EmployeeNameN2":         v.Manager.Name,
			"EmployeeMailAddressN2":  v.Manager.MailAddress,
			"EmployeeDomainN2":       v.Manager.Domain,
			"EmployeePasswordN2":     v.Manager.Password,
			"EmployeePhoneNumberN2":  v.Manager.PhoneNumber,
			"EmployeeEmployeeRoleN2": v.Manager.Role,
			"EmployeeLocationN2":     v.Manager.Location,
		}
	default:
		return map[string]any{}
	}
}
