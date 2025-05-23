package main

import (
	"fmt"
	"log"

	"luhn-card-validator/card"
)

func main() {
	testCases := []string{
		"4111 1111 1111 1111", // Visa
		"5500 0000 0000 0004", // MasterCard
		"3400 0000 0000 009",  // Amex
		"6011 0000 0000 0004", // Discover
		"1234 5678 9012 3456", // Inválida
	}

	for _, tc := range testCases {
		info, err := card.Analyze(tc)
		if err != nil {
			log.Printf("Error con '%s': %v\n", tc, err)
			continue
		}
		fmt.Printf("Tarjeta: %s | Tipo: %s | Válida: %t\n", tc, info.Type, info.Valid)
	}
}
