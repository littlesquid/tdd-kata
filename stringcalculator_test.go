package main

import (
	"fmt"
	"testing"
)

func TestStringCalculator(t *testing.T) {
	testcases := []struct {
		in  string
		out int
	}{
		{"", 0},
		{"1", 1},
		{"1,2", 3},
		{"152, 304", 456},
		{"13, 8", 21},
		{"a", 0},
		{"a, b", 0},
		{"5, b", 5},
		{"r, 20", 20},

		{"13\n7", 20},
		{"25\nm", 25},
		{"1,2\n3", 6},
	}

	for _, test := range testcases {
		t.Run(test.in, func(t *testing.T) {
			res, _ := Add(test.in)
			if test.out != res {
				t.Error(fmt.Sprintf("sum of %v should be %v but have %v", test.in, test.out, res))
			}
		})
	}
}

func TestStringCalculatorNegativeNumber(t *testing.T) {
	testcases := []struct {
		in  string
		out string
	}{
		{"-3, 10", "Negative number(s) not allowed: -3"},
		{"2,-4,-9", "Negative number(s) not allowed: -4,-9"},
	}

	for _, test := range testcases {
		t.Run(test.in, func(t *testing.T) {
			_, err := Add(test.in)
			if test.out != err.Error() {
				t.Error(fmt.Sprintf("sum of %v should be error negative number but have %v", test.in, err.Error()))
			}
		})
	}
}

func TestStringCalculatorInvalidCase(t *testing.T) {
	testcases := []struct {
		in  string
		out string
	}{
		{"1,2,3,", "invalid input format"},
		{"25,\n5", "invalid input format"},
		{"//|\n1|2,3", "invalid input format"},
	}

	for _, test := range testcases {
		t.Run(test.in, func(t *testing.T) {
			_, err := Add(test.in)
			if test.out != err.Error() {
				t.Error(fmt.Sprintf("sum of %v should be invalid input format but have %v", test.in, err.Error()))
			}
		})
	}
}

func TestStringCalculatorMultipleDelimiter(t *testing.T) {
	testcases := []struct {
		in  string
		out int
	}{
		{"//;\n;3", 3},
		{"//;\n1;3", 4},
		{"//|\n1|2|3", 6},
		{"//sep\n2sep5", 7},
	}

	for _, test := range testcases {
		t.Run(test.in, func(t *testing.T) {
			res, _ := Add(test.in)
			if test.out != res {
				t.Error(fmt.Sprintf("sum of %v should be %v but have %v", test.in, test.out, res))
			}
		})
	}
}
