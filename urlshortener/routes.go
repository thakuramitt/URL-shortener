package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	c.File("public/index.html")

}
func shortingUrl(c *gin.Context) {
	db, err := dbcon()
	if err != nil {
		log.Fatalf(err.Error())
	}
	var json struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate the URL
	if !isValidURL(json.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	// Check if the URL already exists in the database
	var existingURL ShortURL
	if result := db.Where("original = ?", json.URL).First(&existingURL); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{"shortened_url": "http://localhost:8080/" + existingURL.Shortened})
		return
	}

	// Generate a random shortened URL string
	shortenedURL := GenerateShortURL()

	// Save the URL in the database
	newURL := ShortURL{Original: json.URL, Shortened: shortenedURL}
	db.Create(&newURL)

	c.JSON(http.StatusOK, gin.H{"shortened_url": "http://localhost:8080/" + shortenedURL})
}
func toRedirect(c *gin.Context) {

	db, err := dbcon()
	if err != nil {
		log.Fatalf(err.Error())
	}

	shortenedURL := c.Param("shortenedURL")
	// Retrieve the original URL from the database
	var url ShortURL
	if result := db.Where("shortened = ?", shortenedURL).First(&url); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.Redirect(http.StatusFound, url.Original)
	c.HTML(http.StatusOK, "public/redirect.html", gin.H{
		"url": url.Original,
	})
}

func geturls(c *gin.Context) {
	db, err := dbcon()
	if err != nil {
		log.Fatal(err)
	}

	var urls []ShortURL
	res := []string{}

	db.Find(&urls, res)
	c.JSON(200, urls)
}
