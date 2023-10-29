package main

import (
	"github.com/gin-gonic/gin"
	"scoreboard-go/server/config"
	"scoreboard-go/server/routes"
)

func main() {
	r := gin.Default()

	config.SetupDatabase()
	routes.SetupDataRoutes(r.Group("/data"))
	r.Run(":8080")
}
