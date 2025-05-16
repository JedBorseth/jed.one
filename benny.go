package main

import (
	"fmt"
	"image"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "image/gif"

	_ "image/jpeg"

	"bytes"
	_ "image/png"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"

	"github.com/disintegration/imaging"
)

// 1. Get all files from a directory
func getAllFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

// 2. Check if a file is a valid image and can be decoded
func isValidImage(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	_, _, err = image.Decode(file)
	return err == nil
}

// 3. Get all valid image paths
func getAllValidImages(dir string) ([]string, error) {
	allFiles, err := getAllFiles(dir)
	if err != nil {
		return nil, err
	}

	var validImages []string
	for _, file := range allFiles {
		ext := filepath.Ext(file)
		switch ext {
		case ".jpg", ".jpeg", ".png", ".gif", ".webp":
			if isValidImage(file) {
				validImages = append(validImages, file)
			}
		}
	}
	return validImages, nil
}

// 4. Get a random image from a list
func getRandomImage(images []string) (string, error) {
	if len(images) == 0 {
		return "", fmt.Errorf("no valid images found")
	}

	return images[rand.Intn(len(images))], nil
}






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


