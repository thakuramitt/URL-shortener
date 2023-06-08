package main

import (
	"github.com/gin-gonic/gin"
)

type ShortURL struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Original  string `gorm:"unique" json:"original"`
	Shortened string `gorm:"unique" json:"shortened"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", Homepage)
	router.Static("/public", "./public")

	router.POST("/shorten", shortingUrl)
	router.GET("/:shortenedURL", toRedirect)
	router.GET("/geturls", geturls)

	router.Run(":8080")
}
