func convert(s string, numRows int) string {
	lengthOfString := len(s)
	if lengthOfString <= 2 || numRows == 1 || lengthOfString <= numRows {
		return s
	}

	output := ""
	skipColumn := 0
	if numRows >= 2 {
		skipColumn = numRows - 2
	}
	for i := 0; i < numRows; i++ {
		index := 0
		rowWordCount := 1
		for index < lengthOfString {
			index = (numRows+skipColumn)*(rowWordCount-1) + i
			if index >= lengthOfString {
				break
			}
			output += string(s[index])
			rowWordCount++
			if i != 0 && i != numRows-1 {
				index = (numRows+skipColumn)*(rowWordCount-1) - i
				if index >= lengthOfString {
					break
				}
				output += string(s[index])
			}
		}
	}
	return output
}
