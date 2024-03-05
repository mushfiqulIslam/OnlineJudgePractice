func minimumLength(s string) int {
	start := 0
	end := len(s) - 1
	for start < end && s[start] == s[end] {
		firstWord := s[start]
		for start <= end && s[end] == firstWord {
			end -= 1
		}
		for start <= end && s[start] == firstWord {
			start += 1
		}
	}
	return end - start + 1
}
