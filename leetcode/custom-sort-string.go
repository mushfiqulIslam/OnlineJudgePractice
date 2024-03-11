package main

func customSortString(order string, s string) string {
	stringCount := make(map[byte]int)
	result := make([]byte, 0, len(s))

	for i, _ := range s {
		stringCount[s[i]]++
	}

	for i, _ := range order {
		if val, ok := stringCount[order[i]]; ok {
			for j := 0; j < val; j++ {
				result = append(result, order[i])
			}
		}
		delete(stringCount, order[i])
	}

	for key, count := range stringCount {
		for i := 0; i < count; i++ {
			result = append(result, key)
		}
	}
	return string(result)
}
