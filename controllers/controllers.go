package controllers

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)



func BennyHandler(c *gin.Context) {
	// Parse query parameters
	widthStr := c.Query("width")
	heightStr := c.Query("height")

	var width, height int
	var err error

	if widthStr != "" {
		width, err = strconv.Atoi(widthStr)
		if err != nil || width <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid width value",
			})
			return
		}
	}

	if heightStr != "" {
		height, err = strconv.Atoi(heightStr)
		if err != nil || height <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid height value",
			})
			return
		}
	}

	// Get valid images
	images, err := getAllValidImages("./images")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to retrieve images",
		})
		return
	}

	// Select random image
	randomImg, err := getRandomImage(images)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to retrieve random image",
		})
		return
	}

	// Open the image
	src, err := imaging.Open(randomImg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to open image",
		})
		return
	}

	// Resize if dimensions were provided
	if width > 0 || height > 0 {
		// If one dimension is missing, preserve aspect ratio
		if width == 0 {
			ratio := float64(height) / float64(src.Bounds().Dy())
			width = int(float64(src.Bounds().Dx()) * ratio)
		} else if height == 0 {
			ratio := float64(width) / float64(src.Bounds().Dx())
			height = int(float64(src.Bounds().Dy()) * ratio)
		}

		src = imaging.Resize(src, width, height, imaging.Lanczos)
	}

	// Encode to memory and stream to client
	var buf bytes.Buffer
	err = imaging.Encode(&buf, src, imaging.PNG)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to encode image",
		})
		return
	}

	c.Header("Content-Type", "image/png")
	c.Data(http.StatusOK, "image/png", buf.Bytes())
}


func UrlHandler(c *gin.Context) {
	
}