func maxScore(s string) int {
	var result []int
	result = make([]int, len(s)-1)
	sSlice := strings.Split(s, "")
	prevCount := 0
	for index, r := range sSlice {
		if index == len(sSlice)-1 {
			break
		}
		val, _ := strconv.Atoi(r)
		if val == 0 {
			prevCount += 1
			result[index] = prevCount

		}
	}
	prevCount = 0
	maximum, decrease := 0, 0
	for i := len(sSlice) - 1; i > 0; i-- {
		val2, _ := strconv.Atoi(sSlice[i])
		decrease = len(sSlice) - i
		if val2 == 1 {
			prevCount += 1
			result[len(result)-decrease] = result[len(result)-decrease] + prevCount
		}
		if result[len(result)-decrease] > maximum {
			maximum = result[len(result)-decrease]
		}
	}
	//fmt.Print(result)
	return maximum
}
