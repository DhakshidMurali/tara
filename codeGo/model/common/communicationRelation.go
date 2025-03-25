package model

import (
	constant "github.com/DhakshidMurali/tara/constant"
	"github.com/DhakshidMurali/tara/util"
)

type CommunicationPostedInCommunity struct {
	Communication Communication
	PostedIn      PostedIn
	Community     Community
}

type CommunicationPostedByUser struct {
	Communication Communication
	PostedBy      PostedBy
	User          User
}

func (v CommunicationPostedInCommunity) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.COMMUNICATION,
			"%r":  constant.POSTEDIN,
			"%n2": constant.COMMUNITY,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_COMMUNICATION_POSTED_IN_COMMUNITY":
		returnData := constant.RETURNDATA_COMMUNICATION
		query := constant.RETRIEVE_DATA_NODE_WHERE
		mapData := map[string]string{
			"%n1":        "COMMUNICATION",
			"%rel":       "POSTEDIN",
			"%n2":        "COMMUNITY",
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

func (v CommunicationPostedInCommunity) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"CommunicationContent": v.Communication.Content,
			"PostedInPostedInDate": v.PostedIn.PostedInDate,
			"CommunityName":        v.Community.Name,
			"CommunityDescription": v.Community.Description,
			"CommunityAccessType":  v.Community.AccessType,
		}
	case "LIST_COMMUNICATION_POSTED_IN_COMMUNITY":
		return map[string]any{
			"NodeId": v.Community.Key,
		}
	default:
		return map[string]any{}
	}
}

func (v CommunicationPostedByUser) MakeQuery(typeOfQuery string) string {
	switch typeOfQuery {
	case "CREATE":
		query := constant.CREATE_NODE_TO_NODE_RELATION
		mapData := map[string]string{
			"%n1": constant.COMMUNICATION,
			"%r":  constant.POSTEDBY,
			"%n2": constant.USEREN1,
		}
		query = util.ReplaceQuery(query, mapData)
		return query
	case "LIST_COMMUNICATION_POSTED_BY_USER":
		returnData := constant.RETURNDATA_COMMUNICATION
		query := constant.RETRIEVE_DATA_NODE_WHERE
		mapData := map[string]string{
			"%n1":        "COMMUNICATION",
			"%rel":       "POSTEDBY",
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

func (v CommunicationPostedByUser) MakeParams(typeOfQuery string) map[string]any {
	switch typeOfQuery {
	case "CREATE":
		return map[string]any{
			"CommunicationLink":    v.Communication.Link,
			"CommunicationContent": v.Communication.Content,
			"PostedByPostedByDate": v.PostedBy.PostedByDate,
			"UserName":             v.User.Name,
			"UserMailAddress":      v.User.MailAddress,
			"UserDomain":           v.User.Domain,
			"UserPassword":         v.User.Password,
			"UserPhoneNumber":      v.User.PhoneNumber,
			"UserUserRole":         v.User.Role,
			"UserLocation":         v.User.Location,
		}
	case "LIST_COMMUNICATION_POSTED_BY_USER":
		return map[string]any{
			"NodeId": v.User.Key,
		}
	default:
		return map[string]any{}
	}
}
