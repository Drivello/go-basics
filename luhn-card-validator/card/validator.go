package card

import (
	"fmt"
	"strings"
	"unicode"
)

func sanitizeInput(input string) (string, error) {
	var sb strings.Builder
	for _, r := range input {
		if unicode.IsSpace(r) {
			continue
		}
		if !unicode.IsDigit(r) {
			return "", fmt.Errorf("input inválido: contiene caracteres no numéricos")
		}
		sb.WriteRune(r)
	}
	return sb.String(), nil
}

// validateLuhn returns true if the given card number passes the Luhn check.
// This algorithm is used to validate identification numbers like credit card numbers.
func validateLuhn(number string) bool {
	totalSum := 0        // final sum to check against mod 10
	doubleDigit := false // flag to determine whether to double the digit

	// Traverse the card number from right to left
	for i := len(number) - 1; i >= 0; i-- {
		digit := charToDigit(number[i]) // convert character to its integer value

		if doubleDigit {
			// Step 1: double the digit
			doubled := digit * 2

			// Step 2: if the result is > 9, subtract 9 (same as adding the two digits)
			if doubled > 9 {
				doubled -= 9
			}

			// Step 3: add to total
			totalSum += doubled
		} else {
			// Add digit as-is
			totalSum += digit
		}

		// Alternate between doubling and not
		doubleDigit = !doubleDigit
	}

	// If the total sum is divisible by 10, the number is valid
	return totalSum%10 == 0
}

// charToDigit converts a digit character ('0'–'9') to its corresponding integer value.
// It works by subtracting the ASCII value of '0' (which is 48) from the input character.
// For example: '5' - '0' = 53 - 48 = 5
func charToDigit(c byte) int {
	return int(c - '0')
}
