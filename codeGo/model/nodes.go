package model

import (
	"fmt"

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

func (v Employee) MakeParams(key string) map[string]any {
	fmt.Println("Name: ")
	fmt.Print(v.Name)
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
}
func (v Tool) MakeParams(key string) map[string]any {
	params := make(map[string]string)
	util.AddToParams(v.Name, "Name", params)
	util.AddToParams(v.ApprovalType, "ApprovalType", params)
	return map[string]any{
		"NodeId": key,
		"v":      params,
	}
}
func (v Department) MakeParams(key string) map[string]any {
	params := make(map[string]string)
	util.AddToParams(v.DepartmentName, "DepartmentName", params)
	return map[string]any{
		"NodeId": key,
		"v":      params,
	}
}
func (v Community) MakeParams(key string) map[string]any {
	params := make(map[string]string)
	util.AddToParams(v.Name, "Name", params)
	util.AddToParams(v.Description, "Description", params)
	util.AddToParams(v.AccessType, "AccessType", params)
	return map[string]any{
		"NodeId": key,
		"v":      params,
	}
}
func (v Communication) MakeParams(key string) map[string]any {
	params := make(map[string]string)
	util.AddToParams(v.Type, "Type", params)
	util.AddToParams(v.Content, "Content", params)
	return map[string]any{
		"NodeId": key,
		"v":      params,
	}
}
func (v Skills) MakeParams(key string) map[string]any {
	params := make(map[string]string)
	util.AddToParams(v.SkillName, "SkillName", params)
	return map[string]any{
		"NodeId": key,
		"v":      params,
	}
}
