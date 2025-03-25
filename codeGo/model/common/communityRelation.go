package model

import (
	"github.com/DhakshidMurali/tara/constant"
	"github.com/DhakshidMurali/tara/util"
)

type CommunityMemberUser struct {
	Community Community
	Member    Member
	User      User
}

type CommunityCreatedByUser struct {
	Community Community
	CreatedBy CreatedBy
	User      User
}

func (v CommunityMemberUser) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.COMMUNITY,
			"%r":  constant.MEMBER,
			"%n2": constant.USEREN1,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_USERS_MEMBER_OF_COMMUNITY":
		returnData := constant.RETURNDATA_USER
		query := constant.RETRIEVE_DATA_NODE_WHERE
		mapData := map[string]string{
			"%n1":        "COMMUNITY",
			"%rel":       "MEMBER",
			"%n2":        "USER",
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

func (v CommunityMemberUser) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"CommunityName":              v.Community.Name,
			"CommunityDescription":       v.Community.Description,
			"CommunityAccessType":        v.Community.AccessType,
			"MemberAddedToCommunityDate": v.Member.AddedToCommunityDate,
			"UserName":                   v.User.Name,
			"UserMailAddress":            v.User.MailAddress,
			"UserDomain":                 v.User.Domain,
			"UserPassword":               v.User.Password,
			"UserPhoneNumber":            v.User.PhoneNumber,
			"UserUserRole":               v.User.Role,
			"UserLocation":               v.User.Location,
		}
	case "LIST_USERS_MEMBER_OF_COMMUNITY":
		return map[string]any{
			"NodeId": v.Community.Key,
		}
	default:
		return map[string]any{}
	}
}

func (v CommunityCreatedByUser) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.COMMUNITY,
			"%r":  constant.CREATEDBY,
			"%n2": constant.USEREN1,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_COMMUNITY_CREATED_BY_USER":
		returnData := constant.RETURNDATA_COMMUNITY
		query := constant.RETRIEVE_DATA_NODE_WHERE
		mapData := map[string]string{
			"%n1":        "COMMMUNITY",
			"%rel":       "CREATEDBY",
			"%n2":        "USER",
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

func (v CommunityCreatedByUser) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"CommunityName":                 v.Community.Name,
			"CommunityDescription":          v.Community.Description,
			"CommunityAccessType":           v.Community.AccessType,
			"CreatedByCommunityCreatedDate": v.CreatedBy.CommunityCreatedDate,
			"UserName":                      v.User.Name,
			"UserMailAddress":               v.User.MailAddress,
			"UserDomain":                    v.User.Domain,
			"UserPassword":                  v.User.Password,
			"UserPhoneNumber":               v.User.PhoneNumber,
			"UserUserRole":                  v.User.Role,
			"UserLocation":                  v.User.Location,
		}
	case "LIST_COMMUNITY_CREATED_BY_USER":
		return map[string]any{
			"NodeId": v.User.Key,
		}
	default:
		return map[string]any{}
	}
}
