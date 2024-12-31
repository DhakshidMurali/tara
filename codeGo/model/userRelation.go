package model

import (
	"github.com/DhakshidMurali/tara/constant"
	"github.com/DhakshidMurali/tara/util"
)

type UserCollaboratedWithUser struct {
	UserN1           User
	CollaboratedWith CollaboratedWith
	UserN2           User
}

type UserWorksInTools struct {
	User    User
	WorksIn WorksIn
	Tool    Tool
}

type UserSkilledInSkills struct {
	User      User
	SkilledIn SkilledIn
	Skills    Skills
}

type UserReportToUser struct {
	User    User
	Manager User
}
type UserComesUnderDomain struct {
	User   User
	Domain Domain
}

func (v UserCollaboratedWithUser) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.USEREN1,
			"%r":  constant.COLLABORATEDWITH,
			"%n2": constant.USEREN2,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_USER_COLLABORATED_WITH":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_USER
		mapData := map[string]string{
			"%n1":        "User",
			"%rel":       "CollaboratedWith",
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

func (v UserCollaboratedWithUser) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"UserName":                        v.UserN1.Name,
			"UserMailAddress":                 v.UserN1.MailAddress,
			"UserDomain":                      v.UserN1.Domain,
			"UserPassword":                    v.UserN1.Password,
			"UserPhoneNumber":                 v.UserN1.PhoneNumber,
			"UserUserRole":                    v.UserN1.Role,
			"UserLocation":                    v.UserN1.Location,
			"CollaboratedWithTaskDescription": v.CollaboratedWith.TaskDescription,
			"UserNameN2":                      v.UserN2.Name,
			"UserMailAddressN2":               v.UserN2.MailAddress,
			"UserDomainN2":                    v.UserN2.Domain,
			"UserPasswordN2":                  v.UserN2.Password,
			"UserPhoneNumberN2":               v.UserN2.PhoneNumber,
			"UserUserRoleN2":                  v.UserN2.Role,
			"UserLocationN2":                  v.UserN2.Location,
		}
	case "LIST_USER_COLLABORATED_WITH":
		return map[string]any{
			"NodeId": v.UserN1.Key,
		}
	default:
		return map[string]any{}
	}
}

func (v UserWorksInTools) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.USEREN1,
			"%r":  constant.WORKSIN,
			"%n2": constant.TOOL,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_TOOLS_USER_WORKS_IN":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_TOOL
		mapData := map[string]string{
			"%n1":        "User",
			"%rel":       "WorksIn",
			"%n2":        "Tool",
			"%condition": "elementId(n1)=$NodeId",
			"%node":      "n2",
			"%return":    returnData,
		}
		query = util.DoubleReplaceQuery(query, mapData)
		return query
	case "LIST_USERS_WORKS_IN_TOOL":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_USER
		mapData := map[string]string{
			"%n1":        "User",
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

func (v UserWorksInTools) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"UserName":           v.User.Name,
			"UserMailAddress":    v.User.MailAddress,
			"UserDomain":         v.User.Domain,
			"UserPassword":       v.User.Password,
			"UserPhoneNumber":    v.User.PhoneNumber,
			"UserUserRole":       v.User.Role,
			"UserLocation":       v.User.Location,
			"WorksInWorksInRole": v.WorksIn.Role,
			"ToolName":           v.Tool.Name,
			"ToolDeliveryFormat": v.Tool.DeliveryFormat,
			"ToolApprovalType":   v.Tool.ApprovalType,
		}
	case "LIST_TOOLS_USER_WORKS_IN":
		return map[string]any{
			"NodeId": v.User.Key,
		}
	case "LIST_USERS_WORKS_IN_TOOL":
		return map[string]any{
			"NodeId": v.Tool.Key,
		}
	default:
		return map[string]any{}
	}
}

func (v UserSkilledInSkills) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.USEREN1,
			"%r":  constant.SKILLEDIN,
			"%n2": constant.SKILLS,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_USER_SKILLED_IN":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_USER
		mapData := map[string]string{
			"%n1":        "User",
			"%rel":       "SkilledIn",
			"%n2":        "Skills",
			"%condition": "elementId(n2)=$NodeId",
			"%node":      "n1",
			"%return":    returnData,
		}
		query = util.DoubleReplaceQuery(query, mapData)
		return query
	case "LIST_SKILLS_SKILLED_BY_USER":
		query := constant.RETRIEVE_DATA_NODE_WHERE
		returnData := constant.RETURNDATA_SKILLS
		mapData := map[string]string{
			"%n1":        "User",
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

func (v UserSkilledInSkills) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"UserName":            v.User.Name,
			"UserMailAddress":     v.User.MailAddress,
			"UserDomain":          v.User.Domain,
			"UserPassword":        v.User.Password,
			"UserPhoneNumber":     v.User.PhoneNumber,
			"UserUserRole":        v.User.Role,
			"UserLocation":        v.User.Location,
			"SkilledInExperience": v.SkilledIn.Experience,
			"SkillsSkills":        v.Skills.SkillName,
		}
	case "LIST_USER_SKILLED_IN":
		return map[string]any{
			"NodeId": v.Skills.Key,
		}
	case "LIST_SKILLS_SKILLED_BY_USER":
		return map[string]any{
			"NodeId": v.User.Key,
		}
	default:
		return map[string]any{}
	}
}
