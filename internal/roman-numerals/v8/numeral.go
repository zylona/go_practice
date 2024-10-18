package v8

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var RomanNumerals = []RomanNumeral{
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

func ConvertToRoman(arabic uint16) string {

	var result strings.Builder

	for _, numeral := range RomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

func (r RomanNumeral) ValueOf(symbol string) uint16 {
	for _, numeral := range RomanNumerals {
		if numeral.Symbol == symbol {
			return numeral.Value
		}
	}
	return 0
}

func ConvertToArabic(roman string) uint16 {
	total := uint16(0)

	previous := uint16(0)
	for i := 0; i < len(roman); i++ {
		current := RomanNumeral{}.ValueOf(string(roman[i]))
		if current > previous {
			total -= previous
		} else {
			total += previous
		}
		previous = current
	}
	total += previous
	return total
}
