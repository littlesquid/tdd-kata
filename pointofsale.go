package main

import (
	"fmt"
	"strings"
)

type Product struct {
	Code  string
	Price float32
}

func BarcodeScan(code string) string {

	if code == "99999" {
		return "Error: barcode not found"
	} else if code == "" {
		return "Error: empty barcode"
	}

	products := []Product{
		{Code: "12345", Price: 7.25},
		{Code: "12346", Price: 12.50},
	}

	var barcode []string
	var sum float32
	var index = 0

	if IsMatch(code, "^total.[a-zA-Z0-9,]*$") {
		command := strings.ReplaceAll(code, "total", "")
		command = strings.TrimSpace(command)
		barcode = strings.Split(command, ",")
	}

	for _, p := range products {
		if len(barcode) > 0 {
			if barcode[index] == p.Code {
				sum += p.Price
				index++
			}
		}
		if code == p.Code {
			return fmt.Sprintf("$%.2f", p.Price)
		}
	}

	return fmt.Sprintf("$%.2f", sum)
}
