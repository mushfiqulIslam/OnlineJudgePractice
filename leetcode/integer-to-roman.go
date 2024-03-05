import (
	"strings"
)

func intToRoman(num int) string {
	romanRepresentation := ""
	var romanNumber string
	var value, count int
	for num > 0 {
		if num/1000 > 0 {
			romanNumber = "M"
			value = 1000
		} else if num/900 > 0 {
			romanNumber = "CM"
			value = 900
		} else if num/500 > 0 {
			romanNumber = "D"
			value = 500
		} else if num/400 > 0 {
			romanNumber = "CD"
			value = 400
		} else if num/100 > 0 {
			romanNumber = "C"
			value = 100
		} else if num/90 > 0 {
			romanNumber = "XC"
			value = 90
		} else if num/50 > 0 {
			romanNumber = "L"
			value = 50
		} else if num/40 > 0 {
			romanNumber = "XL"
			value = 40
		} else if num/10 > 0 {
			romanNumber = "X"
			value = 10
		} else if num/9 > 0 {
			romanNumber = "IX"
			value = 9
		} else if num/5 > 0 {
			romanNumber = "V"
			value = 5
		} else if num/4 > 0 {
			romanNumber = "IV"
			value = 4
		} else if num/1 > 0 {
			romanNumber = "I"
			value = 1
		}

		count = num / value
		romanRepresentation += strings.Repeat(romanNumber, count)
		num = num % value
	}
	return romanRepresentation
}
