package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	var fullURLs []string
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("could not parse html: '%v'", err)
	}
	var traverseNode func(*html.Node)
	traverseNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, anchor := range n.Attr {
				if anchor.Key == "href" {
					href, err := url.Parse(anchor.Val)
					if err != nil {
						fmt.Printf("couldn't parse href '%v': %v\n", anchor.Val, err)
						continue
					}
					resolvedURL := baseURL.ResolveReference(href)
					fullURLs = append(fullURLs, resolvedURL.String())
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverseNode(c)
		}
	}
	traverseNode(doc)
	return fullURLs, nil
}