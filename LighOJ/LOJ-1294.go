package main

import (
	"fmt"
)

func main() {
	var numOfCase, n, m int
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&n, &m)
		fmt.Printf("Case %d: %d\n", i, (n>>1)*m)
	}
}
