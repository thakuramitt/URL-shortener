package main

import "net/url"

func isValidURL(urlString string) bool {
	parsedURL, err := url.Parse(urlString)
	if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		return false
	}

	return true
}
