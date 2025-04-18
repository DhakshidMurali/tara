package toolRepository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DhakshidMurali/tara/db"
	model "github.com/DhakshidMurali/tara/model/tool"
	"github.com/DhakshidMurali/tara/util"
	"github.com/gin-gonic/gin"
)

var toolDataList = []model.ToolPageDetails{}

func ListToolGroupByCommunication(context *gin.Context) {
	var data model.ToolPageDetails
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
		toolDataList = append(toolDataList, data)
	}
	context.JSON(http.StatusOK, toolDataList)
	toolDataList = nil
}
