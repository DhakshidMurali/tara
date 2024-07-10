package model

import (
	"github.com/DhakshidMurali/tara/constant"
	"github.com/DhakshidMurali/tara/util"
)

type Employee struct {
	Name        string
	MailAddress string
	Domain      string
	Password    string
	PhoneNumber string
	Role        string
	Location    string
	Key         string
}

type Tool struct {
	Name         string
	ApprovalType string
	Key          string
}

type Department struct {
	DepartmentName string
	Key            string
}

type Community struct {
	Name        string
	Description string
	AccessType  string
	Key         string
}

type Communication struct {
	Type    string
	Content string
	Key     string
}

type Skills struct {
	SkillName string
	Key       string
}

type Node struct {
	NodeName string
}

type UpdateNode struct {
	Key  string
	Node interface{}
}

func (v Node) MakeQuery() string {
	query := constant.UPDATE_NODE
	mapData := map[string]string{
		"%n": v.NodeName,
	}
	query = util.ReplaceQuery(query, mapData)
	return query
}

func (v Employee) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "LIST_EMPLOYEE":
		returnData := `n1{
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
		query := constant.RETRIEVE_DATA_NODES_ALL
		mapData := map[string]string{
			"%n1":     "Department",
			"%return": returnData,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	default:
		return ""
	}
}
func (v Employee) MakeParams(key string, typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "UPDATE":
		params := make(map[string]string)
		util.AddToParams(v.Name, "Name", params)
		util.AddToParams(v.MailAddress, "MailAddress", params)
		util.AddToParams(v.Domain, "Domain", params)
		util.AddToParams(v.Password, "Password", params)
		util.AddToParams(v.PhoneNumber, "PhoneNumber", params)
		util.AddToParams(v.Role, "Role", params)
		util.AddToParams(v.Location, "Location", params)
		return map[string]any{
			"NodeId": key,
			"v":      params,
		}
	default:
		return map[string]any{}
	}
}
func (v Tool) MakeParams(key string, typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "UPDATE":
		params := make(map[string]string)
		util.AddToParams(v.Name, "Name", params)
		util.AddToParams(v.ApprovalType, "ApprovalType", params)
		return map[string]any{
			"NodeId": key,
			"v":      params,
		}
	default:
		return map[string]any{}
	}
}
func (v Department) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "LIST_DEPARTMENT":
		returnData := `n1{
			DepartmentName:n1.DepartmentName,
			ManagedBy:n2.Name,
			Key:elementId(n1)
			}
		`
		query := constant.RETRIEVE_DATA_NODES
		mapData := map[string]string{
			"%n1":     "Department",
			"%r":      "ManagedBy",
			"%n2":     "Employee",
			"%return": returnData,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	default:
		return ""
	}
}
func (v Department) MakeParams(key string, typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "UPDATE":
		params := make(map[string]string)
		util.AddToParams(v.DepartmentName, "DepartmentName", params)
		return map[string]any{
			"NodeId": key,
			"v":      params,
		}
	default:
		return map[string]any{}
	}
}

func (v Community) MakeParams(key string, typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "UPDATE":
		params := make(map[string]string)
		util.AddToParams(v.Name, "Name", params)
		util.AddToParams(v.Description, "Description", params)
		util.AddToParams(v.AccessType, "AccessType", params)
		return map[string]any{
			"NodeId": key,
			"v":      params,
		}
	default:
		return map[string]any{}
	}
}
func (v Community) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "LIST_COMMUNITY":
		returnData := `n1{
			Name:n1.Name,
			Description:n1.Description,
			AccessType:n1.AccessType,
			Key:elementId(n1)
			}
		`
		query := constant.RETRIEVE_DATA_NODES_ALL
		mapData := map[string]string{
			"%n1":     "Community",
			"%return": returnData,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	default:
		return ""
	}
}
func (v Communication) MakeParams(key string, typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "UPDATE":
		params := make(map[string]string)
		util.AddToParams(v.Type, "Type", params)
		util.AddToParams(v.Content, "Content", params)
		return map[string]any{
			"NodeId": key,
			"v":      params,
		}
	case "LIST_COMMUNITY":
		return map[string]any{
			"nodeId": v.Key,
		}
	default:
		return map[string]any{}
	}
}
func (v Skills) MakeParams(key string, typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "UPDATE":
		params := make(map[string]string)
		util.AddToParams(v.SkillName, "SkillName", params)
		return map[string]any{
			"NodeId": key,
			"v":      params,
		}
	default:
		return map[string]any{}
	}
}
