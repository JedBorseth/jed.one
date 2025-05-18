package controllers

import (
	"fmt"
	"image"

	_ "image/gif"

	_ "image/jpeg"

	_ "image/png"
	"math/rand"
	"os"
	"path/filepath"
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








