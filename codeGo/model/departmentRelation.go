package model

import (
	"github.com/DhakshidMurali/tara/constant"
	"github.com/DhakshidMurali/tara/util"
)

type DepartmentManagedByEmployee struct {
	Department Department
	Employee   Employee
}

func (v DepartmentManagedByEmployee) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.DEPARTMENT,
			"%r":  constant.MANAGEDBY,
			"%n2": constant.EMPLOYEEEN1,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_EMPLOYEE_MANAGES_DEPARTMENT":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_DEPARTMENT
		mapData := map[string]string{
			"%n1":        "Department",
			"%rel":         "ManagedBy",
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

func (v DepartmentManagedByEmployee) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
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
	case "LIST_EMPLOYEE_MANAGES_DEPARTMENT":
		return map[string]any{
			"NodeId": v.Employee.Key,
		}
	default:
		return map[string]any{}
	}
}
