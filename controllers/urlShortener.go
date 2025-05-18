package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"jed.one/config"
	"jed.one/models"
)

func UrlHandler(c *gin.Context) {
    rawURL := c.Query("url")
    if rawURL == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
        return
    }

    normalizedURL, err := normalizeURL(rawURL)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
        return
    }

    shortCode := GenerateShortCode(normalizedURL)
    baseURL := "https://api.jed.one/url/"

    var url models.URL
    result := config.DB.FirstOrCreate(&url, models.URL{
        OriginalURL: normalizedURL,
        ShortCode:   shortCode,
    })

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



func normalizeURL(rawURL string) (string, error) {
    parsed, err := url.Parse(rawURL)
    if err != nil {
        return "", err
    }

    // Add "https" if no scheme
    if parsed.Scheme == "" {
        parsed.Scheme = "https"
    }

    return parsed.String(), nil
}





