package main

import (
	"fmt"
	"testing"
)

func TestSearchWithNoResult(t *testing.T) {
	input := "A"
	res := Search(input)
	if "no result" != res[0] {
		t.Error(fmt.Sprintf("search response of %v should be %v but have %v", input, "no result", res[0]))
	}
}

func TestSearch(t *testing.T) {
	expected := []string{"Valencia", "Vancouver"}
	input := "Va"
	res := Search(input)
	if len(expected) != len(res) {
		t.Error(fmt.Sprintf("search response of %v should be %v but have %v", input, "Valencia,Vancouver", res))
	}
}

func TestSearchLowerCase(t *testing.T) {
	expected := []string{"London"}
	input := "lo"
	res := Search(input)
	if len(expected) != len(res) {
		t.Error(fmt.Sprintf("search response of %v should be %v but have %v", input, "London", res))
	}
}

func TestSearchUpperCase(t *testing.T) {
	expected := []string{"Bangkok", "Dubai"}
	input := "BA"
	res := Search(input)
	if len(expected) != len(res) {
		t.Error(fmt.Sprintf("search response of %v should be %v but have %v", input, "Bangkok,Dubai", res))
	}
}

func TestSearchContain(t *testing.T) {
	expected := []string{"Budapest"}
	input := "ape"
	res := Search(input)
	if len(expected) != len(res) {
		t.Error(fmt.Sprintf("search response of %v should be %v but have %v", input, "Budapest", res))
	}
}

func TestSearchAll(t *testing.T) {
	expected := []string{"Paris", "Budapest", "Skopje", "Rotterdam", "Valencia", "Vancouver", "Amsterdam",
		"Vienna", "Sydney", "New York City", "London", "Bangkok", "Hong Kong",
		"Dubai", "Rome", "Istanbul"}
	input := "*"
	res := Search(input)
	if len(expected) != len(res) {
		t.Error(fmt.Sprintf("search response of %v should return %v cities but got %v cities", input, len(expected), len(res)))
	}
}
