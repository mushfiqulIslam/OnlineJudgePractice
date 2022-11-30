package main

func reverse(x int) int {
	var reverseValue int
	negative := false
	if x < 0 {
		negative = true
		x = x * -1
	}
	for mod := 10; x > 0; x = x / mod {
		reverseValue = reverseValue*10 + (x % mod)
	}

	if negative {
		reverseValue = reverseValue * -1
		if -2147483648 > reverseValue {
			return 0
		}
	} else if reverseValue > 2147483647 {
		return 0
	}

	return reverseValue
}
