import "strings"

func romanToInt(s string) int {
	var output, subtract int
	var nextVal string
	romanIntMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	romanSlice := strings.Split(s, "")
	length := len(romanSlice) - 1
	for index, roman := range romanSlice {
		if index < length {
			nextVal = romanSlice[index+1]
			if roman == "I" && (nextVal == "V" || nextVal == "X") {
				subtract -= romanIntMap["I"]
			} else if roman == "X" && (nextVal == "L" || nextVal == "C") {
				subtract -= romanIntMap["X"]
			} else if roman == "C" && (nextVal == "D" || nextVal == "M") {
				subtract -= romanIntMap["C"]
			} else {
				output += romanIntMap[roman] + subtract
				subtract = 0
			}
		} else {
			output += romanIntMap[roman] + subtract
			subtract = 0
		}
	}
	return output
}
