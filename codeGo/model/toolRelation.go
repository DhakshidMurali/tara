package model

import (
	"github.com/DhakshidMurali/tara/constant"
	"github.com/DhakshidMurali/tara/util"
)

type ToolRequestByUser struct {
	Tool      Tool
	RequestBy RequestBy
	User      User
}

type ToolAccessToUser struct {
	Tool     Tool
	AccessTo AccessTo
	User     User
}

type ToolManagedByUser struct {
	Tool Tool
	User User
}

type ToolComesUnderDomain struct {
	Tool   Tool
	Domain Domain
}

func (v ToolRequestByUser) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.TOOL,
			"%r":  constant.REQUESTBY,
			"%n2": constant.USEREN1,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_TOOL_REQUESTED_BY_USER":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_TOOL
		mapData := map[string]string{
			"%n1":        "Tool",
			"%rel":       "RequestBy",
			"%n2":        "User",
			"%condition": "elementId(n2)=$NodeId",
			"%node":      "n1",
			"%return":    returnData,
		}
		query = util.DoubleReplaceQuery(query, mapData)
		return query
	case "LIST_USERS_REQUESTED_TO_TOOL":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_USER
		mapData := map[string]string{
			"%n1":        "Tool",
			"%rel":       "RequestBy",
			"%n2":        "User",
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

func (v ToolRequestByUser) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"ToolName":                     v.Tool.Name,
			"ToolDeliveryFormat":           v.Tool.DeliveryFormat,
			"RequestByRequestedDate":       v.RequestBy.RequestedDate,
			"RequestByRequestByAccessType": v.RequestBy.AccessType,
			"RequestByLvl1Approval":        v.RequestBy.Lvl1Approval,
			"RequestByLvl2Approval":        v.RequestBy.Lvl2Approval,
			"UserName":                     v.User.Name,
			"UserMailAddress":              v.User.MailAddress,
			"UserDomain":                   v.User.Domain,
			"UserPassword":                 v.User.Password,
			"UserPhoneNumber":              v.User.PhoneNumber,
			"UserUserRole":                 v.User.Role,
			"UserLocation":                 v.User.Location,
		}
	case "LIST_TOOL_REQUESTED_BY_USER":
		return map[string]any{
			"NodeId": v.User.Key,
		}
	case "LIST_USERS_REQUESTED_TO_TOOL":
		return map[string]any{
			"NodeId": v.Tool.Key,
		}
	default:
		return map[string]any{}
	}
}

func (v ToolAccessToUser) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.TOOL,
			"%r":  constant.ACCESSTO,
			"%n2": constant.USEREN1,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_TOOLS_ACCESS_BY_USER":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_TOOL
		mapData := map[string]string{
			"%n1":        "Tool",
			"%rel":       "AccessTo",
			"%n2":        "User",
			"%condition": "elementId(n2)=$NodeId",
			"%node":      "n1",
			"%return":    returnData,
		}
		query = util.DoubleReplaceQuery(query, mapData)
		return query
	case "LIST_USERS_ACCESS_TO_TOOL":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_USER
		mapData := map[string]string{
			"%n1":        "Tool",
			"%rel":       "AccessTo",
			"%n2":        "User",
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

func (v ToolAccessToUser) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"ToolName":            v.Tool.Name,
			"ToolDeliveryFormat":  v.Tool.DeliveryFormat,
			"RequestByAccessType": v.AccessTo.AccessType,
			"UserName":            v.User.Name,
			"UserMailAddress":     v.User.MailAddress,
			"UserDomain":          v.User.Domain,
			"UserPassword":        v.User.Password,
			"UserPhoneNumber":     v.User.PhoneNumber,
			"UserUserRole":        v.User.Role,
			"UserLocation":        v.User.Location,
		}
	case "LIST_TOOLS_ACCESS_BY_USER":
		return map[string]any{
			"NodeId": v.User.Key,
		}
	case "LIST_USERS_ACCESS_TO_TOOL":
		return map[string]any{
			"NodeId": v.Tool.Key,
		}
	default:
		return map[string]any{}
	}
}

func (v ToolManagedByUser) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.TOOL,
			"%r":  constant.MANAGEDBY,
			"%n2": constant.USEREN1,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_TOOLS_MANAGED_BY_USER":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_TOOL
		mapData := map[string]string{
			"%n1":        "Tool",
			"%rel":       "ManagedBy",
			"%n2":        "User",
			"%condition": "elementId(n2)=$NodeId",
			"%node":      "n1",
			"%return":    returnData,
		}
		query = util.DoubleReplaceQuery(query, mapData)
		return query
	case "LIST_USERS_MANAGES_TOOL":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_USER
		mapData := map[string]string{
			"%n1":        "Tool",
			"%rel":       "ManagedBy",
			"%n2":        "User",
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
func (v ToolManagedByUser) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"ToolName":           v.Tool.Name,
			"ToolDeliveryFormat": v.Tool.DeliveryFormat,
			"UserName":           v.User.Name,
			"UserMailAddress":    v.User.MailAddress,
			"UserDomain":         v.User.Domain,
			"UserPassword":       v.User.Password,
			"UserPhoneNumber":    v.User.PhoneNumber,
			"UserUserRole":       v.User.Role,
			"UserLocation":       v.User.Location,
		}
	case "LIST_TOOLS_MANAGED_BY_USER":
		return map[string]any{
			"NodeId": v.User.Key,
		}
	case "LIST_USERS_MANAGES_TOOL":
		return map[string]any{
			"NodeId": v.Tool.Key,
		}
	default:
		return map[string]any{}
	}
}
func (v ToolComesUnderDomain) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.TOOL,
			"%r":  constant.COMESUNDER,
			"%n2": constant.DOMAIN,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_TOOLS_COMES_UNDER_DOMAIN":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_TOOL
		mapData := map[string]string{
			"%n1":        "Tool",
			"%rel":       "ComesUnder",
			"%n2":        "Domain",
			"%condition": "elementId(n2)=$NodeId",
			"%node":      "n1",
			"%return":    returnData,
		}
		query = util.DoubleReplaceQuery(query, mapData)
		return query
	case "LIST_DOMAIN_MAINTAIN_TOOL":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_DOMAIN
		mapData := map[string]string{
			"%n1":        "Tool",
			"%rel":       "ComesUnder",
			"%n2":        "Domain",
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

func (v ToolComesUnderDomain) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"ToolName":           v.Tool.Name,
			"ToolDeliveryFormat": v.Tool.DeliveryFormat,

			"DomainName": v.Domain.DomainName,
		}
	case "LIST_TOOLS_COMES_UNDER_DOMAIN":
		return map[string]any{
			"NodeId": v.Domain.Key,
		}
	case "LIST_DOMAIN_MAINTAIN_TOOL":
		return map[string]any{
			"NodeId": v.Tool.Key,
		}
	default:
		return map[string]any{}
	}
}
