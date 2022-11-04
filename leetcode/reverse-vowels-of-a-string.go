func reverseVowels(s string) string {
	type stringVowelNumberPair struct {
		vowelPos  int
		stringPos int
	}
	vowels := [10]string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	stringSlice := strings.Split(s, "")
	var vowelIndexes []stringVowelNumberPair
	for cPos, char := range stringSlice {
		for vPos, vowel := range vowels {
			if vowel == char {
				vowelIndexes = append(vowelIndexes, stringVowelNumberPair{
					vowelPos:  vPos,
					stringPos: cPos,
				})
				break
			}
		}
	}
	revCursor := len(vowelIndexes) - 1
	for _, indexPair := range vowelIndexes {
		stringSlice[indexPair.stringPos] = vowels[vowelIndexes[revCursor].vowelPos]
		revCursor--
	}
	return strings.Join(stringSlice, "")
}
