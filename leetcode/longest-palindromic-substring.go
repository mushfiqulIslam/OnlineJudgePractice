func longestPalindrome(s string) string {
	start, maxLength := 0, 1

	if len(s) <= 1 {
		return s
	}

	for i := 0; i < len(s)-maxLength/2; i++ {
		newStart, newLength := checkPalindromeSubString(s, i, i)
		if newLength > maxLength {
			maxLength = newLength
			start = newStart
		}

		newStart, newLength = checkPalindromeSubString(s, i, i+1)
		if newLength > maxLength {
			maxLength = newLength
			start = newStart
		}
	}
	return s[start : start+maxLength]
}

func checkPalindromeSubString(s string, leftPos, rightPos int) (int, int) {
	for leftPos >= 0 && rightPos < len(s) && s[leftPos] == s[rightPos] {
		leftPos--
		rightPos++
	}
	return leftPos + 1, rightPos - leftPos - 1
}
