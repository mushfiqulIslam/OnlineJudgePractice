package main

import (
	"fmt"
	"sort"
	"strings"
)

func palindromePairs(words []string) [][]int {
	indexes := make([]int, len(words))
	palindromeMap := map[string]int{}
	var palindromePairs [][]int
	for i := range indexes {
		indexes[i] = i
	}
	fmt.Println(indexes)
	// Sorting to store small words into map first because small words can't be matched with big word
	// But portion of big word can be matched with small word and create a palindrome from map.
	// while checking s we do not know is it palindrome or not. If s is stored previously we can find
	// substring of lls-> s and match it to the index 3 s
	sort.Slice(indexes, func(i, j int) bool {
		return len(words[indexes[i]]) < len(words[indexes[j]])
	})
	for _, wordIndex := range indexes {
		word := words[wordIndex]
		candidate, remainInLeft := getCandidate(word, -1, 0)
		updatePalindromePairs(&palindromePairs, palindromeMap, candidate, wordIndex, remainInLeft)
		for pos := range word {
			candidate, remainInLeft := getCandidate(word, pos, pos)
			updatePalindromePairs(&palindromePairs, palindromeMap, candidate, wordIndex, remainInLeft)
			candidate, remainInLeft = getCandidate(word, pos, pos+1)
			updatePalindromePairs(&palindromePairs, palindromeMap, candidate, wordIndex, remainInLeft)
		}
		palindromeMap[reverse(word)] = wordIndex
	}

	return palindromePairs
}

func getCandidate(word string, leftPos, rightPos int) (string, bool) {
	for leftPos >= 0 && rightPos < len(word) {
		if word[leftPos] != word[rightPos] {
			return "-1", false
		}
		leftPos--
		rightPos++
	}
	if leftPos == -1 {
		return word[rightPos:], true
	}
	return word[:leftPos+1], false
}

func updatePalindromePairs(palindromePairs *[][]int, palindromeMap map[string]int, candidate string, index1 int, remainInLeft bool) {
	if candidate == "-1" {
		return
	}
	if index2, ok := palindromeMap[candidate]; ok {
		if candidate == "" {
			*palindromePairs = append(*palindromePairs, []int{index1, index2})
			*palindromePairs = append(*palindromePairs, []int{index2, index1})
		} else {
			if remainInLeft {
				// matched with a reverse so reverse will come first revesre(abcd)->dcba
				// So, abcd will sit in the second position
				*palindromePairs = append(*palindromePairs, []int{index2, index1})
			} else {
				*palindromePairs = append(*palindromePairs, []int{index1, index2})
			}
		}
	}
}

func reverse(word string) string {
	str := strings.Builder{}
	for i := len(word) - 1; i >= 0; i-- {
		str.WriteByte(word[i])
	}
	return str.String()
}
