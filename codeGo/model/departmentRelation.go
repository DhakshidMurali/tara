package model

import (
	"github.com/DhakshidMurali/tara/constant"
	"github.com/DhakshidMurali/tara/util"
)

type DepartmentManagedByEmployee struct {
	Department Department
	Employee   Employee
}

func (v DepartmentManagedByEmployee) MakeQuery() string {
	query := constant.CREATE_NODE_TO_NODE_RELATION
	mapData := map[string]string{
		"%n1": constant.Department,
		"%r":  constant.ManagedBy,
		"%n2": constant.EmployeeN1,
	}
	query = util.ReplaceQuery(query, mapData)
	return query
}

func (v DepartmentManagedByEmployee) MakeParams() map[string]any {
	return map[string]any{
		"DepartmentName":       v.Department.DepartmentName,
		"EmployeeName":         v.Employee.Name,
		"EmployeeMailAddress":  v.Employee.MailAddress,
		"EmployeeDomain":       v.Employee.Domain,
		"EmployeePassword":     v.Employee.Password,
		"EmployeePhoneNumber":  v.Employee.PhoneNumber,
		"EmployeeEmployeeRole": v.Employee.Role,
		"EmployeeLocation":     v.Employee.Location,
	}
}
