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
	Tool       Tool
	Department Department
}

func (v ToolRequestByEmployee) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.TOOL,
			"%r":  constant.REQUESTBY,
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
	case "LIST_TOOL_REQUESTED_BY_EMPLOYEE":
		return map[string]any{
			"NodeId": v.Employee.Key,
		}
	case "LIST_EMPLOYEES_REQUESTED_TO_TOOL":
		return map[string]any{
			"NodeId": v.Tool.Key,
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
			"%r":  constant.ACCESSTO,
			"%n2": constant.EMPLOYEEEN1,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_TOOLS_ACCESS_BY_EMPLOYEE":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_TOOL
		mapData := map[string]string{
			"%n1":        "Tool",
			"%r":         "AccessTo",
			"%n2":        "Employee",
			"%condition": "elementId(n2)=$NodeId",
			"%node":      "n1",
			"%return":    returnData,
		}
		query = util.DoubleReplaceQuery(query, mapData)
		return query
	case "LIST_EMPLOYEES_ACCESS_TO_TOOL":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_EMPLOYEE
		mapData := map[string]string{
			"%n1":        "Tool",
			"%r":         "AccessTo",
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
	case "LIST_TOOLS_ACCESS_BY_EMPLOYEE":
		return map[string]any{
			"NodeId": v.Employee.Key,
		}
	case "LIST_EMPLOYEES_ACCESS_TO_TOOL":
		return map[string]any{
			"NodeId": v.Tool.Key,
		}
	default:
		return map[string]any{}
	}
}

func (v ToolManagedByEmployee) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.TOOL,
			"%r":  constant.MANAGEDBY,
			"%n2": constant.EMPLOYEEEN1,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_TOOLS_MANAGED_BY_EMPLOYEE":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_TOOL
		mapData := map[string]string{
			"%n1":        "Tool",
			"%r":         "AccessTo",
			"%n2":        "Employee",
			"%condition": "elementId(n2)=$NodeId",
			"%node":      "n1",
			"%return":    returnData,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_EMPLOYEES_MANAGES_TOOL":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_EMPLOYEE
		mapData := map[string]string{
			"%n1":        "Tool",
			"%r":         "AccessTo",
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
	case "LIST_TOOL_REQUESTED_BY_EMPLOYEE":
		return map[string]any{
			"NodeId": v.Employee.Key,
		}
	case "LIST_EMPLOYEES_REQUESTED_TO_TOOL":
		return map[string]any{
			"NodeId": v.Tool.Key,
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
			"%r":  constant.COMESUNDER,
			"%n2": constant.DEPARTMENT,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_TOOLS_COMES_UNDER_DEPARTMENT":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_TOOL
		mapData := map[string]string{
			"%n1":        "Tool",
			"%r":         "ComesUnder",
			"%n2":        "Department",
			"%condition": "elementId(n2)=$NodeId",
			"%node":      "n1",
			"%return":    returnData,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_DEPARTMENT_MAINTAIN_TOOL":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_DEPARTMENT
		mapData := map[string]string{
			"%n1":        "Tool",
			"%r":         "ComesUnder",
			"%n2":        "Department",
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

func (v ToolComesUnderDepartment) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"ToolName":         v.Tool.Name,
			"ToolApprovalType": v.Tool.ApprovalType,
			"DepartmentName":   v.Department.DepartmentName,
		}
	case "LIST_TOOLS_COMES_UNDER_DEPARTMENT":
		return map[string]any{
			"NodeId": v.Department.Key,
		}
	case "LIST_DEPARTMENT_MAINTAIN_TOOL":
		return map[string]any{
			"NodeId": v.Tool.Key,
		}
	default:
		return map[string]any{}
	}
}
