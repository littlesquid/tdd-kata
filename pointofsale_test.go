package main

import (
	"fmt"
	"testing"
)

func TestBarcodeScannerCase1(t *testing.T) {
	input := "12345"
	expected := "$7.25"
	res := BarcodeScan(input)
	if expected != res {
		t.Error(fmt.Sprintf("price of %v should be %v but have %v", input, expected, res))
	}
}

func TestBarcodeScannerCase2(t *testing.T) {
	input := "12346"
	expected := "$12.50"
	res := BarcodeScan(input)
	if expected != res {
		t.Error(fmt.Sprintf("price of %v should be %v but have %v", input, expected, res))
	}
}

func TestBarcodeScannerCase3(t *testing.T) {
	input := "99999"
	expected := "Error: barcode not found"
	res := BarcodeScan(input)
	if expected != res {
		t.Error(fmt.Sprintf("price of %v should be %v but have %v", input, expected, res))
	}
}

func TestBarcodeScannerCase4(t *testing.T) {
	input := ""
	expected := "Error: empty barcode"
	res := BarcodeScan(input)
	if expected != res {
		t.Error(fmt.Sprintf("price of %v should be %v but have %v", input, expected, res))
	}
}

func TestBarcodeScannerTotalCommand(t *testing.T) {
	input := "total 12345,12346"
	expected := "$19.75"
	res := BarcodeScan(input)
	if expected != res {
		t.Error(fmt.Sprintf("price of %v should be %v but have %v", input, expected, res))
	}
}
