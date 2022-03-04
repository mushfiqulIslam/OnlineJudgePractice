package main

import (
	"fmt"
)

func main() {
	var numOfCase, n, l, w, h, v, max, min int
	var name, maxPerson, minPerson string
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&n)
		max, min = 0, 0
		maxPerson, minPerson = "", ""
		for ; n > 0; n-- {
			fmt.Scan(&name, &l, &w, &h)
			v = l * w * h
			if v > max {
				max = v
				maxPerson = name
			}
			if min == 0 || min > v {
				min = v
				minPerson = name
			}
		}
		if max == min {
			fmt.Printf("Case %d: no thief\n", i)
		} else {
			fmt.Printf("Case %d: %s took chocolate from %s\n", i, maxPerson, minPerson)
		}

	}
}
