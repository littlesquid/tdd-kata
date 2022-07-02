package main

import (
	"regexp"
)

func IsMatch(input string, expression string) bool {
	regex := regexp.MustCompile(expression)
	return regex.Match([]byte(input))
}
