package domainRepository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DhakshidMurali/tara/db"
	model "github.com/DhakshidMurali/tara/model/domain"
	"github.com/DhakshidMurali/tara/util"
	"github.com/gin-gonic/gin"
)

var dasboardDataList = model.DashboardPayload{}

func GetDashBoardDetails(context *gin.Context) {
	var data model.DashboardPayload
	query := data.GetQuery("GET_LIST_TOOLS_GROUPBY_DOMAIN")
	params := map[string]any{}
	result := db.Execute(query, params)
	for _, record := range result.Records {
		mapRecord, _ := record.Get("N2")
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &data)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}

	}
	context.JSON(http.StatusOK, dasboardDataList)
}

func ListToolsGroupByDeliveryFormatForTop4Domain(context *gin.Context) {
	var data model.DashboardPayload
	query := data.GetQuery("GET_LIST_TOOLS_GROUPBY_DELIVERYFORMAT_FOR_TOP4DOMAINS")
	params := map[string]any{}
	result := db.Execute(query, params)
	for _, record := range result.Records {
		mapRecord, _ := record.Get("N2")
		byteData := util.MarshalData(mapRecord)
		err := json.Unmarshal(byteData, &data)
		if err != nil {
			fmt.Println("Error Unmarshalling Json")
			panic(err)
		}

	}
	context.JSON(http.StatusOK, dasboardDataList)
}
