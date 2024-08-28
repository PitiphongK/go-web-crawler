package main

import (
	"fmt"
)

func main() {
    rawBaseURL, maxConcurrency, maxPages := validateArgument()
    conf := newConfigure(rawBaseURL, maxConcurrency, maxPages)
    fmt.Printf("starting crawl: %s\n", rawBaseURL)
    conf.wg.Add(1)
    go conf.crawlPage(rawBaseURL)
    conf.wg.Wait()
    for k, v := range conf.pages {
        fmt.Printf("%v, %v\n", k, v)
    }
    conf.printReport()
}