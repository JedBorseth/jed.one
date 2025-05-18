package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jed.one/controllers"
)


func RegisterRoutes(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
	c.File("./docs/static/404.html")
})
	v1 := r.Group("/api/v1")
	{
		v1.GET("/welcome", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome to the API"})
		})

		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "healthy"})
		})

		v1.GET("/benny", controllers.BennyHandler)

		v1.GET("/url", controllers.UrlHandler)
	}
	r.GET("/url/:shortCode", controllers.RedirectToOriginalURL)

	r.Use(staticHandler("./docs/static"))
	
	
	
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"status":  "error",
			"message": "Method not allowed",
		})
	})
}