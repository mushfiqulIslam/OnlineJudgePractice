package main

import "fmt"

func main() {
	var numOfCase, n, count int
	var res string
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&n)
		count = 0
		for ; n > 0; n /= 2 {
			if (n % 2) == 1 {
				count++
			}
		}
		if (count % 2) == 1 {
			res = "odd"
		} else {
			res = "even"
		}
		fmt.Printf("Case %d: %s\n", i, res)
	}
}
