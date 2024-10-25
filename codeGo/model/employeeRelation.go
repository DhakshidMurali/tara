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
type EmployeeComesUnderDepartment struct {
	Employee   Employee
	Department Department
}
type EmployeeGroupByDepartment struct {
	EmployeeCount  int
	DepartmentName string
}

func (v EmployeeCollaboratedWithEmployee) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.EMPLOYEEEN1,
			"%r":  constant.COLLABORATEDWITH,
			"%n2": constant.EMPLOYEEEN2,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_EMPLOYEE_COLLABORATED_WITH":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_EMPLOYEE
		mapData := map[string]string{
			"%n1":        "Employee",
			"%rel":       "CollaboratedWith",
			"%n2":        "Employee",
			"%condition": "elementId(n1)=$NodeId",
			"%node":      "n2",
			"%return":    returnData,
		}
		query = util.DoubleReplaceQuery(query, mapData)
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
			"%n1": constant.EMPLOYEEEN1,
			"%r":  constant.WORKSIN,
			"%n2": constant.TOOL,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_TOOLS_EMPLOYEE_WORKS_IN":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_TOOL
		mapData := map[string]string{
			"%n1":        "Employee",
			"%rel":       "WorksIn",
			"%n2":        "Tool",
			"%condition": "elementId(n1)=$NodeId",
			"%node":      "n2",
			"%return":    returnData,
		}
		query = util.DoubleReplaceQuery(query, mapData)
		return query
	case "LIST_EMPLOYEES_WORKS_IN_TOOL":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_EMPLOYEE
		mapData := map[string]string{
			"%n1":        "Employee",
			"%rel":       "WorksIn",
			"%n2":        "Tool",
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
	case "LIST_EMPLOYEES_WORKS_IN_TOOL":
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
			"%n1": constant.EMPLOYEEEN1,
			"%r":  constant.SKILLEDIN,
			"%n2": constant.SKILLS,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_EMPLOYEE_SKILLED_IN":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_EMPLOYEE
		mapData := map[string]string{
			"%n1":        "Employee",
			"%rel":       "SkilledIn",
			"%n2":        "Skills",
			"%condition": "elementId(n2)=$NodeId",
			"%node":      "n1",
			"%return":    returnData,
		}
		query = util.DoubleReplaceQuery(query, mapData)
		return query
	case "LIST_SKILLS_SKILLED_BY_EMPLOYEE":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_SKILLS
		mapData := map[string]string{
			"%n1":        "Employee",
			"%rel":       "SkilledIn",
			"%n2":        "Skills",
			"%condition": "elementId(n1)=$NodeId",
			"%node":      "n2",
			"%return":    returnData,
		}
		query = util.DoubleReplaceQuery(query, mapData)
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
	case "LIST_EMPLOYEE_SKILLED_IN":
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
			"%n1": constant.EMPLOYEEEN1,
			"%r":  constant.REPORTTO,
			"%n2": constant.EMPLOYEEEN2,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_MANAGER_HIERARCHY_OF_EMPLOYEE":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_EMPLOYEE
		mapData := map[string]string{
			"%n1":        "Employee",
			"%rel":       "ReportTo",
			"%n2":        "Employee",
			"%condition": "elementId(n1)=$NodeId",
			"%node":      "n2",
			"%return":    returnData,
		}
		query = util.DoubleReplaceQuery(query, mapData)
		return query
	case "LIST_EMPLOYEE_REPORTEE_OF_MANAGER":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_EMPLOYEE
		mapData := map[string]string{
			"%n1":        "Employee",
			"%rel":       "ReportTo",
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
	case "LIST_MANAGER_HIERARCHY_OF_EMPLOYEE":
		return map[string]any{
			"NodeId": v.Employee.Key,
		}
	case "LIST_EMPLOYEE_REPORTEE_OF_MANAGER":
		return map[string]any{
			"NodeId": v.Manager.Key,
		}
	default:
		return map[string]any{}
	}
}

func (v EmployeeComesUnderDepartment) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.EMPLOYEEEN1,
			"%r":  constant.COMESUNDER,
			"%n2": constant.DEPARTMENT,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_DEPARTMENT_OF_EMPLOYEE":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_EMPLOYEE
		mapData := map[string]string{
			"%n1":        "Employee",
			"%rel":       "ComesUnder",
			"%n2":        "Department",
			"%condition": "elementId(n1)=$NodeId",
			"%node":      "n2",
			"%return":    returnData,
		}
		query = util.DoubleReplaceQuery(query, mapData)
		return query
	case "LIST_EMPLOYEE_COMES_UNDER_DEPARTMENT":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_EMPLOYEE
		mapData := map[string]string{
			"%n1":        "Employee",
			"%rel":       "ComesUnder",
			"%n2":        "Department",
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

func (v EmployeeComesUnderDepartment) MakeParams(typeOfQuery string) map[string]any {
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
			"DepartmentName":       v.Department.DepartmentName,
		}
	case "LIST_DEPARTMENT_OF_EMPLOYEE":
		return map[string]any{
			"NodeId": v.Employee.Key,
		}
	case "LIST_EMPLOYEE_COMES_UNDER_DEPARTMENT":
		return map[string]any{
			"NodeId": v.Department.Key,
		}
	default:
		return map[string]any{}
	}
}
