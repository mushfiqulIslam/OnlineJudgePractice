package main

import "fmt"

func main() {
	var numOfCase, n, p, q, sumOfWeight, count, w int
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&n, &p, &q)
		sumOfWeight, count = 0, 0
		for ; n > 0; n-- {
			fmt.Scan(&w)
			if count+1 <= p && sumOfWeight+w <= q {
				sumOfWeight += w
				count++
			}
		}
		fmt.Printf("Case %d: %d\n", i, count)

	}
}
