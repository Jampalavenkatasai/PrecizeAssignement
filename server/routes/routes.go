package routes

import (
	"github.com/gin-gonic/gin"
	"scoreboard-go/server/handlers"
)

func SetupDataRoutes(r *gin.RouterGroup) {
	dataRoutes := r.Group("/data")
	{
		dataRoutes.POST("/", handlers.InsertData)
		dataRoutes.GET("/", handlers.ViewAllData)
		dataRoutes.PUT("/:name", handlers.UpdateScore)
		dataRoutes.DELETE("/:name", handlers.DeleteRecord)
	}
}
