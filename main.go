package main

import (
	"fmt"
	"os"

	"jed.one/routes"

	"github.com/gin-gonic/gin"
)
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies([]string{"0.0.0.0/0"})
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}
	routes.RegisterRoutes(r)


	fmt.Printf("\n\n\033[0;31m Server running on http://localhost:"+ port +"\033[0m\n\n")
	r.Run(":" + port) // listen and serve on
}
