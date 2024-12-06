package dashboardRoutes

import (
	"encoding/json"
	"fmt"

	"github.com/DhakshidMurali/tara/db"
	dashboardModel "github.com/DhakshidMurali/tara/model/dashboard"
	"github.com/DhakshidMurali/tara/util"
	"github.com/gin-gonic/gin"
)

func listToolGroupByDomain(context *gin.Context) {
	var data dashboardModel.ToolGroupByDomain
	query := data.MakeQuery()
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
		fmt.Println(data)
	}

}
