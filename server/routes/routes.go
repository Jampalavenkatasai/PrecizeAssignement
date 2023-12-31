package routes

import (
	"github.com/gin-gonic/gin"
	"scoreboard-go/server/handlers"
)

func SetupDataRoutes(r *gin.RouterGroup) {

	dataRoutes := r.Group("/data")
	{
		dataRoutes.POST("/insert", handlers.InsertData)
		dataRoutes.GET("/view", handlers.ViewAllData)
		dataRoutes.GET("/single", handlers.ViewSingleData)
		dataRoutes.GET("/rank", handlers.GetRankStudent)
		dataRoutes.PUT("/updatescore", handlers.UpdateScore)
		dataRoutes.DELETE("/deletevalues", handlers.DeleteRecord)
	}
}
