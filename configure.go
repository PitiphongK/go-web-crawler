package main

import "sync"

type config struct {
	pages              map[string]int
	baseURL            string
	mu                 sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
    maxPages           int
}

func newConfigure(rawBaseURL string, maxConcurrency, maxPages int) *config {
	return &config{
        pages: make(map[string]int),
        baseURL: rawBaseURL,
        concurrencyControl: make(chan struct{}, maxConcurrency),
        wg: &sync.WaitGroup{},
        maxPages: maxPages,
    }
}