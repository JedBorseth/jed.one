package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// API routes first
	v1 := r.Group("/api/v1")
	{
		v1.GET("/welcome", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome to the API"})
		})

		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "healthy"})
		})

		v1.GET("/benny", BennyHandler)
	}

	// Serve Astro static site (after /api routes)
	r.Use(staticHandler())

	// Optional: fallback to index.html for client-side routing
	r.NoRoute(func(c *gin.Context) {
		c.File("./docs/static/index.html")
	})

	fmt.Printf("\n\n\033[0;31m Server running on http://localhost:8080\033[0m\n\n")
	r.Run(":8080")
}
