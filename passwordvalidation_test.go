package main

import (
	"fmt"
	"testing"
)

func TestPasswordValidation(t *testing.T) {
	invalidCharacterLong := "Password must be at least 8 characters"
	invalidNumberCount := "The password must contain at least 2 numbers"
	invalidCapitalCount := "password must contain at least one capital letter"
	invalidSpecialCharacterCount := "password must contain at least one special character"

	allMessageErrorHandle := fmt.Sprintf("%v,%v,%v,%v", invalidCharacterLong, invalidNumberCount, invalidCapitalCount, invalidSpecialCharacterCount)
	invalidCharacterLongAndNumberErrorHandle := fmt.Sprintf("%v,%v", invalidCharacterLong, invalidNumberCount)

	testcases := []struct {
		in  string
		out PasswordValidationResult
	}{
		{"1P@ssw0rd", SetValidationResult(true, nil)},
		{"invalid", SetValidationResult(false, &allMessageErrorHandle)},
		{"p@sswordwithoutnumbeR", SetValidationResult(false, &invalidNumberCount)},
		{"p$sworD", SetValidationResult(false, &invalidCharacterLongAndNumberErrorHandle)},
		{"1Passw0rd", SetValidationResult(false, &invalidSpecialCharacterCount)},
	}

	for _, test := range testcases {
		t.Run(test.in, func(t *testing.T) {
			res := PasswordValidation(test.in)
			if test.out.IsValid != res.IsValid {
				t.Error(fmt.Sprintf("password %v should be %v but got %v", test.in, test.out.IsValid, res.IsValid))
			}
			if test.out.ErrorMessage != nil && *test.out.ErrorMessage != *res.ErrorMessage {
				t.Error(fmt.Sprintf("result of validate [%v] should be %v but got %v", test.in, *test.out.ErrorMessage, *res.ErrorMessage))
			}
		})
	}
}
