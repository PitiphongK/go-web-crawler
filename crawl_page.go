package main

import (
	"fmt"
	"net/url"
)

func (conf *config) crawlPage(rawCurrentURL string) {
	defer func(){ 
		<- conf.concurrencyControl
		conf.wg.Done()
	}()
	conf.concurrencyControl <- struct{}{}

	if isReached := conf.isMaxPagesReached(); isReached {
		return
	}

	err := hasSameBaseURL(conf.baseURL, rawCurrentURL)
	if err != nil {
		// fmt.Println(err)
		return
	}
	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
		return
	}

	if isFirst := conf.addPageVisit(normalizedURL); !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}
	urls, err := getURLsFromHTML(htmlBody, conf.baseURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v", err)
		return
	}
	for _, url := range urls {
		conf.wg.Add(1)
		go conf.crawlPage(url)
	}
}

func (conf *config) isMaxPagesReached() (isReached bool){
	defer conf.mu.Unlock()
	conf.mu.Lock()
	return len(conf.pages) >= conf.maxPages
}

func (conf *config) addPageVisit(normalizedURL string) (isFirst bool) {
	defer conf.mu.Unlock()
	conf.mu.Lock()
	// increment if visited
	if _, visited := conf.pages[normalizedURL]; visited {
		conf.pages[normalizedURL]++
		return false
	}
	// mark as visited
	conf.pages[normalizedURL] = 1
	return true
}

func hasSameBaseURL(rawBaseURL, rawURL string) error {
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return err
	}
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	if parsedBaseURL.Host == parsedURL.Host {
		return nil
	}
	return fmt.Errorf("current and base url have different domain '%v' and '%v'", parsedBaseURL.Host, parsedURL.Host)
}