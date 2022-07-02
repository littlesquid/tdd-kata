package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Greet(names *[]string) string {

	if names == nil {
		return "Hello, my friend."
	}

	greetingTypeToNames := analyzeGreetingType(*names)

	var greetingMsg string = "Hello, "
	var shoutMsg string

	var comma string
	for i, n := range greetingTypeToNames["normal"] {
		n = strings.TrimSpace(n)
		if i == len(greetingTypeToNames["normal"])-1 {
			if len(greetingTypeToNames["normal"]) == 1 {
				greetingMsg += n + "."
				continue
			} else if len(greetingTypeToNames["normal"]) > 2 {
				greetingMsg += ","
			}
			greetingMsg += fmt.Sprintf(" and %v.", n)
		} else {
			greetingMsg += fmt.Sprintf("%v%v", comma, n)
			comma = ", "
		}
	}

	for _, n := range greetingTypeToNames["shoute"] {
		if isUpperCase(n) {
			shoutMsg = fmt.Sprintf("HELLO %v!", n)
			if len(*names) == 1 {
				greetingMsg = ""
			} else {
				greetingMsg += " AND "
			}
		}
	}

	return fmt.Sprintf("%v%v", greetingMsg, shoutMsg)
}

func analyzeGreetingType(input []string) map[string][]string {
	greetingTypeToNames := make(map[string][]string)
	shouteGreet := []string{}
	normalGreet := []string{}
	names := []string{}
	for _, n := range input {
		containingCommaInput := strings.Split(n, ",")
		names = append(names, containingCommaInput...)
	}
	for _, n := range names {
		if isUpperCase(n) {
			shouteGreet = append(shouteGreet, n)
		} else {
			normalGreet = append(normalGreet, n)
		}
	}
	greetingTypeToNames["normal"] = normalGreet
	greetingTypeToNames["shoute"] = shouteGreet

	return greetingTypeToNames
}

func isUpperCase(name string) bool {
	regex := regexp.MustCompile("^[A-Z]+$")
	return regex.MatchString(name)

}
