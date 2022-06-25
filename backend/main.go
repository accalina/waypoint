package main

import (
	"waypoint/src"
	"waypoint/src/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	models.ConnectDatabase()
	src.GetRoute(r)

	r.Run(":8080")
}
