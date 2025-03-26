func isValid(s string) bool {
	var stack []rune
	brackets := map[rune]rune{')': '(', '}': '{', ']': '['}
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != brackets[char] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return false
	}
	return true
}
