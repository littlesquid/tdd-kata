package main

import "fmt"

func FizzBuzz(input int) string {
	if input%3 == 0 && input%5 == 0 {
		return "FizzBuzz"
	} else if input%3 == 0 {
		return "Fizz"
	} else if input%5 == 0 {
		return "Buzz"
	} else {
		return fmt.Sprint(input)
	}
}
