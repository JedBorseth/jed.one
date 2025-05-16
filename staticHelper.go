package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func staticHandler() gin.HandlerFunc {
	fs := http.FileServer(http.Dir("./docs/static"))
	return func(c *gin.Context) {
		// Only serve static content if not an API route
		if len(c.Request.URL.Path) >= 4 && c.Request.URL.Path[:4] == "/api" {
			c.Next()
			return
		}
		fs.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}
