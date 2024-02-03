package main

// https://www.youtube.com/watch?v=mmjDZGSr7EA
// https://www.youtube.com/watch?v=JWTqsNvtwP4
// https://www.youtube.com/watch?v=l3hda49XcDE
// https://www.geeksforgeeks.org/implementing-regular-expression-matching/
func isMatch(s string, p string) bool {
	var empty, nonempty bool
	stringLen := len(s)
	patternLen := len(p)

	table := make([][]bool, stringLen+1)
	for i := 0; i < len(table); i++ {
		table[i] = make([]bool, patternLen+1)
	}
	table[0][0] = true

	for j := 2; j < patternLen+1; j++ {
		if p[j-1] == '*' {
			table[0][j] = table[0][j-2]
		}
	}

	for i := 1; i < stringLen+1; i++ {
		for j := 1; j < patternLen+1; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				table[i][j] = table[i-1][j-1]
			} else if p[j-1] == '*' {
				empty = table[i][j-2]
				nonempty = (s[i-1] == p[j-2] || p[j-2] == '.') && table[i-1][j]
				table[i][j] = empty || nonempty
			}
		}
	}

	return table[stringLen][patternLen]
}
