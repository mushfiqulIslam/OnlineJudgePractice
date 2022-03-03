package main

import (
	"fmt"
	"math"
)

func main() {
	var numOfCase, n, noseSize, preSize, lieCount int
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&n)
		preSize = 2
		for ; n > 0; n-- {
			fmt.Scan(&noseSize)
			lieCount += int(math.Ceil(float64(noseSize-preSize) / 5.0))
			preSize = noseSize
		}
		fmt.Printf("Case %d: %d\n", i, lieCount)
		lieCount = 0
	}
}
