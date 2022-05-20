package main

import (
	"fmt"
	"strings"
)

func max(m, n int) int {
	if n > m {
		return n
	}
	return m
}

func lengthOfLongestSubstring(s string) int {
	var subString string
	var start, longest int
	for i := 0; i < len(s); i++ {
		tmp := string(s[i])
		if strings.Contains(subString, tmp) {
			start = strings.Index(subString, tmp) + 1
			subString = subString[start:]
			subString += tmp
		} else {
			subString += tmp
		}
		longest = max(longest, len(subString))
	}
	return longest
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
