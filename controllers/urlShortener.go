package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"jed.one/config"
	"jed.one/models"
)

func UrlHandler(c *gin.Context) {

    longURL := c.Query("url")
    if longURL == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
        return
    }

    shortCode := GenerateShortCode(longURL)
    baseURL := "https://api.jed.one/url/"


    var url models.URL
    result := config.DB.FirstOrCreate(&url, models.URL{OriginalURL: longURL, ShortCode: shortCode})
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "shortened_url": baseURL + url.ShortCode,
    })
}

func GenerateShortCode(url string) string {
    hash := sha1.Sum([]byte(url))
    return base64.URLEncoding.EncodeToString(hash[:])[:6]
}

func RedirectToOriginalURL(c *gin.Context) {
    shortCode := c.Param("shortCode") 

    var url models.URL
    result := config.DB.Where("short_code = ?", shortCode).First(&url)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
        return
    }

    c.Redirect(http.StatusFound, url.OriginalURL)
}





