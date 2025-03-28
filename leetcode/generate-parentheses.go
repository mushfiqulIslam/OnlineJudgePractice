func generateParenthesis(n int) []string {
	var parenthesisCombinations []string
	var dfs func(openP, closeP int, s string)
	dfs = func(openP, closeP int, s string) {
		if openP == closeP && openP+closeP == n*2 {
			parenthesisCombinations = append(parenthesisCombinations, s)
			return
		}

		if openP < n {
			dfs(openP+1, closeP, s+"(")
		}

		if closeP < openP {
			dfs(openP, closeP+1, s+")")
		}
	}
	dfs(0, 0, "")
	return parenthesisCombinations
}
