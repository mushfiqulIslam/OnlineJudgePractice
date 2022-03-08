package main

import (
	"fmt"
	"math"
)

func main() {
	var numOfCase int
	var v1, v2, v3, a1, a2, d1, d2, d3, t1, t2 float64
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&v1, &v2, &v3, &a1, &a2)
		t1 = v1 / a1
		t2 = v2 / a2
		d1 = v1*t1 - 0.5*a1*t1*t1
		d2 = v2*t2 - 0.5*a2*t2*t2
		d3 = v3 * math.Max(t1, t2)
		fmt.Printf("Case %d: %.10f %.10f\n", i, d1+d2, d3)
	}
}
