package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type Result struct {
	page	string
	links	int
}

func (conf *config) printReport() {
	fmt.Printf(`=============================
  REPORT for %s
=============================`, conf.baseURL)
	fmt.Println()
	results := []Result{}
	for k, v := range conf.pages {
		results = append(results, Result{k, v})
	}
	slices.SortFunc(results, func(a, b Result) int {
		// Descending order
		if n := cmp.Compare(a.links, b.links); n != 0 {
			return n * -1
		}
		// If links are equal, order by page
		return strings.Compare(a.page, b.page)
	})
	for _, result := range results {
		fmt.Printf("Found %d internal links to %s\n", result.links, result.page)
	}
}