package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Add(number string) (int, error) {
	if "" == number {
		return 0, nil
	}

	if isBlock(number) {
		return 0, errors.New("invalid input format")
	}

	var inputs []string
	if isMultiDelimiter(number) {
		inputs = processMultiDelimiter(number)
	} else {
		inputs = strings.FieldsFunc(number, split)
	}

	if len(inputs) == 1 {
		resp, err := strconv.Atoi(number)
		if err != nil {
			return 0, nil
		}
		return resp, nil
	}

	var sum = 0
	var negative []string

	for i, input := range inputs {
		trimInput := strings.Trim(input, " ")
		if !isNumberOrLetter(trimInput) {
			return 0, errors.New("invalid input format")
		}
		n, _ := strconv.Atoi(trimInput)

		if n < 0 {
			negative = append(negative, fmt.Sprint(n))
		}
		if negative != nil && i == len(inputs)-1 {
			return 0, errors.New(fmt.Sprintf("Negative number(s) not allowed: %v", strings.Join(negative, ",")))
		}

		sum += n
	}
	return sum, nil
}

func split(r rune) bool {
	return r == ',' || r == '\n'
}

func isNumberOrLetter(input string) bool {
	regex := regexp.MustCompile("^[a-zA-Z0-9_.-]*$")
	return regex.Match([]byte(input))
}

func isBlock(input string) bool {
	invalidConcanateSymbol := regexp.MustCompile(".+,\\n\\d+")
	invalidInputFormat := regexp.MustCompile("[,]$")

	return invalidConcanateSymbol.Match([]byte(input)) || invalidInputFormat.Match([]byte(input))
}

func isMultiDelimiter(input string) bool {
	if len(input) > 1 {
		return "//" == input[:2]
	}
	return false
}

func processMultiDelimiter(input string) []string {
	inputs := strings.Split(input, "\n")

	delimiter := inputs[0]

	delimiter = strings.ReplaceAll(delimiter, "//", "")

	number := inputs[1]

	return strings.Split(number, delimiter)
}
