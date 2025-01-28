package main

import (
	"sort"
	"strings"
)

func groupAnagrams(input []string) []int {
	anagramMap := make(map[string]int)

	for _, word := range input {
		sortedWord := sortString(word)
		anagramMap[sortedWord]++
	}

	result := make([]int, 0)
	for _, count := range anagramMap {
		if count > 2 {
			multipleOf2 := count / 2
			reminder := count % 2
			for i := 0; i < multipleOf2; i++ {
				result = append(result, 2)
			}
			result = append(result, reminder)
		} else {
			result = append(result, count)
		}
	}

	return result
}

func sortString(word string) string {
	chars := strings.Split(word, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
