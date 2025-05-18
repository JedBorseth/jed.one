package main

import (
	"fmt"
	"os"

	"jed.one/config"
	"jed.one/models"
	"jed.one/routes"

	"github.com/gin-gonic/gin"
)
func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"0.0.0.0/0"})
	routes.RegisterRoutes(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}

	fmt.Printf("\n\n\033[0;31m Server running on http://localhost:"+ port +"\033[0m\n\n")
	r.Run(":" + port)
}

func init() {
	gin.SetMode(gin.ReleaseMode)
	config.InitDB()
	config.DB.AutoMigrate(&models.URL{})
}
