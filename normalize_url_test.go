package main

import (
	"testing"
	"fmt"
	"strings"
)

func TestNormalizeUrl(t *testing.T) {
	type testCase struct {
		name			string
		inputURL 		string
		expected		string
		errorContains 	string
	}
	tests := []testCase{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove trailing slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "lowercase capital letters",
			inputURL: "https://BLOG.boot.dev/PATH",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheme and capitals and trailing slash",
			inputURL: "http://BLOG.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:          "handle invalid URL",
			inputURL:      `:\\invalidURL`,
			expected:      "",
			errorContains: "couldn't parse URL",
		},
	}
	passCount := 0
	failCount := 0
	for _, test := range tests {
		output, err := normalizeURL(test.inputURL)
		if err != nil && !strings.Contains(err.Error(), test.errorContains) {
			failCount++
			t.Errorf(`--------------------------------
Inputs:		(%v)
Expecting:	%v
Actual:		%v
Fail: 		%v
`, test.inputURL, test.expected, output, err)
		} else if output != test.expected {
			failCount++
			t.Errorf(`--------------------------------
Inputs:		(%v)
Expecting:	%v
Actual:		%v
Fail
`, test.inputURL, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`--------------------------------
Inputs:		(%v)
Expecting:	%v
Actual:		%v
Pass
`, test.inputURL, test.expected, output)
		}
	}
	fmt.Println("--------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}