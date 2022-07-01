package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Search(city string) []string {
	cities := []string{"Paris", "Budapest", "Skopje", "Rotterdam", "Valencia", "Vancouver", "Amsterdam",
		"Vienna", "Sydney", "New York City", "London", "Bangkok", "Hong Kong",
		"Dubai", "Rome", "Istanbul"}

	var result []string

	if city == "*" {
		return cities
	}

	if len(city) < 2 {
		return append(result, "no result")
	}
	firstPositionExp := fmt.Sprintf("^(%v)\\w+", strings.Title(strings.ToLower(city)))
	containExp := fmt.Sprintf("\\w+(%v)\\w+", strings.ToLower(city))

	fmt.Println(firstPositionExp)
	fmt.Println(containExp)

	firstPositionRegex := regexp.MustCompile(firstPositionExp)
	containRegex := regexp.MustCompile(containExp)
	for _, c := range cities {
		if firstPositionRegex.Match([]byte(c)) || containRegex.Match([]byte(c)) {
			result = append(result, c)
		}
	}

	return result
}
