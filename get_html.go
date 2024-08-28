package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("could not get from '%s': %v", rawURL, err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("could not read body of '%s': %v", rawURL, err)
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		return "", fmt.Errorf("error status code '%d", res.StatusCode)
	}
	if res.Header.Get("Content-type") == "text/html" {
		return "", fmt.Errorf("invalid content-type '%s', expected 'text/html'", res.Header.Get("Content-type"))
	}
	return string(body), nil
}