package main

import (
	"miiwo/skyfarer/backend/models"
	"miiwo/skyfarer/handler"
	"miiwo/skyfarer/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// INIT
	router := gin.Default()
	godotenv.Load(".env")
	models.ConnectToMariaDB()

	// ROUTES
	v1 := router.Group("v1", middleware.CorsConfig, middleware.ValidateAPIKey)
	v1.GET("/weapons", handler.FetchWeaponsFromDB)

	// Just to show routes are working without authentication
	router.GET("/ping", func(gctx *gin.Context) {
		gctx.IndentedJSON(200, gin.H{"message": "pong"})
	})

	// RUN
	router.Run(os.Getenv("BASE_URL"))
}
