package main

import (
	"fmt"
	"regexp"
	"testing"
)

func IsMatch(input string, expression string) bool {
	regex := regexp.MustCompile(expression)
	return regex.Match([]byte(input))
}

func Assert(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Error(fmt.Sprintf("result should be '%v' but got %v", expected, actual))
	}
}
