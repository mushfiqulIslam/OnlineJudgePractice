package main

import (
	"fmt"
	"math"
)

func angle(a float64, b float64, c float64) float64 {
	theta := math.Acos((b*b + c*c - a*a) / (2 * b * c))
	return theta
}

func area(r float64, theta float64) float64 {
	return 0.5 * r * r * theta
}

func main() {
	var numOfCase int
	var r1, r2, r3, s, at, a, b, c float64
	fmt.Scan(&numOfCase)
	for i := 1; i <= numOfCase; i++ {
		fmt.Scan(&r1, &r2, &r3)
		a = r2 + r3
		b = r1 + r3
		c = r1 + r2
		s = (a + b + c) / 2.0
		at = math.Sqrt(s * (s - a) * (s - b) * (s - c))
		at -= area(r1, angle(a, b, c))
		at -= area(r2, angle(b, a, c))
		at -= area(r3, angle(c, a, b))
		fmt.Printf("Case %d: %.10f\n", i, at)
	}
}
