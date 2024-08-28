package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

func validateArgument() (rawBaseURL string, maxConcurrency, maxPages int) {
    if len(os.Args) == 1 {
        fmt.Println("no website provided")
        os.Exit(1)
    } else if len(os.Args) > 4 {
        fmt.Println("too many arguments provided")
        os.Exit(1)
    }
    rawBaseURL = os.Args[1]
    _, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("couldn't parse base URL: %v\n", err)
        os.Exit(1)
	}
    maxConcurrency, err = strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Printf("couldn't parse maxConcurrency: %v\n", err)
        os.Exit(1)
    }
    maxPages, err = strconv.Atoi(os.Args[3])
    if err != nil {
        fmt.Printf("couldn't parse maxPages: %v\n", err)
        os.Exit(1)
    }
    return rawBaseURL, maxConcurrency, maxPages
}