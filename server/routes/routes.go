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
		dataRoutes.PUT("/:name", handlers.UpdateScore)
		dataRoutes.DELETE("/:name", handlers.DeleteRecord)
	}
}
