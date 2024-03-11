package main

import "strings"

var digitMapper = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	var result []string
	if len(digits) == 0 {
		return result
	}

	for i := 0; i < len(digits); i++ {
		letters := digitMapper[digits[i]]
		var temp []string

		if result == nil {
			result = strings.Split(letters, "")
		} else {
			for _, letter := range letters {
				for _, prevLetter := range result {
					temp = append(temp, prevLetter+string(letter))
				}
			}
			result = temp
		}
	}
	return result
}
