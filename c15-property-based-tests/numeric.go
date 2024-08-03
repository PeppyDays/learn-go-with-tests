package c15

import (
	"errors"
	"strings"
)

type arabicToRomanChar struct {
	arabic    int
	romanChar string
}

var arabicToRomanChars = []arabicToRomanChar{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertArabicToRomanNumeral(arabic int) string {
	var roman strings.Builder
	for arabic > 0 {
		matched, _ := findMatchedLargestRomanChar(arabic)
		roman.WriteString(matched.romanChar)
		arabic -= matched.arabic
	}
	return roman.String()
}

func findMatchedLargestRomanChar(arabic int) (arabicToRomanChar, error) {
	for _, arm := range arabicToRomanChars {
		if arabic >= arm.arabic {
			return arm, nil
		}
	}
	return arabicToRomanChar{0, ""}, errors.New("not found the matched Roman character")
}

func ConvertRomanToArabicNumeral(roman string) int {
	arabic := 0
	for _, arm := range arabicToRomanChars {
		for strings.HasPrefix(roman, arm.romanChar) {
			arabic += arm.arabic
			roman = strings.TrimPrefix(roman, arm.romanChar)
		}
	}
	return arabic
}
