package dashboardRoutes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DhakshidMurali/tara/db"
	dashboardModel "github.com/DhakshidMurali/tara/model/dashboard"
	"github.com/DhakshidMurali/tara/util"
	"github.com/gin-gonic/gin"
)

func listToolGroupByDomain(context *gin.Context) {
	var data dashboardModel.DashboardFields
	query := data.MakeQuery("GET_LIST_TOOLS_GROUPBY_DOMAIN")
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
		dasboardDataList = append(dasboardDataList, data)
	}
	context.JSON(http.StatusOK, dasboardDataList)
	dasboardDataList = nil
}

func listToolsGroupByDeliveryFormatForTop4Domain(context *gin.Context) {
	var data dashboardModel.DashboardFields
	query := data.MakeQuery("GET_LIST_TOOLS_GROUPBY_DELIVERYFORMAT_FOR_TOP4DOMAINS")
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
		dasboardDataList = append(dasboardDataList, data)
	}
	context.JSON(http.StatusOK, dasboardDataList)
	dasboardDataList = nil
}
