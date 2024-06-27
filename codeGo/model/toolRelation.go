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

func (v ToolRequestByEmployee) MakeQuery() string {
	query := constant.CREATE_NODE_TO_NODE_RELATION
	mapData := map[string]string{
		"%n1": constant.Tool,
		"%r":  constant.RequestBy,
		"%n2": constant.EmployeeN1,
	}
	query = util.ReplaceQuery(query, mapData)
	return query
}

func (v ToolRequestByEmployee) MakeParams() map[string]any {
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
}

func (v ToolAccessToEmployee) MakeQuery() string {
	query := constant.CREATE_NODE_TO_NODE_RELATION
	mapData := map[string]string{
		"%n1": constant.Tool,
		"%r":  constant.AccessTo,
		"%n2": constant.EmployeeN1,
	}
	query = util.ReplaceQuery(query, mapData)
	return query
}

func (v ToolAccessToEmployee) MakeParams() map[string]any {
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
}

func (v ToolManagedByEmployee) MakeQuery() string {
	query := constant.CREATE_NODE_TO_NODE_RELATION
	mapData := map[string]string{
		"%n1": constant.Tool,
		"%r":  constant.ManagedBy,
		"%n2": constant.EmployeeN1,
	}
	query = util.ReplaceQuery(query, mapData)
	return query
}
func (v ToolManagedByEmployee) MakeParams() map[string]any {
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
}
func (v ToolComesUnderDepartment) MakeQuery() string {
	query := constant.CREATE_NODE_TO_NODE_RELATION
	mapData := map[string]string{
		"%n1": constant.Tool,
		"%r":  constant.ComesUnder,
		"%n2": constant.EmployeeN1,
	}
	query = util.ReplaceQuery(query, mapData)
	return query
}

func (v ToolComesUnderDepartment) MakeParams() map[string]any {
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
}
