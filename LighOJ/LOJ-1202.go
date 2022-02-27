package main

import (
	"fmt"
	"math"
)

func main() {
	var numOfCase, r1, c1, r2, c2, move1, move2 int
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&r1, &c1, &r2, &c2)
		move1 = (r1 + c1) % 2
		move2 = (r2 + c2) % 2
		if move1 != move2 {
			fmt.Printf("Case %d: impossible\n", i)
		} else {
			if math.Abs(float64(r1-r2)) == math.Abs(float64(c1-c2)) {
				fmt.Printf("Case %d: 1\n", i)
			} else {
				fmt.Printf("Case %d: 2\n", i)
			}
		}
	}
}
