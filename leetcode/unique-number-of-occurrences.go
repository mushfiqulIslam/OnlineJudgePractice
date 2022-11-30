func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, value := range arr {
		m[value]++
	}

	var occurrences []int
	for _, v := range m {
		occurrences = append(occurrences, v)
	}
	sort.Ints(occurrences)
	var prevOccurrence int
	for i := 0; i < len(occurrences); i++ {
		if prevOccurrence == occurrences[i] {
			return false
		}
		prevOccurrence = occurrences[i]
	}
	return true
}
