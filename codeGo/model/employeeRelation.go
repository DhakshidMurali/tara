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
	case "LIST_EMPLOYEE_COLLABORATED_WITH":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := `{
			Name:n1.Name,
			MailAddress:n1.Description,
			Domain:n1.AccessType,
			Password:n1.Password,
			PhoneNumber:n1.PhoneNumber,
			Role:n1.Role,
			Location:n1.Location,
			Key:elementId(n1)
			}
		`
		mapData := map[string]string{
			"%n1":        "Employee",
			"%r":         "CollaboratedWith",
			"%n2":        "Employee",
			"%condition": "elementId(n1)=$NodeId",
			"%return":    returnData,
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
	case "LIST_EMPLOYEE_COLLABORATED_WITH":
		return map[string]any{
			"NodeId": v.EmployeeN1.Key,
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
	case "LIST_TOOLS_EMPLOYEE_WORKS_IN":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := `{
			Name:n1.Name,
			ApprovalType:n1.ApprovalType,
			Key:elementId(n1)
			}
		`
		mapData := map[string]string{
			"%n1":        "Employee",
			"%r":         "WorksIn",
			"%n2":        "Tool",
			"%condition": "elementId(n2)=$NodeId",
			"%return":    returnData,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_EMPLOYEES_WORKING_IN_TOOL":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := `{
			Name:n1.Name,
			MailAddress:n1.Description,
			Domain:n1.AccessType,
			Password:n1.Password,
			PhoneNumber:n1.PhoneNumber,
			Role:n1.Role,
			Location:n1.Location,
			Key:elementId(n1)
			}
		`
		mapData := map[string]string{
			"%n1":        "Employee",
			"%r":         "WorksIn",
			"%n2":        "Tool",
			"%condition": "elementId(n2)=$NodeId",
			"%return":    returnData,
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
	case "LIST_TOOLS_EMPLOYEE_WORKS_IN":
		return map[string]any{
			"NodeId": v.Employee.Key,
		}
	case "LIST_EMPLOYEES_WORKING_IN_TOOL":
		return map[string]any{
			"NodeId": v.Tool.Key,
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
	case "LIST_EMPLOYEE_SKILL_IN":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := `{
			Name:n1.Name,
			MailAddress:n1.Description,
			Domain:n1.AccessType,
			Password:n1.Password,
			PhoneNumber:n1.PhoneNumber,
			Role:n1.Role,
			Location:n1.Location,
			Key:elementId(n1)
			}
		`
		mapData := map[string]string{
			"%n1":        "Employee",
			"%r":         "SkilledIn",
			"%n2":        "Skills",
			"%condition": "elementId(n2)=$NodeId",
			"%return":    returnData,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_SKILLS_SKILLED_BY_EMPLOYEE":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := `{
			SkillName:n1.SkillName,
			Key:elementId(n2)
			}
		`
		mapData := map[string]string{
			"%n1":        "Employee",
			"%r":         "SkilledIn",
			"%n2":        "Skills",
			"%condition": "elementId(n1)=$NodeId",
			"%return":    returnData,
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
	case "LIST_EMPLOYEE_SKILL_IN":
		return map[string]any{
			"NodeId": v.Skills.Key,
		}
	case "LIST_SKILLS_SKILLED_BY_EMPLOYEE":
		return map[string]any{
			"NodeId": v.Employee.Key,
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
	case "LIST_EMPLOYEE_REPORTS_TO_EMPLOYEE":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := `{
			Name:n2.Name,
			MailAddress:n2.Description,
			Domain:n2.AccessType,
			Password:n2.Password,
			PhoneNumber:n2.PhoneNumber,
			Role:n2.Role,
			Location:n2.Location,
			Key:elementId(n2)
			}
		`
		mapData := map[string]string{
			"%n1":        "Employee",
			"%r":         "ReportTo",
			"%n2":        "Employee",
			"%condition": "elementId(n1)=$NodeId",
			"%return":    returnData,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_EMPLOYEE_REPORTEE":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := `{
			Name:n1.Name,
			MailAddress:n1.Description,
			Domain:n1.AccessType,
			Password:n1.Password,
			PhoneNumber:n1.PhoneNumber,
			Role:n1.Role,
			Location:n1.Location,
			Key:elementId(n1)
			}
		`
		mapData := map[string]string{
			"%n1":        "Employee",
			"%r":         "ReportTo",
			"%n2":        "Employee",
			"%condition": "elementId(n2)=$NodeId",
			"%return":    returnData,
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
	case "LIST_EMPLOYEE_REPORTS_TO_EMPLOYEE":
		return map[string]any{
			"NodeId": v.Employee.Key,
		}
	default:
		return map[string]any{}
	}
}
