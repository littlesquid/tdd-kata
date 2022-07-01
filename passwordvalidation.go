package main

import (
	"regexp"
)

type PasswordValidationResult struct {
	IsValid      bool
	ErrorMessage *string
}

func PasswordValidation(input string) PasswordValidationResult {
	invalidCharacterLong := "Password must be at least 8 characters"
	invalidNumberCount := "The password must contain at least 2 numbers"
	invalidCapitalCount := "password must contain at least one capital letter"
	invalidSpecialCharacterCount := "password must contain at least one special character"

	var errorMsgs []string

	if len(input) < 8 {
		errorMsgs = append(errorMsgs, invalidCharacterLong)
	}
	if !isContain(input, "[0-9]+") {
		errorMsgs = append(errorMsgs, invalidNumberCount)
	}

	if !isContain(input, "[A-Z]+") {
		errorMsgs = append(errorMsgs, invalidCapitalCount)
	}

	if !isContain(input, "[!|@|#|$|%|^|&|*|(|)|-|=|+]+") {
		errorMsgs = append(errorMsgs, invalidSpecialCharacterCount)
	}

	if len(errorMsgs) > 0 {
		var message string
		for i, msg := range errorMsgs {
			message += msg
			if i < len(errorMsgs)-1 {
				message += ","
			}
		}
		return SetValidationResult(false, &message)
	}

	return SetValidationResult(true, nil)
}

func SetValidationResult(isValid bool, errorMsg *string) PasswordValidationResult {
	return PasswordValidationResult{
		IsValid:      isValid,
		ErrorMessage: errorMsg,
	}
}

func isContain(input string, expression string) bool {
	regex := regexp.MustCompile(expression)
	numbers := regex.FindAllString(input, -1)
	if len(numbers) > 0 {
		return true
	}
	return false
}
