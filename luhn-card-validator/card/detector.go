package card

import (
	"regexp"
	"strings"
)

type CardType string

const (
	Unknown    CardType = "Unknown"
	Visa       CardType = "Visa"
	MasterCard CardType = "MasterCard"
	Amex       CardType = "American Express"
	Discover   CardType = "Discover"
)

type CardInfo struct {
	Number string
	Valid  bool
	Type   CardType
}

// detectCardType identifies card brand based on prefix and length
func detectCardType(number string) CardType {
	switch {
	case strings.HasPrefix(number, "4") && (len(number) == 13 || len(number) == 16 || len(number) == 19):
		return Visa
	case regexp.MustCompile(`^5[1-5]`).MatchString(number) && len(number) == 16:
		return MasterCard
	case regexp.MustCompile(`^2(2[2-9]|[3-6]|7[01]|720)`).MatchString(number) && len(number) == 16:
		return MasterCard // newer MasterCard range
	case regexp.MustCompile(`^3[47]`).MatchString(number) && len(number) == 15:
		return Amex
	case regexp.MustCompile(`^6(?:011|5|4[4-9])`).MatchString(number) && len(number) == 16:
		return Discover
	default:
		return Unknown
	}
}

// Analyze returns sanitized card data, validity, and type
func Analyze(input string) (CardInfo, error) {
	sanitized, err := sanitizeInput(input)
	if err != nil {
		return CardInfo{}, err
	}

	valid := validateLuhn(sanitized)
	cardType := detectCardType(sanitized)

	return CardInfo{
		Number: sanitized,
		Valid:  valid,
		Type:   cardType,
	}, nil
}
