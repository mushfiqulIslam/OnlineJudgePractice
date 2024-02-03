package main

func isPalindrome(x int) bool {
	var revVal, temp int

	for temp = x; temp > 0; temp /= 10 {
		revVal = revVal*10 + temp%10
	}
	if x == revVal {
		return true
	}

	return false
}
