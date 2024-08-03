package c15_test

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"

	"example.com/c15-property-based-tests"
)

func TestConvertArabicToRomanNumeral(t *testing.T) {
	testScenarios := []struct {
		arabic int
		roman  string
	}{
		{arabic: 1, roman: "I"},
		{arabic: 2, roman: "II"},
		{arabic: 3, roman: "III"},
		{arabic: 4, roman: "IV"},
		{arabic: 5, roman: "V"},
		{arabic: 6, roman: "VI"},
		{arabic: 7, roman: "VII"},
		{arabic: 8, roman: "VIII"},
		{arabic: 9, roman: "IX"},
		{arabic: 10, roman: "X"},
		{arabic: 14, roman: "XIV"},
		{arabic: 18, roman: "XVIII"},
		{arabic: 20, roman: "XX"},
		{arabic: 39, roman: "XXXIX"},
		{arabic: 40, roman: "XL"},
		{arabic: 47, roman: "XLVII"},
		{arabic: 49, roman: "XLIX"},
		{arabic: 50, roman: "L"},
		{arabic: 100, roman: "C"},
		{arabic: 90, roman: "XC"},
		{arabic: 400, roman: "CD"},
		{arabic: 500, roman: "D"},
		{arabic: 900, roman: "CM"},
		{arabic: 1000, roman: "M"},
		{arabic: 1984, roman: "MCMLXXXIV"},
		{arabic: 3999, roman: "MMMCMXCIX"},
		{arabic: 2014, roman: "MMXIV"},
		{arabic: 1006, roman: "MVI"},
		{arabic: 798, roman: "DCCXCVIII"},
	}

	for _, scenario := range testScenarios {
		t.Run("sut converts %v to %v", func(t *testing.T) {
			sut := c15.ConvertArabicToRomanNumeral
			actual := sut(scenario.arabic)
			assert.Equal(t, scenario.roman, actual)
		})
	}

	for _, scenario := range testScenarios {
		t.Run("sut converts %v to %v", func(t *testing.T) {
			sut := c15.ConvertRomanToArabicNumeral
			actual := sut(scenario.roman)
			assert.Equal(t, scenario.arabic, actual)
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic int) bool {
		roman := c15.ConvertArabicToRomanNumeral(arabic)
		recoveredArabic := c15.ConvertRomanToArabicNumeral(roman)
		return recoveredArabic == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
