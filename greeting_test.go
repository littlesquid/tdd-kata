package main

import (
	"fmt"
	"testing"
)

func TestGreetWithName(t *testing.T) {
	//arrange
	input := []string{"Bob"}
	expected := "Hello, Bob."

	//act
	res := Greet(&input)

	//assert
	assert(t, expected, res)
}

func TestGreetWithNullHandle(t *testing.T) {
	//arrange null input
	expected := "Hello, my friend."

	//act
	res := Greet(nil)

	//assert
	assert(t, expected, res)

}

func TestGreetWithAllUpperCase(t *testing.T) {
	//arrange
	input := []string{"JERRY"}
	expected := "HELLO JERRY!"

	//act
	res := Greet(&input)

	//assert
	assert(t, expected, res)
}

func TestGreetWithTwoNameHandle(t *testing.T) {
	//arrange
	input := []string{"Jill", "Jane"}
	expected := "Hello, Jill and Jane."

	//act
	res := Greet(&input)

	//assert
	assert(t, expected, res)
}

func TestGreetWithArbitaryNumberOfNamesHandle(t *testing.T) {
	//arrange
	input := []string{"Amy", "Brian", "Charlotte"}
	expected := "Hello, Amy, Brian, and Charlotte."

	//act
	res := Greet(&input)

	//assert
	assert(t, expected, res)
}

func TestGreetWithMixingNormalAndShout(t *testing.T) {
	//arrange
	input := []string{"Amy", "BRIAN", "Charlotte"}
	expected := "Hello, Amy and Charlotte. AND HELLO BRIAN!"

	//act
	res := Greet(&input)

	//assert
	assert(t, expected, res)
}

func TestGreetWithContainingComma(t *testing.T) {
	//arrange
	input := []string{"Bob", "Charlie, Dianne"}
	expected := "Hello, Bob, Charlie, and Dianne."

	//act
	res := Greet(&input)

	//assert
	assert(t, expected, res)
}

func assert(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Error(fmt.Sprintf("result should be '%v' but got %v", expected, actual))
	}
}
