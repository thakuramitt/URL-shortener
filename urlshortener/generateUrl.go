package main

import (
	"math/rand"
	"time"
)

func GenerateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortenedURL := make([]byte, 6)
	for i := range shortenedURL {
		shortenedURL[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(shortenedURL)
}
