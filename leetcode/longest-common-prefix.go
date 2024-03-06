package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		word := strs[i]
		if len(word) == 0 || prefix == "" {
			return ""
		}
		prefix = prefix[:min(len(prefix), len(word))]

		for k := 0; k < len(word) && k < len(prefix); k++ {
			if word[k] != prefix[k] {
				prefix = prefix[:k]
				break
			}
		}
	}
	return prefix
}
