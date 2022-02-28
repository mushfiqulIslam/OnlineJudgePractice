package main

import "fmt"

func main() {
	var numOfCase, value, revVal, temp int
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		revVal = 0
		fmt.Scan(&value)
		for temp = value; temp > 0; temp /= 10 {
			revVal = revVal*10 + temp%10
		}
		if value == revVal {
			fmt.Printf("Case %d: Yes\n", i)
		} else {
			fmt.Printf("Case %d: No\n", i)
		}

	}
}
