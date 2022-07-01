package main

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	input := 2
	resp := FizzBuzz(input)
	if fmt.Sprint(input) != resp {
		t.Error(fmt.Sprintf("fizzbuzz of %v should %v but have %v", input, input, resp))
	}
}

func TestFizzBuzz_multiplesOfThree(t *testing.T) {
	input := 9
	resp := FizzBuzz(input)
	if "Fizz" != resp {
		t.Error(fmt.Sprintf("fizzbuzz of %v should 'Fizz' but have %v", input, resp))
	}
}

func TestFizzBuzz_multiplesOfFive(t *testing.T) {
	input := 20
	resp := FizzBuzz(input)
	if "Buzz" != resp {
		t.Error(fmt.Sprintf("fizzbuzz of %v should 'Buzz' but have %v", input, resp))
	}
}

func TestFizzBuzz_multiplesOfThreeAndFive(t *testing.T) {
	input := 30
	resp := FizzBuzz(input)
	if "FizzBuzz" != resp {
		t.Error(fmt.Sprintf("fizzbuzz of %v should 'FizzBuzz' but have %v", input, resp))
	}
}

func TestFizzBuzz_tableDrivenTest(t *testing.T) {
	testcases := []struct {
		in  int
		out string
	}{
		{1, "1"},
		{3, "Fizz"},
		{30, "FizzBuzz"},
		{150, "FizzBuzz"},
		{99, "Fizz"},
		{95, "Buzz"},
		{65, "Buzz"},
	}

	for _, test := range testcases {
		t.Run(test.out, func(t *testing.T) {
			res := FizzBuzz(test.in)
			if test.out != res {
				t.Error(fmt.Sprintf("fizzbuzz of %v should %v but have %v", test.in, test.out, res))
			}
		})
	}
}
