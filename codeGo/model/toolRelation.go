package model

import (
	"github.com/DhakshidMurali/tara/constant"
	"github.com/DhakshidMurali/tara/util"
)

type ToolRequestByEmployee struct {
	Tool      Tool
	RequestBy RequestBy
	Employee  Employee
}

type ToolAccessToEmployee struct {
	Tool     Tool
	AccessTo AccessTo
	Employee Employee
}

type ToolManagedByEmployee struct {
	Tool     Tool
	Employee Employee
}

type ToolComesUnderDepartment struct {
	Tool     Tool
	Employee Employee
}

func (v ToolRequestByEmployee) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.TOOL,
			"%r":  constant.RequestBy,
			"%n2": constant.EMPLOYEEEN1,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_TOOL_REQUESTED_BY_EMPLOYEE":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_TOOL
		mapData := map[string]string{
			"%n1":        "Tool",
			"%r":         "RequestBy",
			"%n2":        "Employee",
			"%condition": "elementId(n2)=$NodeId",
			"%node":      "n1",
			"%return":    returnData,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_EMPLOYEES_REQUESTED_TO_TOOL":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_EMPLOYEE
		mapData := map[string]string{
			"%n1":        "Tool",
			"%r":         "RequestBy",
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

func (v ToolRequestByEmployee) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"ToolName":                     v.Tool.Name,
			"ToolApprovalType":             v.Tool.ApprovalType,
			"RequestByRequestedDate":       v.RequestBy.RequestedDate,
			"RequestByRequestByAccessType": v.RequestBy.AccessType,
			"RequestByLvl1Approval":        v.RequestBy.Lvl1Approval,
			"RequestByLvl2Approval":        v.RequestBy.Lvl2Approval,
			"EmployeeName":                 v.Employee.Name,
			"EmployeeMailAddress":          v.Employee.MailAddress,
			"EmployeeDomain":               v.Employee.Domain,
			"EmployeePassword":             v.Employee.Password,
			"EmployeePhoneNumber":          v.Employee.PhoneNumber,
			"EmployeeEmployeeRole":         v.Employee.Role,
			"EmployeeLocation":             v.Employee.Location,
		}
	default:
		return map[string]any{}
	}
}

func (v ToolAccessToEmployee) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.TOOL,
			"%r":  constant.AccessTo,
			"%n2": constant.EMPLOYEEEN1,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	default:
		return ""
	}
}

func (v ToolAccessToEmployee) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"ToolName":             v.Tool.Name,
			"ToolApprovalType":     v.Tool.ApprovalType,
			"RequestByAccessType":  v.AccessTo.AccessType,
			"EmployeeName":         v.Employee.Name,
			"EmployeeMailAddress":  v.Employee.MailAddress,
			"EmployeeDomain":       v.Employee.Domain,
			"EmployeePassword":     v.Employee.Password,
			"EmployeePhoneNumber":  v.Employee.PhoneNumber,
			"EmployeeEmployeeRole": v.Employee.Role,
			"EmployeeLocation":     v.Employee.Location,
		}
	default:
		return map[string]any{}
	}
}

func (v ToolManagedByEmployee) MakeQuery(typeOfQuery string) string {
	query := constant.CREATE_NODE_TO_NODE_RELATION
	mapData := map[string]string{
		"%n1": constant.TOOL,
		"%r":  constant.ManagedBy,
		"%n2": constant.EMPLOYEEEN1,
	}
	query = util.ReplaceQuery(query, mapData)
	return query
}
func (v ToolManagedByEmployee) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"ToolName":             v.Tool.Name,
			"ToolApprovalType":     v.Tool.ApprovalType,
			"EmployeeName":         v.Employee.Name,
			"EmployeeMailAddress":  v.Employee.MailAddress,
			"EmployeeDomain":       v.Employee.Domain,
			"EmployeePassword":     v.Employee.Password,
			"EmployeePhoneNumber":  v.Employee.PhoneNumber,
			"EmployeeEmployeeRole": v.Employee.Role,
			"EmployeeLocation":     v.Employee.Location,
		}
	default:
		return map[string]any{}
	}
}
func (v ToolComesUnderDepartment) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.TOOL,
			"%r":  constant.ComesUnder,
			"%n2": constant.EMPLOYEEEN1,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	default:
		return ""
	}
}

func (v ToolComesUnderDepartment) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"ToolName":             v.Tool.Name,
			"ToolApprovalType":     v.Tool.ApprovalType,
			"EmployeeName":         v.Employee.Name,
			"EmployeeMailAddress":  v.Employee.MailAddress,
			"EmployeeDomain":       v.Employee.Domain,
			"EmployeePassword":     v.Employee.Password,
			"EmployeePhoneNumber":  v.Employee.PhoneNumber,
			"EmployeeEmployeeRole": v.Employee.Role,
			"EmployeeLocation":     v.Employee.Location,
		}
	default:
		return map[string]any{}
	}
}
