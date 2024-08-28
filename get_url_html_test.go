package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		name			string
		inputURL		string
		inputBody		string
		expected		[]string
		errorContains	string
	}
	tests := []testCase{
		{
			name:     "absolute URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="https://blog.boot.dev">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev"},
		},
		{
			name:     "relative URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no href",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a>
			<span>Boot.dev></span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
		{
			name:     "bad HTML",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html body>
	<a href="path/one">
		<span>Boot.dev></span>
	</a>
</html body>
`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:     "invalid href URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href=":\\invalidURL">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
		{
			name:     "handle invalid base URL",
			inputURL: `:\\invalidBaseURL`,
			inputBody: `
<html>
	<body>
		<a href="/path">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected:      nil,
			errorContains: "couldn't parse base URL",
		},
	}
	passCount := 0
	failCount := 0
	for i, test := range tests {
		actual, err := getURLsFromHTML(test.inputBody, test.inputURL)
		if err != nil && !strings.Contains(err.Error(), test.errorContains){
			failCount++
			t.Errorf(`---------------------------------
Test %v - '%s'
Input: 		(%v, %v)
Expected:	%v
Actual:		%v
Fail: 		expected error contains '%v', but got '%v'
`, i, test.name, test.inputBody, test.inputURL, test.expected, actual, test.errorContains, err)
		} else if err != nil && test.errorContains == "" {
			failCount++
			t.Errorf(`---------------------------------
Test %v - '%s'
Input: 		(%v, %v)
Expected:	%v
Actual:		%v
Fail: 		unexpected error: '%v'
`, i, test.name, test.inputBody, test.inputURL, test.expected, actual, err)
		} else if err == nil && test.errorContains != "" {
			failCount++
			t.Errorf(`---------------------------------
Test %v - '%s'
Input: 		(%v, %v)
Expected:	%v
Actual:		%v
Fail: 		expected error contains '%v', but got none
`, i, test.name, test.inputBody, test.inputURL, test.expected, actual, test.errorContains)
		} else if !reflect.DeepEqual(actual, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Test %v - '%s'
Input: 		(%v, %v)
Expected:	%v
Actual:		%v
Fail
`, i, test.name, test.inputBody, test.inputURL, test.expected, actual)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test %v - '%s'
Input: 		(%v, %v)
Expected:	%v
Actual:		%v
Pass
`, i, test.name, test.inputBody, test.inputURL, test.expected, actual)
		}
	}
}

/*

Fail Error Case Map
	Result 						vs	Actual
1. 	expects a specific error 		got a differnet error
2. 	expect an error					got no error
3. 	does not expect error			got an error

*/